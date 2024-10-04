package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/saeedjhn/go-backend-clean-arch/internal/delivery/http"

	"github.com/saeedjhn/go-backend-clean-arch/configs"
	"github.com/saeedjhn/go-backend-clean-arch/internal/bootstrap"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/cmd/migrations"
	"go.uber.org/zap"
)

func main() {
	// Bootstrap
	app, err := bootstrap.App(configs.Development)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%#v", app)

	// Log
	app.Logger.Set().Named("main").Info("config", zap.Any("config", app.Config))

	// Migrations
	if err = migrations.Up(app); err != nil {
		log.Fatal(err)
	}

	// Start server
	server := http.New(app)
	go func() {
		server.Serve()
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt) // more SIGX (SIGINT, SIGTERM, etc)
	<-quit

	ctx := context.Background()
	ctxWithTimeout, cancel := context.WithTimeout(ctx, app.Config.Application.GracefulShutdownTimeout)
	defer cancel()

	if err = server.Router.Shutdown(ctxWithTimeout); err != nil {
		log.Println("http server shutdown error", err)
	}

	log.Println("received interrupt signal, shutting down gracefully..")

	// Close all DB connection, etc
	// if err = app.CloseMysqlConnection(); err != nil {
	//	log.Fatal(err)
	// }
	// if err = app.CloseRedisClientConnection(); err != nil {
	//	log.Fatal(err)
	//}

	defer func(app *bootstrap.Application) {
		err = app.CloseRedisClientConnection()
		if err != nil {
			log.Fatal(err)
		}
	}(app)

	defer func(app *bootstrap.Application) {
		err = app.CloseMysqlConnection()
		if err != nil {
			log.Fatal(err)
		}
	}(app)

	// app.ClosePostgresqlConnection() // Or etc..

	<-ctxWithTimeout.Done()
}
