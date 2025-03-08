package task

import (
	"time"

	"github.com/saeedjhn/go-domain-driven-design/internal/types"

	"github.com/saeedjhn/go-domain-driven-design/internal/entity"
)

type Info struct {
	ID          types.ID          `json:"id"`
	UserID      types.ID          `json:"user_id"`
	Title       string            `json:"title"`
	Description string            `json:"description"`
	Status      entity.TaskStatus `json:"status"`
	CreatedAt   time.Time         `json:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at"`
}
