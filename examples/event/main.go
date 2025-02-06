package main

import (
	"context"
	"github.com/saeedjhn/go-backend-clean-arch/internal/adaptor/jsonfilelogger"
	"github.com/saeedjhn/go-backend-clean-arch/internal/contract"
	"log"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/saeedjhn/go-backend-clean-arch/internal/event"
)

const _eventBufferSize = 1024

type InMemoryBus struct {
	eventStream chan event.Event
}

func NewInMemoryBus() *InMemoryBus {
	return &InMemoryBus{eventStream: make(chan event.Event, _eventBufferSize)}
}

func (b *InMemoryBus) Publish(event event.Event) error {
	b.eventStream <- event

	return nil
}

func (b *InMemoryBus) Consume(ch chan<- event.Event) error {
	go func() {
		for evt := range b.eventStream {
			ch <- evt
		}
	}()
	return nil
}

func handleUserRegistered(event event.Event) error {
	log.Printf("[Notification] Sending welcome email for user: %s\n", string(event.Payload))
	return nil
}

func handleOrderPlaced(event event.Event) error {
	log.Printf("[Order] Processing order: %s\n", string(event.Payload))
	return nil
}

func main() {
	logger := setupLogger()

	urTopic := event.Topic("user.registered")
	opTopic := event.Topic("order.processing")

	router := event.NewRouter()
	router.Register(urTopic, handleUserRegistered)
	router.Register(opTopic, handleOrderPlaced)

	bus := NewInMemoryBus()

	eventConsumer := event.NewEventConsumer(
		_eventBufferSize,
		router,
		bus,
	)
	eventConsumer.WithLogger(logger)

	// Start the event consumer
	var wg sync.WaitGroup
	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()
	// Or
	ctx, cancel := context.WithCancel(context.Background())

	go eventConsumer.Start(ctx, &wg)

	evtUserRegistered := event.Event{Topic: urTopic, Payload: []byte("User123")}
	evtOrderProcessing := event.Event{Topic: opTopic, Payload: []byte("Order123456")}

	go func() {
		for i := range 3 {
			log.Printf("bus.Publish %d invoked\n", i)

			if err := bus.Publish(evtUserRegistered); err != nil {
				log.Fatalf("Failed to publish event: %v", err)
			}

			if err := bus.Publish(evtOrderProcessing); err != nil {
				log.Fatalf("Failed to publish event: %v", err)
			}

			time.Sleep(time.Second)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	log.Println("Waiting for termination signal...")
	<-quit
	log.Println("Termination signal received, shutting down...")

	cancel()

	// Gracefully shut down the consumer
	wg.Wait()
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
