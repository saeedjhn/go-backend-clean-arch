package entity

import (
	"time"

	"github.com/saeedjhn/go-backend-clean-arch/internal/types"
)

type Task struct {
	ID          types.ID
	UserID      types.ID
	Title       string
	Description string
	Status      TaskStatus
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
