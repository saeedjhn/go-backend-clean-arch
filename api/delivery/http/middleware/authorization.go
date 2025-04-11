package middleware

import (
	"net/http"

	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/models"

	"github.com/saeedjhn/go-backend-clean-arch/internal/usecase/authorization"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/claim"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/msg"

	"github.com/labstack/echo/v4"
)

func Authorization(authorizeIntr *authorization.Interactor,
	actions ...models.Action,
) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			claims := claim.GetClaimsFromEchoContext(c)

			isAllowed, err := authorizeIntr.CheckAccess(
				c.Request().Context(),
				claims.RoleIDs,
				c.Path(), // resource (Type: API)
				actions...)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, models.NewErrorResponse(
					msg.ErrorMsg500InternalServerError,
					err.Error()),
				)
			}

			if !isAllowed {
				return echo.NewHTTPError(http.StatusForbidden, models.NewErrorResponse(msg.ErrMsg403Forbidden, nil))
			}

			return next(c)
		}
	}
}
