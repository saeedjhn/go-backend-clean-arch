package main

import (
	"log"
	"os"
	"os/signal"
	"time"

	rmqarc2 "github.com/saeedjhn/go-backend-clean-arch/internal/adaptor/rmqpc"

	"github.com/saeedjhn/go-backend-clean-arch/internal/event"
)

func main() {
	cfg := rmqarc2.Config{
		Connection: rmqarc2.ConnectionConfig{
			Username:         "admin",
			Password:         "password123",
			Host:             "localhost",
			Port:             "5672",
			BaseRetryTimeout: 2 * time.Second,
			Multiplier:       2.0,
			MaxDelay:         10 * time.Second,
			MaxRetry:         5,
		},
		MQ: rmqarc2.MQConfig{
			Exchange: rmqarc2.ExchangeConfig{
				Name:       "my-exchange",
				Kind:       rmqarc2.TopicExchange,
				Durable:    true,
				AutoDelete: false,
				Internal:   false,
				NoWait:     false,
				Args:       nil,
			},
			QueueBind: rmqarc2.QueueBindConfig{
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
			Publish: rmqarc2.PublishConfig{
				RoutingKey: "my-routing-key",
				Mandatory:  false,
				Immediate:  false,
			},
			Consume: rmqarc2.ConsumeConfig{
				AutoAck:   false,
				Exclusive: false,
				NoLocal:   false,
				NoWait:    false,
				Args:      nil,
			},
		},
	}

	conn := rmqarc2.New("my-connection", cfg)

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
