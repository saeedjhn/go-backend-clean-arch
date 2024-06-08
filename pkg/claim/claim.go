package claim

import (
	"github.com/labstack/echo/v4"
	"go-backend-clean-arch/internal/usecase/authusecase"
)

func GetClaimsFromEchoContext(c echo.Context, key string) authusecase.Claims {
	return c.Get(key).(authusecase.Claims)
}

func SetClaimsFromEchoContext(c echo.Context, key string, val interface{}) {
	c.Set(key, val)
}
