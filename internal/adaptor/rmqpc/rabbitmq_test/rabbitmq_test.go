package rabbitmq_test

import (
	"testing"
	"time"

	"github.com/saeedjhn/go-backend-clean-arch/internal/entity"

	"github.com/saeedjhn/go-backend-clean-arch/internal/adaptor/rmqpc"
)

//go:generate go test -v -race -count=1 ./...

func TestConsumeMessage_WhenMessageIsPublished_ShouldReceiveMessagee(t *testing.T) {
	urTopic := entity.Topic("user.registered")

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
				BindingKey:       []entity.Topic{urTopic},
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
		t.Fatalf("Failed to connect: %v", err)
	}

	if err := rMQ.SetupExchange(); err != nil {
		t.Fatalf("Failed to setup exchange: %v", err)
	}

	if err := rMQ.SetupBindQueue(); err != nil {
		t.Fatalf("Failed to setup queue binding: %v", err)
	}

	eventStream := make(chan entity.Event)

	go func() {
		err := rMQ.Consume(eventStream)
		if err != nil {
			t.Errorf("Error consuming messages: %v", err)
			return
		}
	}()

	time.Sleep(2 * time.Second)
	_ = rMQ.Publish(entity.Event{Topic: "user.registered", Payload: []byte("New User Registered")})

	// 	for evt := range eventStream {
	// 		t.Logf("Received evt: Topic=%s, Payload=%s", evt.Topic, string(evt.Payload))
	// 	}

	for {
		select {
		case evt := <-eventStream:
			t.Logf("Received evt: Topic=%s, Payload=%s", evt.Topic, string(evt.Payload))
			return
		case <-time.After(5 * time.Second):
			t.Fatalf("Test timed out waiting for message")
			return
		}
	}
}
