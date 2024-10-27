package taskservice

import (
	"github.com/saeedjhn/go-backend-clean-arch/configs"
	"github.com/saeedjhn/go-backend-clean-arch/internal/domain/entity"
)

type Repository interface {
	Create(u entity.Task) (entity.Task, error)
	GetByID(id uint64) (entity.Task, error)
	GetAllByUserID(userID uint64) ([]entity.Task, error)
	IsExistsUser(id uint64) (bool, error)
	// etc
}

type TaskInteractor struct {
	config     *configs.Config
	repository Repository
}

// var _ taskhandler.Interactor = (*TaskInteractor)(nil) // Commented, because it happens import cycle.

func New(
	config *configs.Config,
	repository Repository,
) *TaskInteractor {
	return &TaskInteractor{
		config:     config,
		repository: repository,
	}
}
