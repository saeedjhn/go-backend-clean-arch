package authentication

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/types"
)

type Claims struct {
	jwt.RegisteredClaims
	UserID  types.ID   `json:"user_id"`
	RoleIDs []types.ID `json:"role_ids"`
}

// func (c Claims) Valid() error {
//	return c.RegisteredClaims.Valid()
// }
