package user

import (
	"time"

	"github.com/saeedjhn/go-backend-clean-arch/internal/types"
)

type UserInfo struct {
	ID        types.ID  `json:"id"`
	Name      string    `json:"name"`
	Mobile    string    `json:"mobile"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
