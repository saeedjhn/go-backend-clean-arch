package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/saeedjhn/go-backend-clean-arch/api/v1/delivery/grpc"
	"github.com/saeedjhn/go-backend-clean-arch/configs"
	"github.com/saeedjhn/go-backend-clean-arch/internal/bootstrap"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/cmd/migrations"
	"go.uber.org/zap"
)

func main() {
	// Bootstrap
	app, err := bootstrap.App(configs.Development)
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

	defer func(app *bootstrap.Application) {
		err = app.CloseRedisClientConnection()
		if err != nil {
			app.Logger.Set().Named("Main").Error("Close.Redis.Connection", zap.Error(err))
		}
	}(app)

	defer func(app *bootstrap.Application) {
		err = app.CloseMysqlConnection()
		if err != nil {
			app.Logger.Set().Named("Main").Error("Close.Mysql.Connection", zap.Error(err))
		}
	}(app)

	// app.ClosePostgresqlConnection() // Or etc..

	<-quit
}
