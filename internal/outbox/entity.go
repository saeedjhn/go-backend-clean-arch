package outbox

import (
	"time"

	"github.com/saeedjhn/go-backend-clean-arch/internal/contract"

	"github.com/saeedjhn/go-backend-clean-arch/internal/types"
)

type Event struct {
	ID            types.ID       `json:"id"`
	Topic         contract.Topic `json:"topic"`
	Payload       []byte         `json:"payload"`
	IsPublished   bool           `json:"is_published"`
	ReTriedCount  uint           `json:"retried_count"`
	LastRetriedAt time.Time      `json:"last_retried_at"`
	PublishedAt   time.Time      `json:"published_at"`
}
