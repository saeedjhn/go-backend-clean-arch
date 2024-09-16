package claim

import (
	"github.com/labstack/echo/v4"
)

func GetClaimsFromEchoContext[T interface{}](c echo.Context, key string) T {
	return c.Get(key).(T)
}

func SetClaimsFromEchoContext(c echo.Context, key string, val interface{}) {
	c.Set(key, val)
}
