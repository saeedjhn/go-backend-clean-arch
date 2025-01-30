package claim

import (
	"github.com/labstack/echo/v4"
	"github.com/saeedjhn/go-backend-clean-arch/configs"
	"github.com/saeedjhn/go-backend-clean-arch/internal/usecase/authentication"
)

func GetClaimsFromEchoContext(c echo.Context) *authentication.Claims {
	//nolint:forcetypeassert //defensive programming vs let it crash - log-metric-recover ,...
	return c.Get(configs.AuthMiddlewareContextKey).(*authentication.Claims)
}
