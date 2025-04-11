package models

import (
	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/types"
)

type Authenticable struct {
	ID   types.ID
	Role string
}
