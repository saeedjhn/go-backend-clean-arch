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
)

func main() {
	var (
		confPath string
		fileExt  string
	)

	flag.StringVar(&confPath, "conf", "configs", "config path, e.g., -conf deployments/development")
	flag.StringVar(&fileExt, "ext", "yml", "file extension, e.g., -ext yml")
	flag.Parse()

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

	cfgOption := configs.Option{
		Prefix:      "",
		Delimiter:   "",
		Separator:   "",
		FilePath:    filesWithExt,
		CallbackEnv: nil,
	}

	config, err := configs.Load(cfgOption)
	if err != nil {
		log.Fatalf("Error loading configuration with option '%v': %v", cfgOption, err)
	}

	app, err := bootstrap.App(config)
	if err != nil {
		log.Fatalf("bootstrap app: %v", err)
	}

	app.Logger.Infow("App.Startup.Config", "config", app.Config)
	app.Logger.Infow("App.Startup.BuildInfo", "buildinfo", app.BuildInfo)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt) // more SIGX (SIGINT, SIGTERM, etc)

	server := grpc.New(app)
	go func() {
		if err = server.Run(); err != nil {
			app.Logger.DPanicf("Server.GRPC.Run: %v", err)
		}
	}()

	// Wait for termination signal (e.g., Ctrl+C)
	<-quit

	ctxWithTimeout, cancel := context.WithTimeout(
		context.Background(),
		app.Config.Application.GracefulShutdownTimeout,
	)
	defer cancel()

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

	if err = app.ShutdownCollector(ctxWithTimeout); err != nil {
		app.Logger.Errorf("Shutdown.Collector: %v", err)
	}

	// Optionally, close PostgreSQL or other database connections

	// Wait until graceful shutdown is complete
	<-ctxWithTimeout.Done()

	// Optionally, log or perform any last steps after shutdown completes
}
