package taskdto

import (
	"github.com/saeedjhn/go-backend-clean-arch/internal/domain/entity"
)

type CreateRequest struct {
	UserID      uint64            `json:"user_id"`
	Title       string            `json:"title"`
	Description string            `json:"description"`
	Status      entity.TaskStatus `json:"status"`
}

type CreateResponse struct {
	Task entity.Task `json:"task"`
}
