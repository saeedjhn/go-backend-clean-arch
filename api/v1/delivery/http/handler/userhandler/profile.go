package userhandler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/saeedjhn/go-backend-clean-arch/configs"
	"github.com/saeedjhn/go-backend-clean-arch/internal/domain/dto/userdto"
	"github.com/saeedjhn/go-backend-clean-arch/internal/service/authservice"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/claim"
	"go.uber.org/zap"
)

func (u *UserHandler) Profile(c echo.Context) error {
	// Give claims
	claims := claim.GetClaimsFromEchoContext[authservice.Claims](c, configs.AuthMiddlewareContextKey)

	// Usage Use-case
	resp, err := u.userInteractor.Profile(userdto.ProfileRequest{ID: claims.UserID})
	if err != nil {
		code, errResp := u.present.Error(err)

		u.app.Logger.Set().Named("users").Error("profile", zap.Any("error", err.Error()))

		return echo.NewHTTPError(code, errResp)
	}

	return c.JSON(http.StatusOK, u.present.Success(resp))
}
