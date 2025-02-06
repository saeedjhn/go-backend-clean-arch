package rmqpc

// ExchangeType defines the type of exchange in RabbitMQ.
// An exchange determines how messages are routed to queues based on various rules.
// The type of exchange determines the routing mechanism.
type ExchangeType string

const (
	// DirectExchange routes messages to queues based on an exact matching between the message's routing key and the queue's binding key.
	// This type is used when you need to deliver messages to specific queues based on exact routing keys.
	DirectExchange ExchangeType = "direct"

	// FanoutExchange routes messages to all bound queues regardless of the routing key.
	// This type is typically used for broadcasting messages to all queues.
	FanoutExchange ExchangeType = "fanout"

	// TopicExchange routes messages to queues based on pattern matching between the routing key and the queue's binding key.
	// This allows for more complex routing, including wildcards.
	TopicExchange ExchangeType = "topic"

	// HeadersExchange routes messages based on matching message headers instead of the routing key.
	// This is useful when routing decisions are based on attributes or metadata in the headers.
	HeadersExchange ExchangeType = "headers"
)

// String returns the string representation of the ExchangeType.
// This is useful for displaying the type as a string or for logging purposes.
func (e ExchangeType) String() string {
	return string(e)
}
