package taskdto

import (
	"go-backend-clean-arch/internal/domain/entity"
)

type CreateRequest struct {
	UserID      uint          `json:"user_id"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	Status      entity.Status `json:"status"`
}

type CreateResponse struct {
	Task entity.Task `json:"task"`
}
