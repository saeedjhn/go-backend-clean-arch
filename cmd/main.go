package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"path/filepath"

	"github.com/saeedjhn/go-backend-clean-arch/api/delivery/grpc"
	"github.com/saeedjhn/go-backend-clean-arch/api/delivery/http"

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

	// Bootstrap the application with the provided configuration options.
	app, err := bootstrap.App(configs.Option{
		Prefix:      "",
		Delimiter:   "",
		Separator:   "",
		FilePath:    filesWithExt,
		CallbackEnv: nil,
	})
	if err != nil {
		log.Fatalf("failed to bootstrap the application: %v", err)
	}

	// Log the application configuration at startup
	app.Logger.Set().Named("Main").Info("Config", zap.Any("config", app.Config))

	// Run database migrations
	if err = migrations.Up(app); err != nil {
		app.Logger.Set().Named("Main").Fatal("Migrations.Up", zap.Error(err))
	}

	// Set up signal handling for graceful shutdown (e.g., SIGINT, SIGTERM)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt) // more SIGX (SIGINT, SIGTERM, etc)

	// Start HTTP server in a goroutine
	hs := http.New(app)
	go func() {
		if err = hs.Run(); err != nil {
			app.Logger.Set().Named("Main").Fatal("Server.HTTP.Run", zap.Error(err))
		}
	}()

	// Start gRPC server in a goroutine
	gs := grpc.New(app)
	go func() {
		if err = gs.Run(); err != nil {
			app.Logger.Set().Named("Main").Fatal("Server.GRPC.Run", zap.Error(err))
		}
	}()

	// Start Pprof server for profiling (Optional - if using)
	go func() {
		// Code for Pprof server setup goes here (if necessary)
	}()
	// Wait for termination signal (e.g., Ctrl+C)

	<-quit

	// Graceful shutdown logic
	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), app.Config.Application.GracefulShutdownTimeout)
	defer cancel()

	// Shutdown HTTP server gracefully
	if err = hs.Router.Shutdown(ctxWithTimeout); err != nil {
		app.Logger.Set().Named("Main").Error("Server.HTTP.Shutdown", zap.Error(err))
	}

	// Log received interrupt signal and shutting down gracefully
	app.Logger.Set().Named("Main").Info("Received.Interrupt.Signal.For.Shutting.Down.Gracefully")

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

	// Optionally, close PostgreSQL or other database connections

	// Wait until graceful shutdown is complete
	<-ctxWithTimeout.Done()

	// Optionally, log or perform any last steps after shutdown completes
}
