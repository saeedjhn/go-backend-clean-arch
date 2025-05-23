package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	_ "net/http/pprof" // #nosec G108
	"os"
	"os/signal"
	"path/filepath"

	"github.com/saeedjhn/go-backend-clean-arch/configs"
	"github.com/saeedjhn/go-backend-clean-arch/internal/bootstrap"
)

func main() {
	var (
		confPath string
		fileExt  string
	)

	flag.StringVar(&confPath, "conf", "deployments/development", "config path, e.g., -conf deployments/development")
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
		log.Fatalf("failed to bootstrap the application: %v", err)
	}

	app.Logger.Infow("App.Startup.Config", "config", app.Config)
	app.Logger.Infow("App.Startup.BuildInfo", "buildinfo", app.BuildInfo)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt) // more SIGX (SIGINT, SIGTERM, etc)

	go func() {
		mux := http.NewServeMux()
		server := http.Server{
			Addr:                         app.Config.Pprof.Port,
			Handler:                      mux,
			DisableGeneralOptionsHandler: false,
			TLSConfig:                    nil,
			ReadTimeout:                  app.Config.Pprof.ReadTimeout,
			ReadHeaderTimeout:            app.Config.Pprof.ReadHeaderTimeout,
			WriteTimeout:                 app.Config.Pprof.WriteTimeout,
			IdleTimeout:                  app.Config.Pprof.WriteTimeout,
			MaxHeaderBytes:               0,
			TLSNextProto:                 nil,
			ConnState:                    nil,
			ErrorLog:                     nil,
			BaseContext:                  nil,
			ConnContext:                  nil,
		}
		app.Logger.Infof("Server.PPROF.Starting - Starting PPROF server on port: %s", app.Config.Pprof.Port)

		if err = server.ListenAndServe(); err != nil {
			app.Logger.Fatalf("Server.PPROF.ListenAndServe - Failed to start PPROF server: %v", err)
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

	// Optionally, close connections

	// Wait until graceful shutdown is complete
	<-ctxWithTimeout.Done()

	// Optionally, log or perform any last steps after shutdown completes
}
