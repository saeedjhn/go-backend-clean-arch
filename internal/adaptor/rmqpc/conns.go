package rmqpc

import (
	"errors"
	"fmt"
	"net"
	"time"

	"github.com/saeedjhn/go-backend-clean-arch/internal/contract"

	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
)

var _connectionPool = make(map[string]*Connection) //nolint:gochecknoglobals // nothing

type Connection struct {
	Name       string
	config     Config
	connection *amqp.Connection
	channel    *amqp.Channel
	errChan    chan error
}

func GetConnection(name string) *Connection {
	return _connectionPool[name]
}

func Connections() map[string]*Connection { return _connectionPool }

func New(
	connectionName string,
	config Config,
) *Connection {
	if c, ok := _connectionPool[connectionName]; ok {
		return c
	}

	c := &Connection{
		Name:    connectionName,
		config:  config,
		errChan: make(chan error),
	}

	_connectionPool[connectionName] = c

	return c
}

func (c *Connection) Config() Config {
	return c.config
}

func (c *Connection) Queue() string {
	return c.config.MQ.QueueBind.Queue
}

func (c *Connection) Channel() *amqp.Channel {
	return c.channel
}

func (c *Connection) ConnectRaw() error {
	var err error

	uri := fmt.Sprintf("amqp://%s:%s@%s",
		c.config.Connection.Username,
		c.config.Connection.Password,
		net.JoinHostPort(c.config.Connection.Host, c.config.Connection.Port),
	)

	c.connection, err = amqp.Dial(uri)
	if err != nil {
		return fmt.Errorf(
			"connection error: failed to connect to RabbitMQ server at %s: %w", uri, err)
	}

	return nil
}

func (c *Connection) Connect() error {
	var err error

	uri := fmt.Sprintf("amqp://%s:%s@%s",
		c.config.Connection.Username,
		c.config.Connection.Password,
		net.JoinHostPort(c.config.Connection.Host, c.config.Connection.Port),
	)

	c.connection, err = amqp.Dial(uri)
	if err != nil {
		return fmt.Errorf(
			"connection error: failed to connect to RabbitMQ server at %s: %w", uri, err)
	}

	go func() {
		<-c.connection.NotifyClose(make(chan *amqp.Error)) // Listen to NotifyClose
		c.errChan <- errors.New("Connection.Closed")
	}()

	c.channel, err = c.connection.Channel()
	if err != nil {
		return fmt.Errorf("channel error: failed to create channel after connection: %w", err)
	}

	return nil
}

func (c *Connection) SetupExchange() error {
	if err := c.channel.ExchangeDeclare(
		c.config.MQ.Exchange.Name,
		c.config.MQ.Exchange.Kind.String(),
		c.config.MQ.Exchange.Durable,
		c.config.MQ.Exchange.AutoDelete,
		c.config.MQ.Exchange.Internal,
		c.config.MQ.Exchange.NoWait,
		amqp.Table(c.config.MQ.Exchange.Args),
	); err != nil {
		return fmt.Errorf("exchange setup error: failed to declare exchange '%s' of type '%s': %w",
			c.config.MQ.Exchange.Name,
			c.config.MQ.Exchange.Kind.String(),
			err,
		)
	}

	return nil
}

func (c *Connection) SetupBindQueue() error {
	q, err := c.channel.QueueDeclare(
		c.config.MQ.QueueBind.Queue,
		c.config.MQ.QueueBind.Durable,
		c.config.MQ.QueueBind.AutoDelete,
		c.config.MQ.QueueBind.Exclusive,
		c.config.MQ.QueueBind.NoWait,
		amqp.Table(c.config.MQ.QueueBind.ArgsQueueDeclare),
	)
	if err != nil {
		return fmt.Errorf("failed to declare queue '%s': %w", c.config.MQ.QueueBind.Queue, err)
	}

	for _, bindingKey := range c.config.MQ.QueueBind.BindingKey {
		if err = c.channel.QueueBind(
			q.Name,
			bindingKey.String(), // binding key
			c.config.MQ.Exchange.Name,
			c.config.MQ.QueueBind.NoWait,
			amqp.Table(c.config.MQ.QueueBind.ArgsQueueBind),
		); err != nil {
			return fmt.Errorf(
				"failed to bind queue '%s' to exchange '%s': %w",
				q.Name,
				c.config.MQ.Exchange.Name,
				err,
			)
		}
	}

	if err = c.channel.Qos(
		c.config.MQ.QueueBind.PrefetchCount,
		c.config.MQ.QueueBind.PrefetchSize,
		c.config.MQ.QueueBind.PrefetchGlobal,
	); err != nil {
		return fmt.Errorf("failed to set QoS for queue '%s': %w", q.Name, err)
	}

	return nil
}

func (c *Connection) Publish(evt contract.Event) error {
	// non-blocking channel - if there is no error will go to default where we do nothing
	select {
	case err := <-c.errChan:
		if err != nil {
			return c.reconnect()
		}
	default:
	}

	if len(c.config.MQ.Exchange.Name) == 0 {
		return errors.New("publish failed: no exchange selected; specify an exchange in the configuration")
	}

	if err := c.channel.Publish(
		c.config.MQ.Exchange.Name,
		string(evt.Topic), // routing key
		c.config.MQ.Publish.Mandatory,
		c.config.MQ.Publish.Immediate,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        evt.Payload,
		},
	); err != nil {
		return fmt.Errorf(
			"publish failed: message delivery to exchange '%s' with routing key '%s' unsuccessful: %w",
			c.config.MQ.Exchange.Name,
			evt.Topic,
			err,
		)
	}

	return nil
}

func (c *Connection) Consume(eventStream chan<- contract.Event) error {
	var (
		deliveries <-chan amqp.Delivery
		err        error
	)
	if c.isEmptyQueue() {
		deliveries, err = c.consumeFromTemporaryQueue(c.config.MQ.Consume)
		if err != nil {
			return fmt.Errorf("failed to consume from temporary queue: %w", err)
		}
	} else {
		deliveries, err = c.consumeFromDefinedQueues(c.config.MQ.Consume)
		if err != nil {
			return fmt.Errorf("failed to consume from temporary queue: %w", err)
		}
	}

	for delivery := range deliveries {
		eventStream <- contract.Event{Topic: contract.Topic(delivery.RoutingKey), Payload: delivery.Body}
		if !c.config.MQ.Consume.AutoAck {
			if err = delivery.Ack(false); err != nil {
				return fmt.Errorf("failed to ACK message: %w", err)
			}
		}
	}

	return nil
}

func (c *Connection) Close() error {
	return c.connection.Close()
}

func (c *Connection) reconnect() error {
	var (
		retryTimeout = c.config.Connection.BaseRetryTimeout

		connErr error
		exErr   error
		qbErr   error
	)

	for range c.config.Connection.MaxRetry {
		if connErr = c.Connect(); connErr == nil {
			return nil
		}

		if exErr = c.SetupExchange(); exErr == nil {
			return nil
		}

		if qbErr = c.SetupBindQueue(); qbErr == nil {
			return nil
		}

		// log.Printf("retryTimeout: %v", retryTimeout)
		retryTimeout = time.Duration(float64(retryTimeout) * c.config.Connection.Multiplier)

		time.Sleep(retryTimeout)

		if retryTimeout > c.config.Connection.MaxDelay {
			retryTimeout = c.config.Connection.BaseRetryTimeout
		}
	}

	return fmt.Errorf(
		"reconnect failed after %d attempts: %w, exchange/queue setup error: %w",
		c.config.Connection.MaxRetry,
		connErr,
		exErr,
	)
}

func (c *Connection) isEmptyQueue() bool {
	return len(c.config.MQ.QueueBind.Queue) == 0
}

func buildConsumeTag(queueName string) string {
	return fmt.Sprintf("%s-Tag", queueName)
}

func (c *Connection) consumeFromTemporaryQueue(cfg ConsumeConfig) (<-chan amqp.Delivery, error) {
	uuID := uuid.New().String()
	uuIDTag := buildConsumeTag(uuID)

	// log.Printf(
	// 	"StartConsume.Starting.To.Consume.From.Queue, ConsumerTag: %v",
	// 	uuIDTag,
	// )

	consume, err := c.channel.Consume(
		uuID,
		uuIDTag,
		cfg.AutoAck,
		cfg.Exclusive,
		cfg.NoLocal,
		cfg.NoWait,
		amqp.Table(cfg.Args),
	)

	if err != nil {
		return nil, fmt.Errorf(
			"failed to consume from temporary queue with ConsumerTag '%s': %w",
			uuIDTag,
			err,
		)
	}

	return consume, nil
}

func (c *Connection) consumeFromDefinedQueues(cfg ConsumeConfig) (<-chan amqp.Delivery, error) {
	qTag := buildConsumeTag(c.config.MQ.QueueBind.Queue)

	// log.Printf(
	// 	"StartConsume.Starting.To.Consume.From.Queue, ConsumerTag: %v",
	// 	buildConsumeTag(c.config.MQ.QueueBind.Queue),
	// )

	consume, err := c.channel.Consume(
		c.config.MQ.QueueBind.Queue,
		qTag,
		cfg.AutoAck,
		cfg.Exclusive,
		cfg.NoLocal,
		cfg.NoWait,
		amqp.Table(cfg.Args),
	)

	if err != nil {
		return nil, fmt.Errorf(
			"failed to consume from defined queue with ConsumerTag '%s': %w",
			qTag,
			err,
		)
	}

	return consume, nil
}
