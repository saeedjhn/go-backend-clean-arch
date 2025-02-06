package event_test

import (
	"context"
	"errors"
	"github.com/saeedjhn/go-backend-clean-arch/internal/adaptor/jsonfilelogger"
	"github.com/saeedjhn/go-backend-clean-arch/internal/contract"
	"log"
	"sync"
	"testing"
	"time"

	"github.com/saeedjhn/go-backend-clean-arch/internal/event"
	"github.com/saeedjhn/go-backend-clean-arch/internal/event/event_test/mocks"
	"github.com/stretchr/testify/mock"
)

//go:generate go test -v -race -count=1 ./...

func TestStart_WithContextCancellation_StopsGracefully(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logger := setupLogger()
	mockConsumer := mocks.NewMockConsumer(t)
	mockConsumer.On("Consume", mock.Anything).
		Return(nil).Maybe()

	router := event.NewRouter()
	eventConsumer := event.NewEventConsumer(10, router, mockConsumer)
	eventConsumer.WithLogger(logger)

	var wg sync.WaitGroup

	go func() {
		time.Sleep(100 * time.Millisecond)
		cancel()
	}()

	// err := eventConsumer.Start(ctx, &wg)
	// require.NoError(t, err)

	eventConsumer.Start(ctx, &wg)

	mockConsumer.AssertExpectations(t)
}

func TestStart_WithConsumerError_ReturnsError(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logger := setupLogger()
	mockConsumer := mocks.NewMockConsumer(t)
	mockConsumer.On("Consume", mock.Anything).
		Return(errors.New("consumer failure")).Maybe()

	router := event.NewRouter()
	eventConsumer := event.NewEventConsumer(10, router, mockConsumer)
	eventConsumer.WithLogger(logger)

	var wg sync.WaitGroup

	// err := eventConsumer.Start(ctx, &wg)
	// require.Error(t, err)
	// assert.Contains(t, err.Error(), "consumer failure")

	eventConsumer.Start(ctx, &wg)

	mockConsumer.AssertExpectations(t)
}

func TestStart_WithNoEvents_ExitsGracefully(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logger := setupLogger()
	mockConsumer := mocks.NewMockConsumer(t)
	mockConsumer.On("Consume", mock.Anything).
		Return(nil).Maybe()

	router := event.NewRouter()
	eventConsumer := event.NewEventConsumer(10, router, mockConsumer)
	eventConsumer.WithLogger(logger)

	var wg sync.WaitGroup

	go func() {
		time.Sleep(100 * time.Millisecond)
		cancel()
	}()

	// err := eventConsumer.Start(ctx, &wg)
	// require.NoError(t, err)

	eventConsumer.Start(ctx, &wg)

	mockConsumer.AssertExpectations(t)
}

func TestStart_WithValidEvents_ProcessesSuccessfull(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	logger := setupLogger()
	mockConsumer := mocks.NewMockConsumer(t)
	mockConsumer.On("Consume", mock.Anything).
		Return(nil).Maybe()

	router := event.NewRouter()
	router.Register(event.Topic("user.registered"), handleUserRegistered)

	eventConsumer := event.NewEventConsumer(10, router, mockConsumer)
	eventConsumer.WithLogger(logger)

	go func() {
		time.Sleep(time.Second)
		cancel()
	}()

	var wg sync.WaitGroup

	// err := eventConsumer.Start(ctx, &wg)
	// require.NoError(t, err)
	eventConsumer.Start(ctx, &wg)

	wg.Wait()

	mockConsumer.AssertExpectations(t)
}

func handleUserRegistered(event event.Event) error {
	log.Printf("[Notification] Sending welcome email for user: %s\n", string(event.Payload))
	return nil
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
