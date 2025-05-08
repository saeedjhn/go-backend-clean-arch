package usecase

import (
	"context"

	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/contract"
)

//go:generate mockery --name OutboxInteractor
type OutboxInteractor interface {
	Create(ctx context.Context, events []contract.DomainEvent) error
}
