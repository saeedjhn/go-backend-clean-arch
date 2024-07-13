package main

import (
	"context"
	"github.com/saeedjhn/go-backend-clean-arch/configs"
	"github.com/saeedjhn/go-backend-clean-arch/internal/bootstrap"
	"go.uber.org/zap"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
)

func main() {
	// Bootstrap
	app := bootstrap.App(configs.Development)

	// Log
	app.Logger.Set().Named("main").Info("config", zap.Any("config", app.Config))

	// Prof
	go func() {
		http.ListenAndServe(app.Config.Pprof.Port, nil)
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
