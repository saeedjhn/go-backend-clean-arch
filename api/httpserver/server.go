package httpserver

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	myMiddleware "github.com/saeedjhn/go-backend-clean-arch/api/httpserver/middleware"
	"github.com/saeedjhn/go-backend-clean-arch/api/httpserver/router"
	"github.com/saeedjhn/go-backend-clean-arch/internal/bootstrap"
	"go.uber.org/zap"
	_ "net/http/pprof"
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
	// s.Router.Use(middleware.logger())
	s.Router.Use(middleware.Recover())
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
		LogValuesFunc: func(_ echo.Context, request middleware.RequestLoggerValues) error {
			errMsg := ""
			if request.Error != nil {
				errMsg = request.Error.Error()
			}

			s.App.Logger.Set().Named("http-server").Info("request",
				zap.String("request_id", request.RequestID),
				zap.String("host", request.Host),
				zap.String("content-length", request.ContentLength),
				zap.String("protocol", request.Protocol),
				zap.String("method", request.Method),
				zap.Duration("latency", request.Latency),
				zap.String("error", errMsg),
				zap.String("remote_ip", request.RemoteIP),
				zap.Int64("response_size", request.ResponseSize),
				zap.String("uri", request.URI),
				zap.Int("status", request.Status),
			)

			return nil
		},
	}))

	// Router Setup
	router.Setup(s.App, s.Router)

	var address = fmt.Sprintf(":%s", s.App.Config.HTTPServer.Port)

	s.App.Logger.Set().Named("server").Info("start-echo-server", zap.Any("server-config", s.App.Config.HTTPServer))
	if err := s.Router.Start(address); err != nil {
		fmt.Print("router start error", err)
	}
}
