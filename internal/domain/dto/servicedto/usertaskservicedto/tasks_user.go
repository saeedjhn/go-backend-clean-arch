package usertaskservicedto

import (
	"github.com/saeedjhn/go-backend-clean-arch/internal/domain/entity"
)

type TasksUserRequest struct {
	UserID uint
}

type TasksUserResponse struct {
	Tasks []entity.Task
}
