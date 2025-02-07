package job

import (
	"log"

	"github.com/saeedjhn/go-backend-clean-arch/internal/entity"
)

func UserRegisteredHandler(event entity.Event) error {
	log.Printf("[Notification] Sending welcome email for user: %s\n", string(event.Payload))

	return nil
}
