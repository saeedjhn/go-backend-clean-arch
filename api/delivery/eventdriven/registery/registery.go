package registery

import (
	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/types"
)

type R struct {
	handlers types.EventRouter
}

func New() *R {
	return &R{
		handlers: make(types.EventRouter),
	}
}

func (r *R) Register(eventType types.Event, handler types.EventRouterHandler) {
	// if _, exists := r.handlers[eventType]; exists {
	// 	panic(fmt.Sprintf("handler for event %v already registered", eventType))
	// }
	r.handlers[eventType] = handler
}

func (r *R) Handlers() types.EventRouter {
	return r.handlers
}
