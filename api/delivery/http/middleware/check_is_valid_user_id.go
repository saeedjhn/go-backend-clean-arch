package middleware

import (
	"net/http"
	"strconv"

	"github.com/saeedjhn/go-backend-clean-arch/internal/entity"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/msg"

	"github.com/saeedjhn/go-backend-clean-arch/pkg/claim"

	"github.com/labstack/echo/v4"
)

func CheckIsValidUserID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		claims := claim.GetClaimsFromEchoContext(c)

		idFromTokenConvertToSTR := strconv.FormatUint(claims.UserID.Uint64(), 10)

		if id != idFromTokenConvertToSTR {
			return echo.NewHTTPError(http.StatusBadRequest, entity.NewErrorResponse(msg.ErrMsg401UnAuthorized, nil))
		}

		return next(c)
	}
}
