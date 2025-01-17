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
)

func main() {
	var (
		confPath string
		fileExt  string
	)

	// Parse command-line flag for environment mode with default value as development
	flag.StringVar(&confPath, "conf", "configs", "config path, e.g., -conf configs")
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
		log.Fatalf("failed to bootstrap the application: %v", err)
	}

	// Log the application configuration at startup
	app.Logger.Infow("App.Startup.Config", "config", app.Config)

	// Set up signal handling for graceful shutdown (e.g., SIGINT, SIGTERM)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt) // more SIGX (SIGINT, SIGTERM, etc)

	// Start HTTP server in a goroutine
	hs := http.New(app)
	go func() {
		if err = hs.Run(); err != nil {
			app.Logger.Fatalf("Server.HTTP.Run: %v", err)
		}
	}()

	// Start gRPC server in a goroutine
	gs := grpc.New(app)
	go func() {
		if err = gs.Run(); err != nil {
			app.Logger.Fatalf("Server.GRPC.Run: %v", err)
		}
	}()

	// Start Pprof server for profiling (Optional - if using)
	go func() {
		// Code for Pprof server setup goes here (if necessary)
	}()

	// Wait for termination signal (e.g., Ctrl+C)
	<-quit

	// Graceful shutdown logic
	ctxWithTimeout, cancel := context.WithTimeout(
		context.Background(),
		app.Config.Application.GracefulShutdownTimeout,
	)
	defer cancel()

	// Shutdown HTTP server gracefully
	if err = hs.Router.Shutdown(ctxWithTimeout); err != nil {
		app.Logger.Errorf("Server.HTTP.Shutdown: %v", err)
	}

	// Log received interrupt signal and shutting down gracefully
	app.Logger.Info("App.Shutdown.Gracefully - Received interrupt signal, shutting down gracefully")

	if err = app.CloseRedisClientConnection(); err != nil {
		app.Logger.Errorf("Close.Redis.Connection: %v", err)
	}

	if err = app.CloseMysqlConnection(); err != nil {
		app.Logger.Errorf("Close.Mysql.Connection: %v", err)
	}

	if err = app.ShutdownTracer(ctxWithTimeout); err != nil {
		app.Logger.Errorf("Shutdown.Tracer: %v", err)
	}

	// Optionally, close PostgreSQL or other database connections

	// Wait until graceful shutdown is complete
	<-ctxWithTimeout.Done()

	// Optionally, log or perform any last steps after shutdown completes
}
