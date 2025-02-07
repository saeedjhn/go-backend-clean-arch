package contract

import (
	"github.com/saeedjhn/go-backend-clean-arch/internal/entity"
)

//go:generate mockery --name Publisher
type Publisher interface {
	Publish(event entity.Event) error
}
