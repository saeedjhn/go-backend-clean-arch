package http

import (
	"fmt"

	myMiddleware "github.com/saeedjhn/go-backend-clean-arch/api/delivery/http/middleware"
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
	// s.Router.Use(middleware.logger())
	s.Router.Use(middleware.Recover())
	s.Router.Use(middleware.RequestID())
	s.Router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     s.App.Config.CORS.AllowOrigins,
		AllowMethods:     s.App.Config.CORS.AllowMethods,
		AllowHeaders:     s.App.Config.CORS.AllowHeaders,
		AllowCredentials: s.App.Config.CORS.AllowCredentials,
	}))

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

			s.App.Logger.Infow("HTTP", "Request", map[string]interface{}{
				"request_id":     request.RequestID,
				"host":           request.Host,
				"content-length": request.ContentLength,
				"protocol":       request.Protocol,
				"method":         request.Method,
				"latency":        request.Latency,
				"error":          errMsg,
				"remote_ip":      request.RemoteIP,
				"response_size":  request.ResponseSize,
				"uri":            request.URI,
				"status":         request.Status,
			})

			// s.App.Logger.Set().Named("HTTP.Server").Info("Request",
			// 	zap.String("request_id", request.RequestID),
			// 	zap.String("host", request.Host),
			// 	zap.String("content-length", request.ContentLength),
			// 	zap.String("protocol", request.Protocol),
			// 	zap.String("method", request.Method),
			// 	zap.Duration("latency", request.Latency),
			// 	zap.String("error", errMsg),
			// 	zap.String("remote_ip", request.RemoteIP),
			// 	zap.Int64("response_size", request.ResponseSize),
			// 	zap.String("uri", request.URI),
			// 	zap.Int("status", request.Status),
			// )

			return nil
		},
	}))

	// Router Setup
	router.Setup(s.App, s.Router)

	address := fmt.Sprintf(":%s", s.App.Config.HTTPServer.Port)

	s.App.Logger.Infow("Start.HTTP.Router", "Server.HTTP.Config", s.App.Config.HTTPServer)

	return s.Router.Start(address)
}
