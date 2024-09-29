package taskservice

import (
	"github.com/saeedjhn/go-backend-clean-arch/api/httpserver/handler/taskhandler"
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
	repository Repository
}

var _ taskhandler.Interactor = (*TaskInteractor)(nil)

func New(repository Repository) *TaskInteractor {
	return &TaskInteractor{repository: repository}
}
