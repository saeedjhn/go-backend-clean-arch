package event

import (
	"fmt"
	"sync"

	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/contract"
)

type RouterHandler func(event contract.Event) error

type Router struct {
	mu       sync.RWMutex
	handlers map[contract.Topic]RouterHandler
}

func NewRouter() *Router {
	return &Router{handlers: make(map[contract.Topic]RouterHandler)}
}

func (r *Router) Register(topic contract.Topic, handler RouterHandler) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.handlers[topic] = handler
}

func (r *Router) Handle(event contract.Event) error {
	r.mu.RLock()
	defer r.mu.RUnlock()

	handler, exists := r.handlers[event.Topic]
	if !exists {
		return fmt.Errorf("no handler for topic: %s", event.Topic)
	}

	return handler(event)
}
