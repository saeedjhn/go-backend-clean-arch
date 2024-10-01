package usertaskservicedto

import (
	"github.com/saeedjhn/go-backend-clean-arch/internal/domain/entity"
)

type CreateTaskRequest struct {
	UserID      uint64
	Title       string
	Description string
	Status      entity.TaskStatus
}

type CreateTaskResponse struct {
	Task entity.Task
}
