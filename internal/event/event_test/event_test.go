package event_test

import (
	"context"
	"errors"
	"log"
	"sync"
	"testing"
	"time"

	"github.com/saeedjhn/go-backend-clean-arch/internal/event"
	"github.com/saeedjhn/go-backend-clean-arch/internal/event/event_test/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

//go:generate go test -v -race -count=1 ./...

func TestStart_WithContextCancellation_StopsGracefully(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mockConsumer := mocks.NewMockConsumer(t)
	mockConsumer.On("Consume", mock.Anything).
		Return(nil).Maybe()

	router := event.NewRouter()
	eventConsumer := event.NewEventConsumer(10, router, mockConsumer)

	var wg sync.WaitGroup

	go func() {
		time.Sleep(100 * time.Millisecond)
		cancel()
	}()

	err := eventConsumer.Start(ctx, &wg)

	require.NoError(t, err)
	mockConsumer.AssertExpectations(t)
}

func TestStart_WithConsumerError_ReturnsError(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mockConsumer := mocks.NewMockConsumer(t)
	mockConsumer.On("Consume", mock.Anything).Return(errors.New("consumer failure"))

	router := event.NewRouter()
	eventConsumer := event.NewEventConsumer(10, router, mockConsumer)

	var wg sync.WaitGroup

	err := eventConsumer.Start(ctx, &wg)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "consumer failure")
	mockConsumer.AssertExpectations(t)
}

func TestStart_WithNoEvents_ExitsGracefully(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mockConsumer := mocks.NewMockConsumer(t)
	mockConsumer.On("Consume", mock.Anything).
		Return(nil).Maybe()

	router := event.NewRouter()
	eventConsumer := event.NewEventConsumer(10, router, mockConsumer)

	var wg sync.WaitGroup

	go func() {
		time.Sleep(100 * time.Millisecond)
		cancel()
	}()

	err := eventConsumer.Start(ctx, &wg)

	require.NoError(t, err)
	mockConsumer.AssertExpectations(t)
}

func TestStart_WithValidEvents_ProcessesSuccessfull(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	mockConsumer := mocks.NewMockConsumer(t)
	mockConsumer.On("Consume", mock.Anything).
		Return(nil).Maybe()

	router := event.NewRouter()
	router.Register(event.Topic("user.registered"), handleUserRegistered)

	eventConsumer := event.NewEventConsumer(10, router, mockConsumer)

	var wg sync.WaitGroup
	go func() {
		time.Sleep(time.Second)
		cancel()
	}()

	err := eventConsumer.Start(ctx, &wg)

	require.NoError(t, err)
	mockConsumer.AssertExpectations(t)
}

func handleUserRegistered(event event.Event) error {
	log.Printf("[Notification] Sending welcome email for user: %s\n", string(event.Payload))
	return nil
}
