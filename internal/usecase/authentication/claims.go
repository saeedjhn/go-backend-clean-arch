package authentication

import (
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	jwt.RegisteredClaims
	UserID  uint64   `json:"user_id"`
	RoleIDs []uint64 `json:"role_ids"`
}

// func (c Claims) Valid() error {
//	return c.RegisteredClaims.Valid()
// }
