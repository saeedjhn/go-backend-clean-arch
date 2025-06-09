package outbox

import (
	"context"

	"github.com/saeedjhn/go-backend-clean-arch/configs"
	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/contract"
	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/models"
)

type Repository interface {
	Create(ctx context.Context, evt models.OutboxEvent) error
}

// var _ outboxhandler.Interactor = (Interactor)(nil) // Commented, because it happens import cycle.

type Interactor struct {
	cfg        *configs.Config
	trc        contract.Tracer
	repository Repository
}

func New(
	cfg *configs.Config,
	trc contract.Tracer,
	repository Repository,
) Interactor {
	return Interactor{
		cfg:        cfg,
		trc:        trc,
		repository: repository,
	}
}
