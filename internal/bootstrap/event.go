package bootstrap

import (
	"context"
	"fmt"

	"github.com/saeedjhn/go-backend-clean-arch/internal/adaptor/rmqpc"
	"github.com/saeedjhn/go-backend-clean-arch/internal/contract"
	"github.com/saeedjhn/go-backend-clean-arch/internal/entity"
	"github.com/saeedjhn/go-backend-clean-arch/internal/event"
	"github.com/saeedjhn/go-backend-clean-arch/internal/job"
)

const (
	_rabbitMQConnection = "my_connection"
	_eventBufferSize    = 1024

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

var _routerRegister = map[entity.Topic]entity.RouterHandler{ //nolint:gochecknoglobals // nothing
	entity.UserRegisteredTopic: job.UserRegisteredHandler,
}

func NewEvent(
	ctx context.Context,
	configRabbitmq rmqpc.ConnectionConfig,
	logger contract.Logger,
) error {
	router := event.NewRouter()

	for t, h := range _routerRegister {
		router.Register(t, h)
	}

	rMQ, err := setupRabbitMQ(configRabbitmq)
	if err != nil {
		return fmt.Errorf("failed rabbitmq setup (host: %s, port: %s): %w",
			configRabbitmq.Host, configRabbitmq.Port, err)
	}

	eventConsumer := event.NewEventConsumer(
		_eventBufferSize,
		router,
		rMQ,
	)
	eventConsumer.WithLogger(logger)

	eventConsumer.Start(ctx)

	return nil
}

func setupRabbitMQ(config rmqpc.ConnectionConfig) (*rmqpc.Connection, error) {
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
				BindingKey:     ListBindingKeys(),
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
		return nil, fmt.Errorf("failed to setup exchange: %w", err)
	}

	if err := rMQ.SetupBindQueue(); err != nil {
		return nil, fmt.Errorf("failed to setup queue binding: %w", err)
	}

	return rMQ, nil
}

func ListBindingKeys() []entity.Topic {
	var bindingkeys []entity.Topic

	for t := range _routerRegister {
		bindingkeys = append(bindingkeys, t)
	}

	return bindingkeys
}
