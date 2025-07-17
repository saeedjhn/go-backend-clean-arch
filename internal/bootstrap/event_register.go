package bootstrap

import (
	userevthandler "github.com/saeedjhn/go-backend-clean-arch/internal/eventhandler/user"
	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/events"
	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/types"
)

func NewEventRegister() types.EventRouter {
	return types.EventRouter{
		events.UsersRegistered: userevthandler.NewRegistered().Execute,
	}
}
