package middleware

import (
	"net/http"

	"github.com/saeedjhn/go-backend-clean-arch/internal/entity"
	"github.com/saeedjhn/go-backend-clean-arch/internal/usecase/authorization"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/claim"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/msg"

	"github.com/labstack/echo/v4"
)

func Authorization(authorizeIntr *authorization.Interactor,
	actions ...entity.Action,
) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			claims := claim.GetClaimsFromEchoContext(c)

			isAllowed, err := authorizeIntr.CheckAccess(
				c.Request().Context(),
				claims.RoleIDs,
				c.Path(), // resource (Type: API)
				actions...)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, entity.NewErrorResponse(
					msg.ErrorMsg500InternalServerError,
					err.Error()),
				)
			}

			if !isAllowed {
				return echo.NewHTTPError(http.StatusForbidden, entity.NewErrorResponse(msg.ErrMsg403Forbidden, nil))
			}

			return next(c)
		}
	}
}
