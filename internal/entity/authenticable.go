package entity

import "github.com/saeedjhn/go-domain-driven-design/internal/types"

type Authenticable struct {
	ID   types.ID
	Role string
}
