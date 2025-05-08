package bootstrap

import (
	"github.com/saeedjhn/go-backend-clean-arch/internal/event"
	userevthandler "github.com/saeedjhn/go-backend-clean-arch/internal/eventhandler/user"
	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/events"
	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/models"
)

func NewEventRegister() map[models.EventType]event.RouterHandler {
	return map[models.EventType]event.RouterHandler{
		events.UsersRegistered: userevthandler.NewRegistered().Execute,
	}
}
