package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/saeedjhn/go-backend-clean-arch/configs"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/claim"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/message"
	"net/http"
	"strconv"
)

func CheckIsValidUserID(next echo.HandlerFunc) echo.HandlerFunc {

	return func(c echo.Context) error {
		id := c.Param("id")
		idFromToken := claim.GetClaimsFromEchoContext(c, configs.AuthMiddlewareContextKey).UserId

		idFromTokenConvertToSTR := strconv.FormatUint(uint64(idFromToken), 10)

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
