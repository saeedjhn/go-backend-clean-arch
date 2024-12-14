package middleware

import (
	"net/http"
	"strconv"

	"github.com/saeedjhn/go-backend-clean-arch/pkg/claim"

	"github.com/labstack/echo/v4"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/message"
)

func CheckIsValidUserID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		claims := claim.GetClaimsFromEchoContext(c)

		idFromTokenConvertToSTR := strconv.FormatUint(claims.UserID, 10)

		if id != idFromTokenConvertToSTR {
			return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
				"status":  false,
				"message": message.ErrorMsg401UnAuthorized,
				"errors":  nil,
			})
		}

		return next(c)
	}
}
