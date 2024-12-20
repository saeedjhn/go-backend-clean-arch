package middleware

import (
	mw "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/saeedjhn/go-backend-clean-arch/configs"
	"github.com/saeedjhn/go-backend-clean-arch/internal/usecase/auth"
)

func Auth(
	authIntr *auth.Interactor,
) echo.MiddlewareFunc {
	return mw.WithConfig(mw.Config{
		ContextKey: configs.AuthMiddlewareContextKey,
		SigningKey: []byte(authIntr.Config.AccessTokenSecret),
		// TODO  - as sign method string to config
		SigningMethod: "HS256",
		ParseTokenFunc: func(_ echo.Context, auth string) (interface{}, error) {
			claims, err := authIntr.ParseToken(authIntr.Config.AccessTokenSecret, auth)
			if err != nil {
				return nil, err
			}

			return claims, nil
		},
	})
}

// const _lenValidAuthorizationKeyFromHeader = 2
//
// func Auth(
//	authInteractor *authusecase.Interactor,
// ) echo.MiddlewareFunc {
//	return func(next echo.HandlerFunc) echo.HandlerFunc {
//		return func(c echo.Context) error {
//			authHeader := c.Request().Header.Get("Authorization")
//			t := strings.Split(authHeader, " ")
//
//			if len(t) == _lenValidAuthorizationKeyFromHeader {
//				authToken := t[1]
//				authorized, err := authInteractor.IsAuthorized(
//					authToken,
//					authInteractor.Config.AccessTokenSecret,
//				)
//				if authorized {
//					parseTokenResp, errParse := authInteractor.ParseToken(userauthservicedto.ParseTokenRequest{
//						Secret: authInteractor.Config.AccessTokenSecret,
//						Token:  authToken,
//					})
//					if errParse != nil {
//						return c.JSON(http.StatusUnauthorized, echo.Map{
//							"status":  false,
//							"message": message.ErrorMsg401UnAuthorized,
//							"errors":  nil,
//						})
//					}
//
//					claim.SetClaimsFromEchoContext(c, configs.AuthMiddlewareContextKey, parseTokenResp.Claims)
//
//					return next(c)
//				}
//
//				return c.JSON(http.StatusUnauthorized, echo.Map{
//					"status":  false,
//					"message": message.ErrorMsg401UnAuthorized,
//					"errors":  err.Error(),
//				})
//			}
//
//			return c.JSON(http.StatusUnauthorized, echo.Map{
//				"status":  false,
//				"message": message.ErrorMsg401UnAuthorized,
//				"errors":  nil,
//			})
//		}
//	}
// }
