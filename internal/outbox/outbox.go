package outbox

import (
	"context"
	"fmt"
	"time"

	"github.com/saeedjhn/go-backend-clean-arch/internal/entity"

	"github.com/saeedjhn/go-backend-clean-arch/internal/contract"
	"github.com/saeedjhn/go-backend-clean-arch/internal/types"
)

//go:generate mockery --name Repository
type Repository interface {
	InsertEvent(ctx context.Context, evt entity.Outbox) error
	UpdatePublished(ctx context.Context, eventIDs []types.ID, publishedAt time.Time) error
	UpdateUnpublished(ctx context.Context, eventIDs []types.ID, lastRetriedAt time.Time) error
	UnpublishedCount(ctx context.Context, retryThreshold int64) (int64, error)
	GetUnPublished(ctx context.Context, offset, limit, retryThreshold int) ([]entity.Outbox, error)
}

//go:generate mockery --name Scheduler
type Scheduler interface {
	RepeatTaskEvery(ctx context.Context, fn func(), t time.Duration) error
}

type O struct {
	config     Config
	logger     contract.Logger
	scheduler  Scheduler
	publisher  contract.Publisher
	repository Repository
}

func New(
	config Config,
	logger contract.Logger,
	sch Scheduler,
	pub contract.Publisher,
	repository Repository,
) O {
	return O{
		config:     config,
		logger:     logger,
		scheduler:  sch,
		publisher:  pub,
		repository: repository,
	}
}

func (s O) StartProcessing(ctx context.Context) {
	err := s.scheduler.RepeatTaskEvery(ctx, func() {
		if err := s.processOutBoxEvents(ctx); err != nil {
			s.logger.Errorf("error publishing outbox events: %v", err)
		}
	}, s.config.Interval)
	if err != nil {
		s.logger.Errorf("failed to start processing outbox: %v", err)
	}
}

func (s O) InsertEvent(ctx context.Context, topic string, payload []byte, reTriedCount uint) error {
	e := entity.Outbox{
		Topic:        entity.Topic(topic),
		Payload:      payload,
		IsPublished:  false,
		ReTriedCount: reTriedCount,
	}

	if err := s.repository.InsertEvent(ctx, e); err != nil {
		return err
	}

	return nil
}

func (s O) processOutBoxEvents(ctx context.Context) error {
	unPublishedOutBoxEvents, err := s.repository.GetUnPublished(ctx,
		0, s.config.BatchSize, s.config.RetryThreshold)
	if err != nil {
		s.logger.Errorf("failed to fetch unpublished OutBoxEvents: %w", err)
	}

	if len(unPublishedOutBoxEvents) == 0 {
		s.logger.Info("no unpublished events found.")
		return nil
	}

	outBoxEventsIDs := make([]types.ID, 0, len(unPublishedOutBoxEvents))

	for _, outBoxEvent := range unPublishedOutBoxEvents {
		if err = s.publisher.Publish(entity.Event{
			Topic:   outBoxEvent.Topic,
			Payload: outBoxEvent.Payload,
		}); err != nil {
			s.logger.Infof("failed to publish event ID %v: %v", outBoxEvent.ID, err)
			continue
		}
		outBoxEventsIDs = append(outBoxEventsIDs, outBoxEvent.ID)
	}

	if len(outBoxEventsIDs) == 0 {
		s.logger.Warn("no events were successfully published.")
		return nil
	}

	if err = s.repository.UpdatePublished(ctx, outBoxEventsIDs, time.Now()); err != nil {
		return fmt.Errorf("failed to update published status for events: %w", err)
	}

	s.logger.Infof("published events successfully with IDS: %v", outBoxEventsIDs)

	return nil
}
