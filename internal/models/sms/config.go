package sms

import (
	"time"

	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/types"
)

type Config struct {
	ID           types.ID
	Title        string
	Slug         string
	Priority     float64
	Template     Template
	Status       Status
	ProviderID   types.ID
	SenderLineID types.ID
	Type         types.ID
	CreatedAt    time.Time
	UpdatedAt    time.Time
	// IsDefault    bool
}
