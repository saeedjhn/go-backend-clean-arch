package middleware

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RequestID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		to := middleware.RequestIDWithConfig(middleware.RequestIDConfig{
			Skipper: middleware.DefaultSkipper,
			Generator: func() string {
				return uuid.New().String()
			},
			//RequestIDHandler: nil,
			TargetHeader: echo.HeaderXRequestID,
		})

		return to(next)(c)
	}
}

func RequestID2() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		to := middleware.RequestIDWithConfig(middleware.RequestIDConfig{
			Generator: func() string {
				return uuid.New().String()
			},
		})

		return func(c echo.Context) (err error) {
			return to(next)(c)
		}
	}
}
