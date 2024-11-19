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
		log.Fatal(err)
	}

	// Log
	app.Logger.Set().Named("main").Info("config", zap.Any("config", app.Config))

	// Prof
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

		log.Printf("Server Pprof is starting on %server", app.Config.Pprof.Port)

		if err = server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt) // more SIGX (SIGINT, SIGTERM, etc)
	<-quit

	ctx := context.Background()
	ctxWithTimeout, cancel := context.WithTimeout(ctx, app.Config.Application.GracefulShutdownTimeout)
	defer cancel()

	log.Println("received interrupt signal, shutting down gracefully..")

	<-ctxWithTimeout.Done()
}
