package taskservice

import (
	"github.com/saeedjhn/go-backend-clean-arch/internal/domain/entity"
)

type Repository interface {
	Create(u entity.Task) (entity.Task, error)
	GetByID(id uint) (entity.Task, error)
	GetAllByUserID(userID uint) ([]entity.Task, error)
	IsExistsUser(id uint) (bool, error)
	// etc
}

type TaskInteractor struct {
	repository Repository
}

func New(repository Repository) *TaskInteractor {
	return &TaskInteractor{repository: repository}
}
