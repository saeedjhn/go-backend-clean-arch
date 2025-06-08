package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/saeedjhn/go-backend-clean-arch/internal/adapter/jsonfilelogger"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/supervisor"
)

const (
	_port              = ":8000"
	_sleepDuration     = 10 * time.Second
	_readTimeout       = 10 * time.Second
	_readHeaderTimeout = 5 * time.Second
	_writeTimeout      = 10 * time.Second
	_idleTimeout       = 2 * time.Minute
)

func httpServer(ctx context.Context, processName string, terminateChannel chan<- string) error {
	srv := &http.Server{
		Addr:              _port,
		ReadTimeout:       _readTimeout,
		ReadHeaderTimeout: _readHeaderTimeout,
		WriteTimeout:      _writeTimeout,
		IdleTimeout:       _idleTimeout,
	}

	log.Println("server is running on 8080")

	for {
		select {
		case <-ctx.Done():
			return srv.Shutdown(ctx)
		default:
			return srv.ListenAndServe()
		}
	}
}

func main() {
	loggerStrategy := jsonfilelogger.NewDevelopmentStrategy(jsonfilelogger.Config{
		FilePath:         "./logs",
		Console:          true,
		File:             false,
		EnableCaller:     true,
		EnableStacktrace: true,
		Level:            "debug",
	})
	logger := jsonfilelogger.New(loggerStrategy).Configure()

	sv := supervisor.New(10*time.Second, logger)

	sv.Register("http-server", httpServer, nil)
	sv.Start()

	sv.WaitOnShutdownSignal()
}
