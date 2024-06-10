package taskusecase

import (
	"go-backend-clean-arch/internal/domain"
)

type Repository interface {
	Create(u domain.Task) (domain.Task, error)
	GetByID(id uint) (domain.Task, error)
	GetAllByUserID(userID uint) ([]domain.Task, error)
	IsExistsUser(id uint) (bool, error)
	// etc
}

type TaskInteractor struct {
	repository Repository
}

func New(taskGateway Repository) *TaskInteractor {
	return &TaskInteractor{repository: taskGateway}
}
