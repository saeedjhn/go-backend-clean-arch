package repository

import (
	"context"
	"time"

	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/models"
	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/types"
)

//go:generate mockery --name OutboxEvent
type OutboxEvent interface {
	Create(ctx context.Context, evt models.OutboxEvent) error
	UpdatePublished(ctx context.Context, eventIDs []types.ID, publishedAt time.Time) error
	UpdateUnpublished(ctx context.Context, eventIDs []types.ID, lastRetriedAt time.Time) error
	UnpublishedCount(ctx context.Context, retryThreshold int) (int64, error)
	GetUnPublished(ctx context.Context, offset, limit, retryThreshold int) ([]models.OutboxEvent, error)
}
