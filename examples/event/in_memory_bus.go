package main

import (
	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/types"
)

type InMemoryBus struct {
	// contractStream chan contract.EventStream
	contractStream chan types.EventStream
}

func NewInMemoryBus() *InMemoryBus {
	return &InMemoryBus{contractStream: make(chan types.EventStream, _eventBufferSize)}
}

func (b *InMemoryBus) Publish(contract types.EventStream) error {
	b.contractStream <- contract

	return nil
}

func (b *InMemoryBus) Consume(ch chan<- types.EventStream) error {
	// go func() {
	for evt := range b.contractStream {
		ch <- evt
	}
	// }()
	return nil
}

// type EventStream struct {
// 	ID   string
// 	Data string
// }
//
// type InMemoryBus struct {
// 	contractStream  chan EventStream
// 	messageStore sync.Map
// 	mu           sync.Mutex
// 	counter      int64
// }
//
// func NewInMemoryBus(bufferSize int) *InMemoryBus {
// 	return &InMemoryBus{
// 		contractStream: make(chan EventStream, bufferSize),
// 	}
// }
//
// func (b *InMemoryBus) Publish(contract EventStream) error {
// 	b.mu.Lock()
// 	defer b.mu.Unlock()
//
// 	b.messageStore.Create(contract.ID, contract)
// 	b.contractStream <- contract
//
// 	return nil
// }
//
// func (b *InMemoryBus) Consume(handler func(EventStream) bool) {
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
// 		contract := value.(EventStream)
// 		fmt.Println("Retrying contract:", contract.ID)
// 		b.contractStream <- contract
// 		return true
// 	})
// }
