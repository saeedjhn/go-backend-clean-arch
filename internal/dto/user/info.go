package user

import (
	"time"

	"github.com/saeedjhn/go-domain-driven-design/internal/types"
)

type Info struct {
	ID        types.ID  `json:"id"`
	Name      string    `json:"name"`
	Mobile    string    `json:"mobile"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
