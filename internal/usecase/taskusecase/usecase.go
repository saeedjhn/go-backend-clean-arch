package taskusecase

import (
	"context"
	"github.com/saeedjhn/go-backend-clean-arch/configs"
	"github.com/saeedjhn/go-backend-clean-arch/internal/domain/entity"
)

type Repository interface {
	Create(ctx context.Context, u entity.Task) (entity.Task, error)
	GetByID(ctx context.Context, id uint64) (entity.Task, error)
	GetAllByUserID(ctx context.Context, userID uint64) ([]entity.Task, error)
	IsExistsUser(ctx context.Context, id uint64) (bool, error)
	// etc
}

type Interactor struct {
	config     *configs.Config
	repository Repository
}

// var _ taskhandler.Interactor = (*Interactor)(nil) // Commented, because it happens import cycle.

func New(
	config *configs.Config,
	repository Repository,
) *Interactor {
	return &Interactor{
		config:     config,
		repository: repository,
	}
}