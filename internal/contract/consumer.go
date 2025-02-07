package contract

import (
	"github.com/saeedjhn/go-backend-clean-arch/internal/entity"
)

//go:generate mockery --name Consumer
type Consumer interface {
	Consume(chan<- entity.Event) error
}
