package event

import (
	"context"
	"fmt"
	"sync"

	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/types"
)

type Router struct {
	mu       sync.RWMutex
	handlers map[types.Event]types.EventRouterHandler
}

func NewRouter() *Router {
	return &Router{handlers: make(map[types.Event]types.EventRouterHandler)}
}

func (r *Router) Register(topic types.Event, handler types.EventRouterHandler) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.handlers[topic] = handler
}

func (r *Router) Handle(ctx context.Context, event types.EventStream) error {
	r.mu.RLock()
	defer r.mu.RUnlock()

	handler, exists := r.handlers[event.Type]
	if !exists {
		return fmt.Errorf("no handler for topic: %s", event.Type)
	}

	return handler(ctx, event.Payload)
}
