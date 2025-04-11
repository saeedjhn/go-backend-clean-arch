package contract

import (
	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/models"
)

type StateMachine interface {
	ApplyEvent(event models.EventFSM) error
	GetCurrentState() models.StateFSM
}
