package httpserver

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	myMiddleware "go-backend-clean-arch-according-to-go-standards-project-layout/api/httpserver/middleware"
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
	s.Router.Debug = s.App.Config.Application.Debug

	// Global Middleware Setup
	s.Router.Use(middleware.Logger())
	s.Router.Use(middleware.Recover())
	s.Router.Use(myMiddleware.Timeout(s.App.Config.HTTPServer.Timeout))

	g := s.Router.Group("")
	g.Use(middleware.RequestID())

	// Router Setup
	router.Setup(s.App, g)

	address := fmt.Sprintf(":%s", s.App.Config.HTTPServer.Port)
	log.Printf("start echo server on %s\n", address)

	if err := s.Router.Start(address); err != nil {
		fmt.Println("router start error", err)
	}
}
