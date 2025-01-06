package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"path/filepath"

	"github.com/saeedjhn/go-backend-clean-arch/api/delivery/grpc"

	"github.com/saeedjhn/go-backend-clean-arch/configs"
	"github.com/saeedjhn/go-backend-clean-arch/internal/bootstrap"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/cmd/migrations"
)

func main() {
	var (
		confPath string
		fileExt  string
	)

	// Parse command-line flag for environment mode with default value as development
	flag.StringVar(&confPath, "conf", "deployments/development", "config path, e.g., -conf deployments/development")
	flag.StringVar(&fileExt, "ext", "yml", "file extension, e.g., -ext yml")
	flag.Parse()

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

	// Initialize configuration options to specify how the configuration should be loaded.
	cfgOption := configs.Option{
		Prefix:      "",
		Delimiter:   "",
		Separator:   "",
		FilePath:    filesWithExt,
		CallbackEnv: nil,
	}

	// Attempt to load the configuration using the specified options.
	config, err := configs.Load(cfgOption)
	if err != nil {
		log.Fatalf("Error loading configuration with option '%v': %v", cfgOption, err)
	}

	// Bootstrap the application with the provided configuration options.
	app, err := bootstrap.App(config)
	if err != nil {
		log.Fatalf("bootstrap app: %v", err)
	}

	// Log the application configuration at startup
	app.Logger.Infow("App.Startup.Config", "config", app.Config)

	// Run database migrations
	if err = migrations.Up(app); err != nil {
		app.Logger.Fatalf("DB.Migrations.Up: %v", err)
	}

	// Set up signal handling for graceful shutdown (e.g., SIGINT, SIGTERM)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt) // more SIGX (SIGINT, SIGTERM, etc)

	// Start gRPC server in a goroutine
	server := grpc.New(app)
	go func() {
		if err = server.Run(); err != nil {
			app.Logger.Fatalf("Server.GRPC.Run: %v", err)
		}
	}()

	// Wait for termination signal (e.g., Ctrl+C)
	<-quit

	// Graceful shutdown logic
	ctxWithTimeout, cancel := context.WithTimeout(
		context.Background(),
		app.Config.Application.GracefulShutdownTimeout,
	)
	defer cancel()

	// Log received interrupt signal and shutting down gracefully
	app.Logger.Info("App.Shutdown.Gracefully - Received interrupt signal, shutting down gracefully")

	// Close Redis client connection during shutdown
	if err = app.CloseRedisClientConnection(); err != nil {
		app.Logger.Errorf("Close.Redis.Connection: %v", err)
	}

	// Close MySQL connection during shutdown
	if err = app.CloseMysqlConnection(); err != nil {
		app.Logger.Errorf("Close.Mysql.Connection: %v", err)
	}

	// Shutdown tracer during shutdown
	if err = app.ShutdownTracer(ctxWithTimeout); err != nil {
		app.Logger.Errorf("Shutdown.Tracer: %v", err)
	}

	// Optionally, close PostgreSQL or other database connections

	// Wait until graceful shutdown is complete
	<-ctxWithTimeout.Done()

	// Optionally, log or perform any last steps after shutdown completes
}
