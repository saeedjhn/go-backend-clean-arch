package main

import (
	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/models"
)

type InMemoryBus struct {
	// contractStream chan contract.Event
	contractStream chan models.Event
}

func NewInMemoryBus() *InMemoryBus {
	return &InMemoryBus{contractStream: make(chan models.Event, _eventBufferSize)}
}

func (b *InMemoryBus) Publish(contract models.Event) error {
	b.contractStream <- contract

	return nil
}

func (b *InMemoryBus) Consume(ch chan<- models.Event) error {
	// go func() {
	for evt := range b.contractStream {
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
// 	contractStream  chan Event
// 	messageStore sync.Map
// 	mu           sync.Mutex
// 	counter      int64
// }
//
// func NewInMemoryBus(bufferSize int) *InMemoryBus {
// 	return &InMemoryBus{
// 		contractStream: make(chan Event, bufferSize),
// 	}
// }
//
// func (b *InMemoryBus) Publish(contract Event) error {
// 	b.mu.Lock()
// 	defer b.mu.Unlock()
//
// 	b.messageStore.Create(contract.ID, contract)
// 	b.contractStream <- contract
//
// 	return nil
// }
//
// func (b *InMemoryBus) Consume(handler func(Event) bool) {
// 	go func() {
// 		for contract := range b.contractStream {
// 			success := handler(contract)
// 			if success {
// 				b.messageStore.Delete(contract.ID)
// 			} else {
// 				fmt.Println("Processing failed, keeping contract for retry:", contract.ID)
// 			}
// 		}
// 	}()
// }
//
// func (b *InMemoryBus) RetryFailedMessages() {
// 	b.messageStore.Range(func(key, value any) bool {
// 		contract := value.(Event)
// 		fmt.Println("Retrying contract:", contract.ID)
// 		b.contractStream <- contract
// 		return true
// 	})
// }
