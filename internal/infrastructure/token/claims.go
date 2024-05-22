package token

import (
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	jwt.RegisteredClaims
	UserId uint `json:"id"`
}

//func (c Claims) Valid() error {
//	return c.RegisteredClaims.Valid()
//}
