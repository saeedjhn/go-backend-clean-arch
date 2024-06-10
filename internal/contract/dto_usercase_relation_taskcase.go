package contract

import (
	"go-backend-clean-arch/internal/domain"
	"time"
)

type Task struct {
	ID          uint          `json:"id"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	Status      domain.Status `json:"status"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
}

type CreateTaskRequestDTO struct {
	UserID      uint          `param:"id" json:"user_id"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	Status      domain.Status `json:"status"`
}

type CreateTaskResponseDTO struct {
	Task Task `json:"task"`
}

type GetTasksRequestDTO struct {
	UserID uint `json:"user_id"`
}

type GetTasksResponseDTO struct {
	Tasks []Task `json:"tasks"`
}
