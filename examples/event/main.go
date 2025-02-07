package main

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/saeedjhn/go-backend-clean-arch/internal/entity"

	"github.com/saeedjhn/go-backend-clean-arch/internal/adaptor/rmqpc"

	"github.com/saeedjhn/go-backend-clean-arch/internal/adaptor/jsonfilelogger"
	"github.com/saeedjhn/go-backend-clean-arch/internal/contract"

	"github.com/saeedjhn/go-backend-clean-arch/internal/event"
)

const _eventBufferSize = 1024

func main() {
	logger := setupLogger()

	urTopic := entity.Topic("user.registered")
	opTopic := entity.Topic("order.processing")

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
				BindingKey:       []entity.Topic{urTopic, opTopic},
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

	// Start the event consumer
	var wg sync.WaitGroup
	wg.Add(1)
	// Or
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		defer wg.Done()
		eventConsumer.Start(ctx)
	}()

	evtUserRegistered := entity.Event{Topic: urTopic, Payload: []byte("User123")}
	evtOrderProcessing := entity.Event{Topic: opTopic, Payload: []byte("Order123456")}

	go func() {
		for i := range 3 {
			log.Printf("bus.Publish %d invoked\n", i)

			// Publish or insert event to db & background run outbox with scheduler, 5s internal
			if err := rMQ.Publish(evtUserRegistered); err != nil {
				log.Fatalf("Failed to publish event: %v", err)
			}

			if err := rMQ.Publish(evtOrderProcessing); err != nil {
				log.Fatalf("Failed to publish event: %v", err)
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
	// Gracefully shut down the consumer
	wg.Wait()

	log.Println("Consumer shut down gracefully.")
	// Wait until graceful shutdown is complete
}

func handleUserRegistered(event entity.Event) error {
	log.Printf("[Notification] Sending welcome email for user: %s\n", string(event.Payload))
	return nil
}

func handleOrderPlaced(event entity.Event) error {
	log.Printf("[Order] Processing order: %s\n", string(event.Payload))
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

// func main() {
// 	logger := setupLogger()
//
// 	urTopic := event.Topic("user.registered")
// 	opTopic := event.Topic("order.processing")
//
// 	router := event.NewRouter()
// 	router.Register(urTopic, handleUserRegistered)
// 	router.Register(opTopic, handleOrderPlaced)
//
// 	bus := NewInMemoryBus()
//
// 	eventConsumer := event.NewEventConsumer(
// 		_eventBufferSize,
// 		router,
// 		bus,
// 	)
// 	eventConsumer.WithLogger(logger)
//
// 	// Start the event consumer
// 	var wg sync.WaitGroup
// 	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	// defer cancel()
// 	// Or
// 	ctx, cancel := context.WithCancel(context.Background())
//
// 	go eventConsumer.Start(ctx, &wg)
//
// 	evtUserRegistered := event.Event{Topic: urTopic, Payload: []byte("User123")}
// 	evtOrderProcessing := event.Event{Topic: opTopic, Payload: []byte("Order123456")}
//
// 	go func() {
// 		for i := range 3 {
// 			log.Printf("bus.Publish %d invoked\n", i)
//
// 			if err := bus.Publish(evtUserRegistered); err != nil {
// 				log.Fatalf("Failed to publish event: %v", err)
// 			}
//
// 			if err := bus.Publish(evtOrderProcessing); err != nil {
// 				log.Fatalf("Failed to publish event: %v", err)
// 			}
//
// 			time.Sleep(time.Second)
// 		}
// 	}()
//
// 	quit := make(chan os.Signal, 1)
// 	signal.Notify(quit, os.Interrupt)
//
// 	log.Println("Waiting for termination signal...")
// 	<-quit
// 	log.Println("Termination signal received, shutting down...")
//
// 	cancel()
//
// 	// Gracefully shut down the consumer
// 	wg.Wait()
// }
