package httpserver

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	myMiddleware "go-backend-clean-arch-according-to-go-standards-project-layout/api/httpserver/middleware"
	"go-backend-clean-arch-according-to-go-standards-project-layout/api/httpserver/router"
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/bootstrap"
	"go.uber.org/zap"
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
	s.Router.Use(myMiddleware.Timeout(s.App.Config.HTTPServer.Timeout))
	//s.Router.Use(middleware.logger())
	//s.Router.Use(middleware.Recover())
	s.Router.Use(middleware.RequestID())

	s.Router.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:           true,
		LogStatus:        true,
		LogHost:          true,
		LogRemoteIP:      true,
		LogRequestID:     true,
		LogMethod:        true,
		LogContentLength: true,
		LogResponseSize:  true,
		LogLatency:       true,
		LogError:         true,
		LogProtocol:      true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			errMsg := ""
			if v.Error != nil {
				errMsg = v.Error.Error()
			}

			s.App.Logger.Set().Named("http-server").Info("request",
				zap.String("request_id", v.RequestID),
				zap.String("host", v.Host),
				zap.String("content-length", v.ContentLength),
				zap.String("protocol", v.Protocol),
				zap.String("method", v.Method),
				zap.Duration("latency", v.Latency),
				zap.String("error", errMsg),
				zap.String("remote_ip", v.RemoteIP),
				zap.Int64("response_size", v.ResponseSize),
				zap.String("uri", v.URI),
				zap.Int("status", v.Status),
			)

			return nil
		},
	}))

	// Router Setup
	router.Setup(s.App, s.Router)

	address := fmt.Sprintf(":%s", s.App.Config.HTTPServer.Port)
	log.Printf("start echo server on %s\n", address)

	if err := s.Router.Start(address); err != nil {
		fmt.Println("router start error", err)
	}
}
