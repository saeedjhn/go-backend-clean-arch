package contract

import "github.com/saeedjhn/go-backend-clean-arch/internal/entity"

type StateMachine interface {
	ApplyEvent(event entity.EventFSM) error
	GetCurrentState() entity.StateFSM
}
