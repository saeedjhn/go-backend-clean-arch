package event

import (
	"fmt"
	"sync"

	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/models"
)

type RouterHandler func(payload []byte) error

type Router struct {
	mu       sync.RWMutex
	handlers map[models.EventType]RouterHandler
}

func NewRouter() *Router {
	return &Router{handlers: make(map[models.EventType]RouterHandler)}
}

func (r *Router) Register(topic models.EventType, handler RouterHandler) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.handlers[topic] = handler
}

func (r *Router) Handle(event models.Event) error {
	r.mu.RLock()
	defer r.mu.RUnlock()

	handler, exists := r.handlers[event.Type]
	if !exists {
		return fmt.Errorf("no handler for topic: %s", event.Type)
	}

	return handler(event.Payload)
}
