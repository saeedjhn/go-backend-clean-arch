package middleware

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/saeedjhn/go-backend-clean-arch/configs"
	"github.com/saeedjhn/go-backend-clean-arch/internal/service/authservice"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/claim"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/message"
)

const _lenValidAuthorizationKeyFromHeader = 2

func Auth(config authservice.Config, authInteractor *authservice.AuthInteractor) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			t := strings.Split(authHeader, " ")

			if len(t) == _lenValidAuthorizationKeyFromHeader {
				authToken := t[1]
				authorized, err := authInteractor.IsAuthorized(authToken, config.AccessTokenSecret)
				if authorized {
					claims, errParse := authInteractor.ParseAccessToken(authToken)
					if errParse != nil {
						return c.JSON(http.StatusUnauthorized, echo.Map{
							"status":  false,
							"message": message.ErrorMsg401UnAuthorized,
							"errors":  nil,
						})
					}
					claim.SetClaimsFromEchoContext(c, configs.AuthMiddlewareContextKey, claims)
					return next(c)
				}
				return c.JSON(http.StatusUnauthorized, echo.Map{
					"status":  false,
					"message": message.ErrorMsg401UnAuthorized,
					"errors":  err.Error(),
				})
			}
			return c.JSON(http.StatusUnauthorized, echo.Map{
				"status":  false,
				"message": message.ErrorMsg401UnAuthorized,
				"errors":  nil,
			})
		}
	}
}
