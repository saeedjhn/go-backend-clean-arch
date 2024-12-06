package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/saeedjhn/go-backend-clean-arch/internal/bootstrap"
)

func Logger(app *bootstrap.Application) echo.MiddlewareFunc {
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
		LogValuesFunc: func(_ echo.Context, request middleware.RequestLoggerValues) error {
			errMsg := ""
			if request.Error != nil {
				errMsg = request.Error.Error()
			}

			app.Logger.Infow("HTTP", "Request", map[string]interface{}{
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

			return nil
		},
	})
}
