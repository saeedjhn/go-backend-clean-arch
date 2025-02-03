package entity

import "github.com/saeedjhn/go-backend-clean-arch/internal/types"

type Authenticable struct {
	ID   types.ID
	Role string
}
