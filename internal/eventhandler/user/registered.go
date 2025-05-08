package user

import (
	"log"
)

type Registered struct {
}

func NewRegistered() *Registered {
	return &Registered{}
}

func (c Registered) Execute(payload []byte) error {
	log.Printf("[Notification] Sending welcome email for user: %s\n", payload)

	return nil
}
