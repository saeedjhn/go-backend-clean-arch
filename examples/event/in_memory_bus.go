package main

import "github.com/saeedjhn/go-backend-clean-arch/internal/event"

const _eventBufferSize = 1024

type InMemoryBus struct {
	eventStream chan event.Event
}

func NewInMemoryBus() *InMemoryBus {
	return &InMemoryBus{eventStream: make(chan event.Event, _eventBufferSize)}
}

func (b *InMemoryBus) Publish(event event.Event) error {
	b.eventStream <- event

	return nil
}

func (b *InMemoryBus) Consume(ch chan<- event.Event) error {
	// go func() {
	for evt := range b.eventStream {
		ch <- evt
	}
	// }()
	return nil
}

// type Event struct {
// 	ID   string
// 	Data string
// }
//
// type InMemoryBus struct {
// 	eventStream  chan Event
// 	messageStore sync.Map
// 	mu           sync.Mutex
// 	counter      int64
// }
//
// func NewInMemoryBus(bufferSize int) *InMemoryBus {
// 	return &InMemoryBus{
// 		eventStream: make(chan Event, bufferSize),
// 	}
// }
//
// func (b *InMemoryBus) Publish(event Event) error {
// 	b.mu.Lock()
// 	defer b.mu.Unlock()
//
// 	b.messageStore.Store(event.ID, event)
// 	b.eventStream <- event
//
// 	return nil
// }
//
// func (b *InMemoryBus) Consume(handler func(Event) bool) {
// 	go func() {
// 		for event := range b.eventStream {
// 			success := handler(event)
// 			if success {
// 				b.messageStore.Delete(event.ID)
// 			} else {
// 				fmt.Println("Processing failed, keeping event for retry:", event.ID)
// 			}
// 		}
// 	}()
// }
//
// func (b *InMemoryBus) RetryFailedMessages() {
// 	b.messageStore.Range(func(key, value any) bool {
// 		event := value.(Event)
// 		fmt.Println("Retrying event:", event.ID)
// 		b.eventStream <- event
// 		return true
// 	})
// }
