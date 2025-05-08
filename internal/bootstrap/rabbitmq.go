package bootstrap

import (
	"fmt"

	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/models"

	"github.com/saeedjhn/go-backend-clean-arch/internal/event"

	"github.com/saeedjhn/go-backend-clean-arch/internal/adapter/rmqpc"
)

const (
	_rabbitMQConnection = "my_connection"

	ExchangeName       = "my-exchange"
	ExchangeKind       = "topic"
	ExchangeDurable    = true
	ExchangeAutoDelete = false
	ExchangeInternal   = false
	ExchangeNoWait     = false

	Queue               = "test-queue"
	QueueDurable        = true
	QueueAutoDelete     = false
	QueueExclusive      = false
	QueueNoWait         = false
	QueuePrefetchCount  = 1
	QueuePrefetchSize   = 0
	QueuePrefetchGlobal = false

	PublishMandatory = false
	PublishImmediate = false

	ConsumeAutoAck   = false
	ConsumeExclusive = false
	ConsumeNoLocal   = false
	ConsumeNoWait    = false
)

func NewRabbitmq(
	config rmqpc.ConnectionConfig,
	eventRouter map[models.EventType]event.RouterHandler,
) (*rmqpc.Connection, error) {
	cfg := rmqpc.Config{
		Connection: rmqpc.ConnectionConfig{
			Username:         config.Username,
			Password:         config.Password,
			Host:             config.Host,
			Port:             config.Port,
			BaseRetryTimeout: config.BaseRetryTimeout,
			Multiplier:       config.Multiplier,
			MaxDelay:         config.MaxDelay,
			MaxRetry:         config.MaxRetry,
		},
		MQ: rmqpc.MQConfig{
			Exchange: rmqpc.ExchangeConfig{
				Name:       ExchangeName,
				Kind:       ExchangeKind,
				Durable:    ExchangeDurable,
				AutoDelete: ExchangeAutoDelete,
				Internal:   ExchangeInternal,
				NoWait:     ExchangeNoWait,
			},
			QueueBind: rmqpc.QueueBindConfig{
				Queue:          Queue,
				BindingKey:     bindingKeys(eventRouter),
				Durable:        QueueDurable,
				AutoDelete:     QueueAutoDelete,
				Exclusive:      QueueExclusive,
				NoWait:         QueueNoWait,
				PrefetchCount:  QueuePrefetchCount,
				PrefetchSize:   QueuePrefetchSize,
				PrefetchGlobal: QueuePrefetchGlobal,
			},
			Publish: rmqpc.PublishConfig{
				Mandatory: PublishMandatory,
				Immediate: PublishImmediate,
			},
			Consume: rmqpc.ConsumeConfig{
				AutoAck:   ConsumeAutoAck,
				Exclusive: ConsumeExclusive,
				NoLocal:   ConsumeNoLocal,
				NoWait:    ConsumeNoWait,
			},
		},
	}

	rMQ := rmqpc.New(_rabbitMQConnection, cfg)

	if err := rMQ.Connect(); err != nil {
		return nil, fmt.Errorf("failed to connect: %w", err)
	}

	if err := rMQ.SetupExchange(); err != nil {
		return nil, fmt.Errorf("failed to SetupRabbitMQ exchange: %w", err)
	}

	if err := rMQ.SetupBindQueue(); err != nil {
		return nil, fmt.Errorf("failed to SetupRabbitMQ queue binding: %w", err)
	}

	return rMQ, nil
}

func bindingKeys(eventRouter map[models.EventType]event.RouterHandler) []models.EventType {
	var bKeys []models.EventType

	for t := range eventRouter {
		bKeys = append(bKeys, t)
	}

	return bKeys
}
