package main

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/events"
	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/types"

	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/models"

	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/contract"

	"github.com/saeedjhn/go-backend-clean-arch/internal/adapter/jsonfilelogger"
	"github.com/saeedjhn/go-backend-clean-arch/internal/adapter/rmqpc"
	"github.com/saeedjhn/go-backend-clean-arch/internal/event"
)

const _eventBufferSize = 1024

func main() {
	logger := setupLogger()

	urTopic := models.EventType("user.registered")
	opTopic := models.EventType("order.processing")

	// Definination JOBS
	router := event.NewRouter()
	router.Register(urTopic, handleUserRegistered)
	router.Register(opTopic, handleOrderPlaced)

	cfg := rmqpc.Config{
		Connection: rmqpc.ConnectionConfig{
			Username:         "admin",
			Password:         "password123",
			Host:             "localhost",
			Port:             "5672",
			BaseRetryTimeout: 2 * time.Second,
			Multiplier:       2.0,
			MaxDelay:         10 * time.Second,
			MaxRetry:         5,
		},
		MQ: rmqpc.MQConfig{
			Exchange: rmqpc.ExchangeConfig{
				Name:       "my-exchange",
				Kind:       rmqpc.TopicExchange,
				Durable:    true,
				AutoDelete: false,
				Internal:   false,
				NoWait:     false,
				Args:       nil,
			},
			QueueBind: rmqpc.QueueBindConfig{
				Queue:            "test-queue",
				BindingKey:       []models.EventType{urTopic, opTopic},
				Durable:          true,
				AutoDelete:       false,
				Exclusive:        false,
				NoWait:           false,
				ArgsQueueDeclare: nil,
				ArgsQueueBind:    nil,
				PrefetchCount:    1,
				PrefetchSize:     0,
				PrefetchGlobal:   false,
			},
			Publish: rmqpc.PublishConfig{
				Mandatory: false,
				Immediate: false,
			},
			Consume: rmqpc.ConsumeConfig{
				AutoAck:   false,
				Exclusive: false,
				NoLocal:   false,
				NoWait:    false,
				Args:      nil,
			},
		},
	}

	rMQ := rmqpc.New("my-connection", cfg)

	if err := rMQ.Connect(); err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	if err := rMQ.SetupExchange(); err != nil {
		log.Fatalf("Failed to setup exchange: %v", err)
	}

	if err := rMQ.SetupBindQueue(); err != nil {
		log.Fatalf("Failed to setup queue binding: %v", err)
	}

	eventConsumer := event.NewEventConsumer(
		_eventBufferSize,
		router,
		rMQ,
	)
	eventConsumer.WithLogger(logger)

	// Start the contract consumer
	var wg sync.WaitGroup
	wg.Add(1)
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		defer wg.Done()
		eventConsumer.Start()
	}()

	e := events.NewUserRegisteredEvent(types.ID(1), "reason")
	b, _ := e.Marshal()

	evtUserRegistered := models.Event{Type: urTopic, Payload: b}
	evtOrderProcessing := models.Event{Type: opTopic, Payload: []byte("Order123456")}

	go func() {
		for i := range 3 {
			log.Printf("bus.Publish %d invoked\n", i)

			// Publish or insert contract to db & background run outbox with scheduler, 5s internal
			if err := rMQ.Publish(evtUserRegistered); err != nil {
				log.Fatalf("Failed to publish contract: %v", err)
			}

			if err := rMQ.Publish(evtOrderProcessing); err != nil {
				log.Fatalf("Failed to publish contract: %v", err)
			}

			time.Sleep(time.Second)
		}
	}()

	quit := make(chan bool)
	// quit := make(chan os.Signal, 1)
	// signal.Notify(quit, os.Interrupt)

	log.Println("Waiting for termination signal...")
	go func() {
		time.Sleep(5 * time.Second)
		quit <- true
	}()

	<-quit
	log.Println("Termination signal received, shutting down...")

	cancel()
	_ = eventConsumer.Shutdown(ctx)

	// Gracefully shut down the consumer
	wg.Wait()

	log.Println("Consumer shut down gracefully.")
	// Wait until graceful shutdown is complete
}

func handleUserRegistered(payload []byte) error {
	log.Printf("[Notification] Sending welcome email for user: %s\n", payload)

	var ur events.UserRegisteredEvent
	_ = ur.Unmarshal(payload)
	log.Printf("%#v", ur)

	return nil
}

func handleOrderPlaced(payload []byte) error {
	log.Printf("[Order] Processing order: %s\n", payload)
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
