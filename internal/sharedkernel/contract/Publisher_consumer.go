package contract

import (
	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/types"
)

//go:generate mockery --name Consumer
type Consumer interface {
	Consume(chan<- types.EventStream) error
}

//go:generate mockery --name Publisher
type Publisher interface {
	Publish(event types.EventStream) error
}

//go:generate mockery --name PublisherConsumer
type PublisherConsumer interface {
	Publisher
	Consumer
}
