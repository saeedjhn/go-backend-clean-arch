package http

import (
	"fmt"
	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/saeedjhn/go-backend-clean-arch/configs"

	myMiddleware "github.com/saeedjhn/go-backend-clean-arch/api/delivery/http/middleware" //nolint:typecheck // nothing
	"github.com/saeedjhn/go-backend-clean-arch/api/delivery/http/router"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/saeedjhn/go-backend-clean-arch/internal/bootstrap"
)

type Server struct {
	App    *bootstrap.Application
	Router *echo.Echo
}

func New(
	app *bootstrap.Application,
) *Server {
	return &Server{
		App:    app,
		Router: echo.New(),
	}
}

func (s Server) Run() error {
	s.Router.Debug = s.App.Config.Application.Debug

	// Global Middleware Setup
	s.Router.Use(myMiddleware.Timeout(s.App.Config.HTTPServer.Timeout))
	s.Router.Use(middleware.Recover())
	s.Router.Use(middleware.RequestID())
	s.Router.Use(echoprometheus.NewMiddleware(configs.PrometheusSubSytemName))
	s.Router.Use(myMiddleware.CORS(s.App.Config.CORS))
	s.Router.Use(myMiddleware.Secure())
	s.Router.Use(myMiddleware.Logger(s.App.Logger))

	// Router Setup
	router.Setup(s.App, s.Router)

	address := fmt.Sprintf(":%s", s.App.Config.HTTPServer.Port)

	s.App.Logger.Infow("Start.HTTP.Router", "Server.HTTP.Config", s.App.Config.HTTPServer)

	return s.Router.Start(address)
}
