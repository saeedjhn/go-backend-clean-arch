package task

import (
	"github.com/saeedjhn/go-backend-clean-arch/internal/entity"
	"time"
)

type TaskInfo struct {
	ID          uint64            `json:"id"`
	UserID      uint64            `json:"user_id"`
	Title       string            `json:"title"`
	Description string            `json:"description"`
	Status      entity.TaskStatus `json:"status"`
	CreatedAt   time.Time         `json:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at"`
}
