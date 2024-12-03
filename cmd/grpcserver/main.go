package main

import (
	"context"
	"flag"
	"github.com/saeedjhn/go-backend-clean-arch/api/delivery/grpc"
	"log"
	"os"
	"os/signal"
	"path/filepath"

	"github.com/saeedjhn/go-backend-clean-arch/configs"
	"github.com/saeedjhn/go-backend-clean-arch/internal/bootstrap"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/cmd/migrations"
	"go.uber.org/zap"
)

func main() {
	var (
		env      string
		confPath string
		fileExt  string
	)

	// Parse command-line flag for environment mode with default value as development
	flag.StringVar(&env, "env", configs.Development.String(), "environment mode, e.g., -env development")
	flag.StringVar(&confPath, "conf", "deployments/development", "config path, e.g., -conf deployments/development")
	flag.StringVar(&fileExt, "ext", "yml", "file extension, e.g., -ext yml")
	flag.Parse()

	log.Println("Environment mode:", env)

	// Get the current working directory
	workingDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting current working directory: %v", err)
	}

	log.Println("Working Directory:", workingDir)

	filesWithExt, err := configs.CollectFilesWithExt(
		filepath.Join(workingDir, confPath),
		fileExt,
	)
	if err != nil {
		log.Fatalf(
			"Unexpected error while loading configuration files from directory: %s. Error: %v",
			filepath.Join(workingDir, confPath),
			err,
		)
	}
	// Bootstrap
	app, err := bootstrap.App(configs.Option{
		Prefix:      "",
		Delimiter:   "",
		Separator:   "",
		FilePath:    filesWithExt,
		CallbackEnv: nil,
	})
	if err != nil {
		log.Fatalf("bootstrap app: %v", err)
	}

	// Log
	app.Logger.Set().Named("Main").Info("Config", zap.Any("config", app.Config))

	// Migrations
	if err = migrations.Up(app); err != nil {
		app.Logger.Set().Named("Main").Fatal("Migrations.Up", zap.Error(err))
	}

	// Signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt) // more SIGX (SIGINT, SIGTERM, etc)

	// Start server
	server := grpc.New(app)
	go func() {
		if err = server.Run(); err != nil {
			app.Logger.Set().Named("Main").Fatal("Server.GRPC.Run", zap.Error(err))
		}
	}()

	// Wait for termination signal (e.g., Ctrl+C)
	<-quit

	// Graceful shutdown logic
	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), app.Config.Application.GracefulShutdownTimeout)
	defer cancel()

	// Close Redis client connection during shutdown
	func(app *bootstrap.Application) {
		err = app.CloseRedisClientConnection()
		if err != nil {
			app.Logger.Set().Named("Main").Error("Close.Redis.Connection", zap.Error(err))
		}
	}(app)

	// Close MySQL connection during shutdown
	func(app *bootstrap.Application) {
		err = app.CloseMysqlConnection()
		if err != nil {
			app.Logger.Set().Named("Main").Error("Close.Mysql.Connection", zap.Error(err))
		}
	}(app)

	// Shutdown tracer during shutdown
	func(ctx context.Context, app *bootstrap.Application) {
		err = app.ShutdownTracer(ctx)
		if err != nil {
			app.Logger.Set().Named("Main").Error("Shutdown.Tracer", zap.Error(err))
		}
	}(ctxWithTimeout, app)

	// app.ClosePostgresqlConnection() // Or etc..

	<-quit
}
