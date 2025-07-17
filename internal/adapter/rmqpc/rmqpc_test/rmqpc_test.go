package rmqpc_test

import (
	"testing"
	"time"

	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/types"

	"github.com/saeedjhn/go-backend-clean-arch/internal/adapter/rmqpc"
)

//go:generate go test -v -race -count=1 ./...

func TestConsumeMessage_WhenMessageIsPublished_ShouldReceiveMessagee(t *testing.T) {
	urTopic := types.Event("user.registered")

	cfg := rmqpc.Config{
		Connection: _myRabbitMQConnectionConfig,
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
				BindingKey:       []types.Event{urTopic},
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

	eventStream := make(chan types.EventStream)

	go func() {
		err := rMQ.Consume(eventStream)
		if err != nil {
			t.Errorf("Error consuming messages: %v", err)
			return
		}
	}()

	time.Sleep(2 * time.Second)
	_ = rMQ.Publish(types.EventStream{Type: "user.registered", Payload: []byte("New User RegisteredHandler")})

	// 	for evt := range eventStream {
	// 		t.Logf("Received evt: Event=%s, Payload=%s", evt.Event, string(evt.Payload))
	// 	}

	for {
		select {
		case evt := <-eventStream:
			t.Logf("Received evt: Event=%s, Payload=%s", evt.Type, string(evt.Payload))
			return
		case <-time.After(5 * time.Second):
			t.Fatalf("Test timed out waiting for message")
			return
		}
	}
}
