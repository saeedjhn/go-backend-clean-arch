package userhandler

import (
	"github.com/labstack/echo/v4"
	"go-backend-clean-arch/configs"
	"go-backend-clean-arch/internal/domain/dto/userdto"
	"go-backend-clean-arch/internal/infrastructure/httpstatus"
	"go-backend-clean-arch/internal/infrastructure/richerror"
	"go-backend-clean-arch/pkg/claim"
	"go-backend-clean-arch/pkg/message"
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
