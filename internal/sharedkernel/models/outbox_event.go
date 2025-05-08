package models

import (
	"time"

	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/types"
)

type OutboxEvent struct {
	ID            types.ID  `json:"id"`
	Type          EventType `json:"type"`
	Payload       []byte    `json:"payload"`
	IsPublished   bool      `json:"is_published"`
	ReTriedCount  uint      `json:"retried_count"`
	LastRetriedAt time.Time `json:"last_retried_at"`
	PublishedAt   time.Time `json:"published_at"`
}
