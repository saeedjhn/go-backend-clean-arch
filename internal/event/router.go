package event

import (
	"fmt"
	"sync"

	"github.com/saeedjhn/go-backend-clean-arch/internal/entity"
)

type Router struct {
	mu       sync.RWMutex
	handlers map[entity.Topic]entity.RouterHandler
}

func NewRouter() *Router {
	return &Router{handlers: make(map[entity.Topic]entity.RouterHandler)}
}

func (r *Router) Register(topic entity.Topic, handler entity.RouterHandler) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.handlers[topic] = handler
}

func (r *Router) Handle(event entity.Event) error {
	r.mu.RLock()
	defer r.mu.RUnlock()

	handler, exists := r.handlers[event.Topic]
	if !exists {
		return fmt.Errorf("no handler for topic: %s", event.Topic)
	}

	return handler(event)
}
