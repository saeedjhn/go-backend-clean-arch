package setuptest

import (
	"github.com/saeedjhn/go-backend-clean-arch/internal/adaptor/rmqpc"
)

func NewRabbitMQ(connName string, config rmqpc.ConnectionConfig) *rmqpc.Connection {
	conn := rmqpc.New(connName, rmqpc.Config{
		Connection: config,
		MQ:         rmqpc.MQConfig{},
	})

	return conn
}
