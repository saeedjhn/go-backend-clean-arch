package rabbitmq_test

import (
	"log"
	"testing"
	"time"

	"github.com/saeedjhn/go-backend-clean-arch/internal/adaptor/rmqpc"
	"github.com/saeedjhn/go-backend-clean-arch/internal/event"
)

func TestExample(t *testing.T) {
	config := rmqpc.Config{
		Host:     "localhost",
		Port:     5672,
		Username: "admin",
		Password: "password123",
	}

	queue := "test_queue"
	topics := []event.Topic{"order.created", "order.updated"}

	adapter := rmqpc.New(config, queue, topics)

	eventStream := make(chan event.Event)

	go func() {
		err := adapter.Consume(eventStream)
		if err != nil {
			log.Fatalf("Error consuming messages: %v", err)
		}
	}()

	time.Sleep(2 * time.Second)
	adapter.Publish(event.Event{Topic: "order.created", Payload: []byte("New Order Created")})
	adapter.Publish(event.Event{Topic: "order.updated", Payload: []byte("Order Updated")})

	for evt := range eventStream {
		log.Printf("Received evt: Topic=%s, Payload=%s", evt.Topic, string(evt.Payload))
	}
}
