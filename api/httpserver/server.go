package httpserver

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go-backend-clean-arch-according-to-go-standards-project-layout/api/httpserver/router"
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/bootstrap"
	"log"
)

type HTTPServer struct {
	App    *bootstrap.Application
	Router *echo.Echo
}

func New(
	app *bootstrap.Application,
) *HTTPServer {
	return &HTTPServer{
		App:    app,
		Router: echo.New(),
	}
}

func (s HTTPServer) Serve() {
	s.Router.Debug = true

	// Global Middleware Setup
	s.Router.Use(middleware.Logger())
	s.Router.Use(middleware.Recover())

	// Router Setup
	router.Setup(s.App, s.Router)

	address := fmt.Sprintf(":%s", s.App.Config.HTTPServer.Port)
	log.Printf("start echo server on %s\n", address)

	if err := s.Router.Start(address); err != nil {
		fmt.Println("router start error", err)
	}
}
