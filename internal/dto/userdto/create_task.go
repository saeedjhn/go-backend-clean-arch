package userdto

import (
	"go-backend-clean-arch/internal/domain"
)

type CreateTaskRequest struct {
	UserID      uint          `param:"id" json:"user_id"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	Status      domain.Status `json:"status"`
}

type CreateTaskResponse struct {
	Task Task `json:"task"`
}
