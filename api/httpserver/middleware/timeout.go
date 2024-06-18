package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"time"
)

const errorMessage = "custom timeout error message returns to client"

func Timeout(timeout time.Duration) echo.MiddlewareFunc {

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		to := middleware.TimeoutWithConfig(middleware.TimeoutConfig{
			Skipper:      middleware.DefaultSkipper,
			ErrorMessage: errorMessage,
			OnTimeoutRouteErrorHandler: func(err error, c echo.Context) {
				log.Println(c.Path()) // TODO - Better impl - timeout_middleware
			},
			Timeout: timeout, // for example 30 * time.Second
		})

		return func(c echo.Context) (err error) {
			return to(next)(c)
		}
	}
}
