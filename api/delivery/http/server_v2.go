package http

import (
	"fmt"
	"github.com/saeedjhn/go-backend-clean-arch/api/delivery/http/handler"

	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4/middleware"
	mymiddleware "github.com/saeedjhn/go-backend-clean-arch/api/delivery/http/middleware"
	"github.com/saeedjhn/go-backend-clean-arch/configs"

	"github.com/labstack/echo/v4"
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

	// s.SetupMiddleware()
	s.Router.Use(middleware.Recover())
	s.Router.Use(middleware.RequestID())
	s.Router.Use(echoprometheus.NewMiddleware(configs.PrometheusSubSytemName))
	s.Router.Use(mymiddleware.Timeout(s.App.Config.HTTPServer.Timeout))
	s.Router.Use(mymiddleware.CORS(s.App.Config.CORS))
	s.Router.Use(mymiddleware.Secure())
	s.Router.Use(mymiddleware.Logger(s.App, configs.LoggerExcludePath))

	// s.RegisterRoutes()
	handler.Setup(s.App, s.Router)

	address := fmt.Sprintf(":%s", s.App.Config.HTTPServer.Port)

	s.App.Logger.Infow("Server.HTTP.Start", "config", s.App.Config.HTTPServer)

	return s.Router.Start(address)
}

// func (s Server) SetupMiddleware() {
// 	s.Router.Use(middleware.Recover())
// 	s.Router.Use(middleware.RequestID())
// 	s.Router.Use(echoprometheus.NewMiddleware(configs.PrometheusSubSytemName))
// 	s.Router.Use(mymiddleware.Timeout(s.App.Config.HTTPServer.Timeout))
// 	s.Router.Use(mymiddleware.CORS(s.App.Config.CORS))
// 	s.Router.Use(mymiddleware.Secure())
// 	s.Router.Use(mymiddleware.Logger(s.App, configs.LoggerExcludePath))
// }
