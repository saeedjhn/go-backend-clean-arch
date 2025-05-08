package main //nolint:cyclop // nothing

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"path/filepath"

	eventdriven "github.com/saeedjhn/go-backend-clean-arch/api/delivery/event_driven"
	grpcserver "github.com/saeedjhn/go-backend-clean-arch/api/delivery/grpc"
	httpserver "github.com/saeedjhn/go-backend-clean-arch/api/delivery/http"
	"github.com/saeedjhn/go-backend-clean-arch/configs"
	"github.com/saeedjhn/go-backend-clean-arch/internal/bootstrap"
)

func main() { //nolint:funlen // +100 lines
	var (
		confPath string
		fileExt  string
	)

	flag.StringVar(&confPath, "conf", "configs", "config path, e.g., -conf configs")
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

	// Set up signal handling for graceful shutdown (e.g., SIGINT, SIGTERM)
	// quit := make(chan bool)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt) // more SIGX (SIGINT, SIGTERM, etc)

	// Start a goroutine to send an interrupt after 20 seconds
	// go func() {
	// 	log.Println("Start a gorotine to send an interrupt after 20 seconds")
	// 	time.Sleep(20 * time.Second)

	// 	quit <- syscall.SIGINT // Sending SIGINT after 20 seconds
	// }()

	hs := httpserver.New(app)
	go func() {
		if err = hs.Run(); err != nil {
			app.Logger.DPanicf("Server.HTTP.Run: %v", err)
		}
	}()

	gs := grpcserver.New(app)
	go func() {
		if err = gs.Run(); err != nil {
			app.Logger.DPanicf("Server.GRPC.Run: %v", err)
		}
	}()

	ed := eventdriven.New(app)
	go func() {
		if err = ed.Run(); err != nil {
			app.Logger.DPanicf("EventDriven.Run: %v", err)
		}
	}()

	// go func() {
	// 	Outbox pattern running..
	// 	time.Sleep(10 * time.Second)
	// 	mq, _ := eventdriven.SetupRabbitMQ(app.Config.RabbitMQ, app.EventRegister)
	// 	err = mq.Publish(contract.Event{
	// 		Type:   events.UsersAccountCreated,
	// 		Payload: []byte("User-123"),
	// 	})
	// 	if err != nil {
	// 		log.Println("rabbitmq error: ", err)
	// 	}
	//
	// 	log.Println("rabbitmq message publish successfull")
	// }()

	go func() {
		// Code for Pprof server setup goes here (if necessary)
	}()

	// Wait for termination signal (e.g., Ctrl+C)
	<-quit

	shutdownCtx, shutdownCancel := context.WithTimeout(
		context.Background(),
		app.Config.Application.GracefulShutdownTimeout,
	)
	defer shutdownCancel()

	app.Logger.Info("App.Shutdown.Gracefully - Received interrupt signal, shutting down gracefully")

	if err = hs.Router.Shutdown(shutdownCtx); err != nil {
		app.Logger.Errorf("Server.HTTP.Shutdown: %v", err)
	}

	if err = ed.Shutdown(shutdownCtx); err != nil {
		app.Logger.Errorf("EventDriven.Shutdown: %v", err)
	}

	if err = app.CloseRedisClientConnection(); err != nil {
		app.Logger.Errorf("Close.Redis.Connection: %v", err)
	}

	if err = app.CloseMysqlConnection(); err != nil {
		app.Logger.Errorf("Close.Mysql.Connection: %v", err)
	}

	if err = app.ShutdownTracer(shutdownCtx); err != nil {
		app.Logger.Errorf("Shutdown.Tracer: %v", err)
	}

	if err = app.ShutdownCollector(shutdownCtx); err != nil {
		app.Logger.Errorf("Shutdown.Collector: %v", err)
	}

	// Optionally, close PostgreSQL or other database connections

	// Wait until graceful shutdown is complete
	<-shutdownCtx.Done()

	// After shutdown is complete, cancel to stop remaining operations.

	// Optionally, log or perform any last steps after shutdown completes
}
