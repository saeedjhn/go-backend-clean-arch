package userhandler

import (
	"github.com/labstack/echo/v4"
	"github.com/saeedjhn/go-backend-clean-arch/configs"
	"github.com/saeedjhn/go-backend-clean-arch/internal/domain/dto/userdto"
	"github.com/saeedjhn/go-backend-clean-arch/internal/infrastructure/httpstatus"
	"github.com/saeedjhn/go-backend-clean-arch/internal/infrastructure/richerror"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/claim"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/message"
	"go.uber.org/zap"
	"net/http"
)

func (u *UserHandler) Profile(c echo.Context) error {
	// Give claims
	claims := claim.GetClaimsFromEchoContext(c, configs.AuthMiddlewareContextKey)

	// Usage Use-case
	resp, err := u.userInteractor.Profile(
		userdto.ProfileRequest{ID: claims.UserId},
	)
	if err != nil {
		richErr, _ := richerror.Analysis(err)
		code := httpstatus.FromKind(richErr.Kind())

		u.app.Logger.Set().Named("users").Error("profile", zap.Any("error", err.Error()))
		return echo.NewHTTPError(code,
			echo.Map{
				"status":  false,
				"message": richErr.Message(),
				"errors":  richErr.Error(),
			})
	}
	return c.JSON(http.StatusOK,
		echo.Map{
			"status":  true,
			"message": message.MsgUserSeeProfileSuccessfully,
			"data":    resp,
		},
	)
}
