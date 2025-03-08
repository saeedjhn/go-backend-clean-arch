package claim

import (
	"github.com/labstack/echo/v4"
	"github.com/saeedjhn/go-domain-driven-design/configs"
	"github.com/saeedjhn/go-domain-driven-design/internal/usecase/authentication"
)

func GetClaimsFromEchoContext(c echo.Context) *authentication.Claims {
	return c.Get(configs.AuthMiddlewareContextKey).(*authentication.Claims) //nolint:errcheck // nothing
}
