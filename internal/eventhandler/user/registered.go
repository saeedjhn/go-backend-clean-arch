package user

import (
	"log"

	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/contract"
)

type Registered struct {
}

func NewRegistered() *Registered {
	return &Registered{}
}

func (c Registered) Execute(evt contract.Event) error {
	log.Printf("[Notification] Sending welcome email for user: %s\n", string(evt.Payload))

	return nil
}
