package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	// Setup
	e := echo.New()
	e.Logger.SetLevel(log.INFO)
	e.GET("/", func(c echo.Context) error {
		time.Sleep(5 * time.Second)
		return c.JSON(http.StatusOK, "OK")
	})

	// Start server
	go func() {
		if err := e.Start(":1323"); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Errorf("err: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt) // more SIGX (SIGINT, SIGTERM, etc)
	<-quit

	ctx := context.Background()
	ctxWithTimeout, cancel := context.WithTimeout(ctx, 8*time.Second)
	defer cancel()

	if err := e.Shutdown(ctxWithTimeout); err != nil {
		fmt.Println("http server shutdown error", err)
	}

	fmt.Println("received interrupt signal, shutting down gracefully..")
	// Close all db connection, etc
	<-ctxWithTimeout.Done()
}
