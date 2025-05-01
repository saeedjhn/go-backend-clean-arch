package bootstrap

import (
	"github.com/saeedjhn/go-backend-clean-arch/internal/event"
	userevthandler "github.com/saeedjhn/go-backend-clean-arch/internal/eventhandler/user"
	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/contract"
	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/events"
)

func NewEventRegister() map[contract.Topic]event.RouterHandler {
	return map[contract.Topic]event.RouterHandler{
		events.UsersRegistered: userevthandler.NewRegistered().Execute,
	}
}
