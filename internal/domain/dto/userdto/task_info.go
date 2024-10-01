package userdto

import (
	"time"

	"github.com/saeedjhn/go-backend-clean-arch/internal/domain/entity"
)

type TaskInfo struct {
	ID          uint64            `json:"id"`
	Title       string            `json:"title"`
	Description string            `json:"description"`
	Status      entity.TaskStatus `json:"status"`
	CreatedAt   time.Time         `json:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at"`
}
