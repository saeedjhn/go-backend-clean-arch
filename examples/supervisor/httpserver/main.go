package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/saeedjhn/go-backend-clean-arch/internal/adaptor/jsonfilelogger"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/supervisor"
)

func httpServer(ctx context.Context, processName string, terminateChannel chan<- string) error {
	srv := &http.Server{Addr: ":8080"}

	log.Println("server is running on 8080")

	for {
		select {
		case <-ctx.Done():
			return srv.Shutdown(context.Background())
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
