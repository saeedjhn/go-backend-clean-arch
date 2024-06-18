package userdto

import (
	"go-backend-clean-arch/internal/domain/entity"
)

type CreateTaskRequest struct {
	UserID      uint   `param:"id" json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type CreateTaskResponse struct {
	Task entity.Task `json:"task"`
}
