package contract

import "github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/models"

type DomainEvent interface {
	GetID() uint32
	GetType() models.EventType
	GetEscalationReason() string
	GetEscalationTime() int64
	Marshal() ([]byte, error)
	Unmarshal(b []byte) error
}
