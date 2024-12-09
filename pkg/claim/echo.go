package claim

import (
	"github.com/labstack/echo/v4"
	"github.com/saeedjhn/go-backend-clean-arch/configs"
	"github.com/saeedjhn/go-backend-clean-arch/internal/usecase/authusecase"
)

func GetClaimsFromEchoContext(c echo.Context) *authusecase.Claims {
	//nolint:forcetypeassert //defensive programming vs let it crash - log-metric-recover ,...
	return c.Get(configs.AuthMiddlewareContextKey).(*authusecase.Claims)
}
