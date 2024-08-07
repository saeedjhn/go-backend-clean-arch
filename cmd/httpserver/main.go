package main

import (
	"context"
	"fmt"
	"github.com/saeedjhn/go-backend-clean-arch/api/httpserver"
	"github.com/saeedjhn/go-backend-clean-arch/configs"
	"github.com/saeedjhn/go-backend-clean-arch/internal/bootstrap"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/cmd/migrations"
	"go.uber.org/zap"
	"log"
	"os"
	"os/signal"
)

func main() {
	// Bootstrap
	app := bootstrap.App(configs.Development)
	log.Printf("%#v", app)

	// Log
	app.Logger.Set().Named("main").Info("config", zap.Any("config", app.Config))

	// Migrations
	migrations.Up(app)

	// Prof
	//go func() {
	//	log.Println("pprof server running")
	//	http.ListenAndServe(":8001", nil)
	//}()

	// Start server
	server := httpserver.New(app)
	go func() {
		server.Serve()
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt) // more SIGX (SIGINT, SIGTERM, etc)
	<-quit

	ctx := context.Background()
	ctxWithTimeout, cancel := context.WithTimeout(ctx, app.Config.Application.GracefulShutdownTimeout)
	defer cancel()

	if err := server.Router.Shutdown(ctxWithTimeout); err != nil {
		fmt.Println("http server shutdown error", err)
	}

	log.Println("received interrupt signal, shutting down gracefully..")
	// Close all db connection, etc
	app.CloseMysqlConnection()
	app.CloseRedisClientConnection()
	//app.ClosePostgresqlConnection() // Or etc..

	<-ctxWithTimeout.Done()
}
