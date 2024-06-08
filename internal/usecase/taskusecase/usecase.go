package taskusecase

import (
	"go-backend-clean-arch/internal/domain"
)

type Repository interface {
	Create(u domain.Task) (domain.Task, error)
	// etc
}

type TaskInteractor struct {
	repository Repository
}

func New(taskGateway Repository) *TaskInteractor {
	return &TaskInteractor{repository: taskGateway}
}
