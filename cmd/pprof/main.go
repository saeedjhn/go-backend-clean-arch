package main

import (
	"context"
	"log"
	"net/http"
	_ "net/http/pprof" // #nosec G108
	"os"
	"os/signal"

	"github.com/saeedjhn/go-backend-clean-arch/configs"
	"github.com/saeedjhn/go-backend-clean-arch/internal/bootstrap"
	"go.uber.org/zap"
)

func main() {
	// Bootstrap
	app, err := bootstrap.App(configs.Development)
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
