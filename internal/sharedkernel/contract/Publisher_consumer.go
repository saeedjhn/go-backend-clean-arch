package contract

import "github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/models"

//go:generate mockery --name Consumer
type Consumer interface {
	Consume(chan<- models.Event) error
}

//go:generate mockery --name Publisher
type Publisher interface {
	Publish(event models.Event) error
}

//go:generate mockery --name PublisherConsumer
type PublisherConsumer interface {
	Publisher
	Consumer
}
