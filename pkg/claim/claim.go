package claim

import (
	"github.com/labstack/echo/v4"
	"github.com/saeedjhn/go-backend-clean-arch/internal/service/authservice"
)

func GetClaimsFromEchoContext(c echo.Context, key string) authservice.Claims {
	return c.Get(key).(authservice.Claims)
}

func SetClaimsFromEchoContext(c echo.Context, key string, val interface{}) {
	c.Set(key, val)
}
