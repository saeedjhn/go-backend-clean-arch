package rmqpc

import (
	"time"

	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/types"
)

type Table map[string]interface{}

type Config struct {
	Connection ConnectionConfig
	MQ         MQConfig
}

// ConnectionConfig defines the configuration needed to establish a connection to RabbitMQ.
type ConnectionConfig struct {
	Host                string        `mapstructure:"host"`
	Username            string        `mapstructure:"username"`
	Password            string        `mapstructure:"password"`
	Port                string        `mapstructure:"port"`
	Multiplier          float64       `mapstructure:"multiplier"`
	MaxRetry            int           `mapstructure:"max_retry"`
	BaseRetryTimeout    time.Duration `mapstructure:"base_retry_timeout"`
	MaxDelay            time.Duration `mapstructure:"max_delay"`
	CheckConnectionTime time.Duration `mapstructure:"check_connection_time"`
}

type MQConfig struct {
	Exchange  ExchangeConfig
	QueueBind QueueBindConfig
	Publish   PublishConfig
	Consume   ConsumeConfig
}

// ExchangeConfig defines the configuration needed to declare an exchange in RabbitMQ.
type ExchangeConfig struct {
	Name       string       // Name of the exchange to declare.
	Kind       ExchangeType // Type of the exchange: "direct", "fanout", "topic", or "headers".
	Durable    bool         // Whether the exchange should survive server restarts (true for durable exchanges).
	AutoDelete bool         // Whether the exchange should be automatically deleted when no longer in use.
	Internal   bool         // Whether the exchange is internal (only used for routing between exchanges).
	NoWait     bool         // Whether the server should wait for a response before declaring the exchange.
	Args       Table        // Additional arguments for exchange declaration.
}

// QueueBindConfig defines the configuration needed to declare and bind a queue in RabbitMQ.
type QueueBindConfig struct {
	Queue            string        // queue to declare.
	BindingKey       []types.Event // The routing key for binding the queue to an exchange.
	Durable          bool          // Whether the queue should survive server restarts (true for durable queues).
	AutoDelete       bool          // Whether the queue should be automatically deleted when no longer in use.
	Exclusive        bool          // Whether the queue is exclusive to the connection that declared it.
	NoWait           bool          // Whether the server should wait for a response before declaring the queue.
	ArgsQueueDeclare Table         // Additional arguments for queue declaration.
	ArgsQueueBind    Table         // Additional arguments for binding the queue.
	PrefetchCount    int           // The number of messages to prefetch (limit of unacknowledged messages).
	PrefetchSize     int           // The size limit for prefetching messages.
	PrefetchGlobal   bool          // Whether the prefetch settings apply globally across channels.
}

// ConsumeConfig defines the configuration for consuming messages from a queue in RabbitMQ.
type ConsumeConfig struct {
	AutoAck   bool  // Whether to automatically acknowledge messages after receiving them.
	Exclusive bool  // Whether the consumer should have exclusive access to the queue.
	NoLocal   bool  // Whether to prevent the consumer from receiving messages published on the same connection.
	NoWait    bool  // Whether the server should wait for a response before starting the consumer.
	Args      Table // Additional arguments for consuming messages.
}

// PublishConfig defines the configuration for publishing a message to an exchange in RabbitMQ.
type PublishConfig struct {
	ExchangeName string // Name of the exchange to publish the message to.
	// RoutingKey   string // The routing key used for routing the message.
	Mandatory bool // Whether the message must be delivered to at least one queue.
	Immediate bool // Whether the message must be immediately delivered to a consumer.
}
