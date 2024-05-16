package definemiddeware

import "github.com/labstack/echo/v4"

func Fn1(param interface{}) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			return next(c)
		}
	}
}

func Fn2(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		return next(c)
	}
}
