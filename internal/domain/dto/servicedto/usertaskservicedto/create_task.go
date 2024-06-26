package usertaskservicedto

import (
	"github.com/saeedjhn/go-backend-clean-arch/internal/domain/entity"
)

type CreateTaskRequest struct {
	UserID      uint
	Title       string
	Description string
	Status      entity.Status
}

type CreateTaskResponse struct {
	Task entity.Task
}
