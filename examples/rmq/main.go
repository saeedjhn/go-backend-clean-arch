package main

import (
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/saeedjhn/go-backend-clean-arch/internal/adaptor/rmqpc"

	"github.com/saeedjhn/go-backend-clean-arch/internal/event"
)

func main() {
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
				Queue:            "my-queue",
				BindingKey:       []string{"my-routing-key"},
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
				RoutingKey: "my-routing-key",
				Mandatory:  false,
				Immediate:  false,
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

	conn := rmqpc.New("my-connection", cfg)

	if err := conn.Connect(); err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	if err := conn.SetupExchange(); err != nil {
		log.Fatalf("Failed to setup exchange: %v", err)
	}

	if err := conn.SetupBindQueue(); err != nil {
		log.Fatalf("Failed to setup queue binding: %v", err)
	}

	eventStream := make(chan event.Event)

	go func() {
		if err := conn.Consume(eventStream); err != nil {
			log.Fatalf("Failed to consume messages: %v", err)
		}
	}()

	eventMsg := event.Event{
		Topic:   "my-routing-key",
		Payload: []byte("Hello, RabbitMQ!"),
	}

	if err := conn.Publish(eventMsg); err != nil {
		log.Fatalf("Failed to publish message: %v", err)
	}
	log.Println("Message Published!")

	go func() {
		for evt := range eventStream {
			log.Printf("Received event: %s - %s", evt.Topic, string(evt.Payload))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	log.Println("Waiting for termination signal...")
	<-quit
	log.Println("Termination signal received, shutting down...")
}
