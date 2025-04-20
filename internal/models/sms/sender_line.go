package sms

import (
	"time"

	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/types"
)

type SenderLine struct {
	ID          types.ID
	Number      string
	Capacity    int
	IsActive    bool
	Description string
	ProviderID  types.ID
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
