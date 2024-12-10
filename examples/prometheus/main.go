package main

import (
	"errors"
	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus"
	"log"
	"net/http"
)

func main() {
	customCounter := prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "ping_request_count",
			Help: "How many HTTP requests processed, partitioned by status code and HTTP method.",
		},
	)

	// register your new counter metric with default metrics registry
	if err := prometheus.Register(customCounter); err != nil { // Or MustRegister
		log.Fatal(err)
	}

	e := echo.New()
	e.Use(echoprometheus.NewMiddleware("myapp")) // adds middleware to gather metrics

	e.GET("/metrics", echoprometheus.NewHandler()) // adds route to serve gathered metrics

	e.GET("/ping", func(c echo.Context) error {
		customCounter.Inc()

		return c.String(http.StatusOK, "pong")
	})

	if err := e.Start(":8080"); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}
}
