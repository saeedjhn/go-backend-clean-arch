package task

import (
	"context"
	"github.com/saeedjhn/go-backend-clean-arch/internal/entity"

	"github.com/saeedjhn/go-backend-clean-arch/configs"
)

type Repository interface {
	Create(ctx context.Context, u entity.Task) (entity.Task, error)
	GetByID(ctx context.Context, id uint64) (entity.Task, error)
	GetAllByUserID(ctx context.Context, userID uint64) ([]entity.Task, error)
	IsExistsUser(ctx context.Context, id uint64) (bool, error)
	// etc
}

type Interactor struct {
	cfg        *configs.Config
	repository Repository
}

// var _ taskhandler.Interactor = (*Interactor)(nil) // Commented, because it happens import cycle.

func New(
	config *configs.Config,
	repository Repository,
) *Interactor {
	return &Interactor{
		cfg:        config,
		repository: repository,
	}
}
