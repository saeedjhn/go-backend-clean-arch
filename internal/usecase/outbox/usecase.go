package outbox

import (
	"github.com/saeedjhn/go-backend-clean-arch/configs"
	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/contract"
	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/repository"
)

// var _ outboxhandler.Interactor = (*Interactor)(nil) // Commented, because it happens import cycle.

type Interactor struct {
	cfg        *configs.Config
	trc        contract.Tracer
	repository repository.OutboxEvent
}

func New(
	cfg *configs.Config,
	trc contract.Tracer,
	repository repository.OutboxEvent,
) *Interactor {
	return &Interactor{
		cfg:        cfg,
		trc:        trc,
		repository: repository,
	}
}
