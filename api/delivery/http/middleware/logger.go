package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/saeedjhn/go-backend-clean-arch/internal/bootstrap"
)

func Logger(app *bootstrap.Application, excludePath string) echo.MiddlewareFunc {
	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
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
		LogValuesFunc: func(ctx echo.Context, request middleware.RequestLoggerValues) error {
			if request.URI == excludePath {
				return nil
			}

			errMsg := ""
			if request.Error != nil {
				errMsg = request.Error.Error()
			}

			app.Logger.Infow("Sever.HTTP.Request", "request", map[string]interface{}{
				"environment":    app.Config.Application.Env,
				"uri":            request.URI,
				"method":         request.Method,
				"query_string":   request.QueryParams,
				"status":         request.Status,
				"content-length": request.ContentLength,
				"response_size":  request.ResponseSize,
				"remote_ip":      request.RemoteIP,
				"user_agent":     request.UserAgent,
				"host":           request.Host,
				"request_id":     request.RequestID,
				"protocol":       request.Protocol,
				"latency":        request.Latency,
				"error":          errMsg,
			})

			return nil
		},
	})
}
