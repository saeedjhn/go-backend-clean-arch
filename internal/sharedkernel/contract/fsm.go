package contract

import "github.com/saeedjhn/go-domain-driven-design/internal/entity"

type StateMachine interface {
	ApplyEvent(event entity.EventFSM) error
	GetCurrentState() entity.StateFSM
}
