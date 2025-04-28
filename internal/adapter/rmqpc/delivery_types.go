package rmqpc

// DeliveryType represents the type of message delivery mode in RabbitMQ.
// It defines whether the message is transient (will not be persisted) or persistent (will be persisted).
type DeliveryType uint8

const (
	// Transient indicates that the message is not persisted and may be lost in case of server crashes.
	// This delivery mode is typically used for non-critical or fast-processing messages.
	Transient DeliveryType = 1

	// Persistent indicates that the message will be persisted to disk and can survive server restarts.
	// This is suitable for critical messages that must be reliably delivered.
	Persistent DeliveryType = 2
)

// Uint8 returns the uint8 representation of the DeliveryType.
// This can be useful for interfacing with APIs or libraries that require the delivery mode as a uint8 value.
func (d DeliveryType) Uint8() uint8 {
	return uint8(d)
}
