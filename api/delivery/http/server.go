package http

import (
	"fmt"

	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4/middleware"
	healthcheckhandler "github.com/saeedjhn/go-backend-clean-arch/api/delivery/http/handler/healthcheck"
	prometheushandler "github.com/saeedjhn/go-backend-clean-arch/api/delivery/http/handler/prometheus"
	usertaskhandler "github.com/saeedjhn/go-backend-clean-arch/api/delivery/http/handler/user/task"
	userhandler "github.com/saeedjhn/go-backend-clean-arch/api/delivery/http/handler/user/user"
	mymiddleware "github.com/saeedjhn/go-backend-clean-arch/api/delivery/http/middleware"
	"github.com/saeedjhn/go-backend-clean-arch/configs"

	"github.com/labstack/echo/v4"
	"github.com/saeedjhn/go-backend-clean-arch/internal/bootstrap"
)

type Server struct {
	App                *bootstrap.Application
	Router             *echo.Echo
	userHandler        *userhandler.Handler
	userTaskHandler    *usertaskhandler.Handler
	prometheusHandler  *prometheushandler.Handler
	healthcheckHandler *healthcheckhandler.Handler
}

func New(
	app *bootstrap.Application,
) *Server {
	return &Server{
		App:                app,
		Router:             echo.New(),
		userHandler:        userhandler.New(app.Trc, app.Usecase.AuthIntr, app.Usecase.UserIntr),
		userTaskHandler:    usertaskhandler.New(app.Trc, app.Usecase.AuthIntr, app.Usecase.TaskIntr),
		prometheusHandler:  prometheushandler.New(),
		healthcheckHandler: healthcheckhandler.New(),
	}
}

func (s Server) Run() error {
	s.Router.Debug = s.App.Config.Application.Debug

	s.RegisterMiddleware()

	s.RegisterRoutes()

	address := fmt.Sprintf(":%s", s.App.Config.HTTPServer.Port)

	s.App.Logger.Infow("Server.HTTP.Start", "config", s.App.Config.HTTPServer)

	return s.Router.Start(address)
}

func (s Server) RegisterMiddleware() {
	s.Router.Use(middleware.Recover())
	s.Router.Use(middleware.RequestID())
	s.Router.Use(echoprometheus.NewMiddleware(configs.PrometheusSubSytemName))
	s.Router.Use(mymiddleware.Timeout(s.App.Config.HTTPServer.Timeout))
	s.Router.Use(mymiddleware.CORS(s.App.Config.CORS))
	s.Router.Use(mymiddleware.Secure())
	s.Router.Use(mymiddleware.Logger(s.App, configs.LoggerExcludePath))
}

func (s Server) RegisterRoutes() {
	s.userHandler.SetRoutes(s.Router)
	s.userTaskHandler.SetRoutes(s.Router)
	s.prometheusHandler.SetRoutes(s.Router)
	s.healthcheckHandler.SetRoutes(s.Router)
}
