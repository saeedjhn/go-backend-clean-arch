package event

import (
	"fmt"
	"sync"
)

type handlerFunc func(event Event) error

type Router struct {
	mu       sync.RWMutex
	handlers map[Topic]handlerFunc
}

func NewRouter() *Router {
	return &Router{handlers: make(map[Topic]handlerFunc)}
}

func (r *Router) Register(topic Topic, handler handlerFunc) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.handlers[topic] = handler
}

func (r *Router) Handle(event Event) error {
	r.mu.RLock()
	defer r.mu.RUnlock()

	handler, exists := r.handlers[event.Topic]
	if !exists {
		return fmt.Errorf("no handler for topic: %s", event.Topic)
	}

	return handler(event)
}
