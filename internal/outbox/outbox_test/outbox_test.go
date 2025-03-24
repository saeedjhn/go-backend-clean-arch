package outbox_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/saeedjhn/go-backend-clean-arch/internal/adaptor/jsonfilelogger"
	"github.com/saeedjhn/go-backend-clean-arch/internal/contract"
	"github.com/saeedjhn/go-backend-clean-arch/internal/outbox"
	"github.com/saeedjhn/go-backend-clean-arch/internal/outbox/outbox_test/mocks"
	"github.com/saeedjhn/go-backend-clean-arch/internal/types"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/scheduler"
	"github.com/stretchr/testify/mock"
)

//go:generate go test -v -race -count=1 ./...

const _sleep = 2 * time.Second

func TestOutbox_ProcessOutBoxEvents_NoUnpublishedEvents_ReturnsNil(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	config := outbox.Config{
		Interval:       time.Second,
		BatchSize:      2,
		RetryThreshold: 3,
	}
	var events []outbox.Event

	logger := setupLogger()
	sch, err := setupScheduler()
	if err != nil {
		t.Fatalf("failed to set up scheduler: %v", err)
	}
	mockPublisher := mocks.NewMockPublisher(t)
	mockRepo := mocks.NewMockRepository(t)
	mockRepo.On("GetUnPublished", mock.Anything, 0, config.BatchSize, config.RetryThreshold).Return(events, nil).Maybe()

	ob := outbox.New(
		config, logger, sch, mockPublisher, mockRepo,
	)

	ob.StartProcessing(ctx)

	if err = sch.Start(); err != nil {
		t.Fatalf("failed to start scheduler: %v", err)
	}

	time.Sleep(_sleep)

	if err = sch.Shutdown(); err != nil {
		t.Fatalf("failed to shutdown scheduler: %v", err)
	}
}

func TestOutbox_ProcessOutBoxEvents_FailToPublishEvent_ReturnsError(t *testing.T) {
	ctx := context.Background()
	config := outbox.Config{
		Interval:       time.Second,
		BatchSize:      1,
		RetryThreshold: 1,
	}
	events := []outbox.Event{{
		ID:            1,
		Topic:         "user.sign.up",
		Payload:       []byte("user data"),
		IsPublished:   false,
		ReTriedCount:  0,
		LastRetriedAt: time.Time{},
		PublishedAt:   time.Time{},
	}}

	logger := setupLogger()
	sch, err := setupScheduler()
	if err != nil {
		t.Fatalf("failed to set up scheduler: %v", err)
	}
	mockPublisher := mocks.NewMockPublisher(t)
	mockPublisher.On("Publish", mock.Anything).Return(errors.New("publish error"))

	mockRepo := mocks.NewMockRepository(t)
	mockRepo.On("GetUnPublished", mock.Anything, 0, config.BatchSize, config.RetryThreshold).Return(events, nil)
	mockRepo.On("UpdatePublished", mock.Anything, []types.ID{events[0].ID}, mock.Anything).Return(nil).Maybe()

	ob := outbox.New(
		config, logger, sch, mockPublisher, mockRepo,
	)

	ob.StartProcessing(ctx)

	if err = sch.Start(); err != nil {
		t.Fatalf("failed to start scheduler: %v", err)
	}

	time.Sleep(_sleep)

	if err = sch.Shutdown(); err != nil {
		t.Fatalf("failed to shutdown scheduler: %v", err)
	}
}

func TestOutbox_ProcessOutBoxEvents_FailToUpdatePublished_ReturnsError(t *testing.T) {
	ctx := context.Background()
	config := outbox.Config{
		Interval:       time.Second,
		BatchSize:      2,
		RetryThreshold: 3,
	}
	events := []outbox.Event{{
		ID:            1,
		Topic:         "user.sign.up",
		Payload:       []byte("user data"),
		IsPublished:   false,
		ReTriedCount:  0,
		LastRetriedAt: time.Time{},
		PublishedAt:   time.Time{},
	}}

	logger := setupLogger()
	sch, err := setupScheduler()
	if err != nil {
		t.Fatalf("failed to set up scheduler: %v", err)
	}
	mockPublisher := mocks.NewMockPublisher(t)
	mockPublisher.EXPECT().Publish(mock.Anything).Return(nil)
	mockRepo := mocks.NewMockRepository(t)
	mockRepo.On("GetUnPublished", mock.Anything, 0, config.BatchSize, config.RetryThreshold).Return(events, nil)
	// m, _ := mockRepo.GetUnPublished(ctx, 0, config.BatchSize, config.RetryThreshold)
	mockRepo.On("UpdatePublished", mock.Anything, []types.ID{events[0].ID}, mock.Anything).
		Return(errors.New("update error"))

	ob := outbox.New(
		config, logger, sch, mockPublisher, mockRepo,
	)

	ob.StartProcessing(ctx)

	if err = sch.Start(); err != nil {
		t.Fatalf("failed to start scheduler: %v", err)
	}

	time.Sleep(_sleep)

	if err = sch.Shutdown(); err != nil {
		t.Fatalf("failed to shutdown scheduler: %v", err)
	}
}

func TestOutbox_ProcessOutBoxEvents_SuccessfullyPublished_ReturnsNil(t *testing.T) {
	ctx := context.Background()
	config := outbox.Config{
		Interval:       time.Second,
		BatchSize:      2,
		RetryThreshold: 3,
	}
	events := []outbox.Event{{
		ID:            1,
		Topic:         "user.sign.up",
		Payload:       []byte("user data"),
		IsPublished:   false,
		ReTriedCount:  0,
		LastRetriedAt: time.Time{},
		PublishedAt:   time.Time{},
	}, {
		ID:            2,
		Topic:         "order.created",
		Payload:       []byte("order details"),
		IsPublished:   false,
		ReTriedCount:  0,
		LastRetriedAt: time.Time{},
		PublishedAt:   time.Time{},
	}}

	logger := setupLogger()
	sch, err := setupScheduler()
	if err != nil {
		t.Fatalf("failed to set up scheduler: %v", err)
	}
	mockPublisher := mocks.NewMockPublisher(t)
	mockPublisher.EXPECT().Publish(mock.Anything).Return(nil)
	mockRepo := mocks.NewMockRepository(t)
	mockRepo.On("GetUnPublished", mock.Anything, 0, config.BatchSize, config.RetryThreshold).Return(events, nil)
	mockRepo.On("UpdatePublished", mock.Anything, []types.ID{events[0].ID, events[1].ID}, mock.Anything).Return(nil)

	ob := outbox.New(
		config, logger, sch, mockPublisher, mockRepo,
	)
	ob.StartProcessing(ctx)

	if err = sch.Start(); err != nil {
		t.Fatalf("failed to start scheduler: %v", err)
	}

	time.Sleep(_sleep)

	if err = sch.Shutdown(); err != nil {
		t.Fatalf("failed to shutdown scheduler: %v", err)
	}
}

func setupLogger() contract.Logger {
	config := jsonfilelogger.Config{
		LocalTime:        true,
		Console:          true,
		EnableCaller:     true,
		EnableStacktrace: true,
		Level:            "debug",
	}

	return jsonfilelogger.New(jsonfilelogger.NewDevelopmentStrategy(config)).Configure()
}

func setupScheduler() (*scheduler.Scheduler, error) {
	sch := scheduler.New()
	if err := sch.Configure(); err != nil {
		return nil, err
	}

	return sch, nil
}
