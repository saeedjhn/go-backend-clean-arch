package event

import (
	"context"
	"sync"

	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/contract"
)

// type C struct {
// 	logger         contract.Logger
// 	Consumers      []contract.Consumer
// 	Router         *Router
// 	chanBufferSize uint64
// 	eventStream    chan contract.Event
// 	shutdownOnce   sync.Once
// 	shutdownChan   chan struct{}  // استفاده از کانال برای سیگنال shutdown
// 	wg             sync.WaitGroup // برای مدیریت گوروتین‌ها
// }
//
// func NewEventConsumer(
// 	chanBufferSize uint64,
// 	router *Router,
// 	consumers ...contract.Consumer,
// ) *C {
// 	return &C{
// 		chanBufferSize: chanBufferSize,
// 		Consumers:      consumers,
// 		Router:         router,
// 		shutdownChan:   make(chan struct{}),
// 		eventStream:    make(chan contract.Event, chanBufferSize),
// 	}
// }
//
// func (c *C) WithLogger(logger contract.Logger) *C {
// 	c.logger = logger
//
// 	return c
// }
//
// func (c *C) Start() {
// 	c.wg.Add(len(c.Consumers) + 1) // +1 برای پردازشگر رویدادها
//
// 	// شروع مصرف‌کنندگان
// 	for _, consumer := range c.Consumers {
// 		go func(consumer contract.Consumer) {
// 			defer c.wg.Done()
// 			for {
// 				select {
// 				case <-c.shutdownChan:
// 					c.logger.Info("Consumer shutting down")
// 					return
// 				default:
// 					if err := consumer.Consume(c.eventStream); err != nil {
// 						c.logger.Errorf("Consumer failed: %v", err)
// 					}
// 				}
// 			}
// 		}(consumer)
// 	}
//
// 	// پردازش رویدادها
// 	go func() {
// 		defer c.wg.Done()
// 		for {
// 			select {
// 			case <-c.shutdownChan:
// 				c.logger.Info("Event processor shutting down")
// 				return
// 			case e, ok := <-c.eventStream:
// 				if !ok {
// 					return
// 				}
// 				if err := c.Router.Handle(e); err != nil {
// 					c.logger.Errorf("Error handling event: %v", err)
// 				}
// 			}
// 		}
// 	}()
// }
//
// func (c *C) Shutdown(ctx context.Context) error {
// 	var shutdownErr error
//
// 	c.shutdownOnce.Do(func() {
// 		close(c.shutdownChan) // سیگنال shutdown به همه گوروتین‌ها
// 		close(c.eventStream)  // بستن کانال رویدادها
//
// 		done := make(chan struct{})
// 		go func() {
// 			c.wg.Wait() // منتظر پایان تمام گوروتین‌ها
// 			close(done)
// 		}()
//
// 		select {
// 		case <-done:
// 			c.logger.Info("Shutdown completed successfully")
// 		}
// 	})
//
// 	return shutdownErr
// }

type C struct {
	logger         contract.Logger
	Consumers      []contract.Consumer
	Router         *Router
	chanBufferSize uint64
	shutdownChan   chan struct{}
	eventStream    chan contract.Event
}

func NewEventConsumer(
	chanBufferSize uint64,
	router *Router,
	consumers ...contract.Consumer,
) *C {
	return &C{
		chanBufferSize: chanBufferSize,
		Consumers:      consumers,
		Router:         router,
		shutdownChan:   make(chan struct{}),
		eventStream:    make(chan contract.Event, chanBufferSize),
	}
}

func (c *C) WithLogger(logger contract.Logger) *C {
	c.logger = logger

	return c
}

func (c *C) Start() { //nolint:gocognit // nothing
	// c.eventStream = make(chan contract.Event, c.chanBufferSize)
	// var once sync.Once // To ensure eventStream is closed only once

	// Start consumers
	for _, consumer := range c.Consumers {
		// consumer := consumer // Prevent closure issue
		go func() {
			for {
				select {
				case <-c.shutdownChan:
					c.logger.Info("[Start] Consumer shutting down due to context cancellation")
					return
				default:
					if err := consumer.Consume(c.eventStream); err != nil {
						c.logger.Errorf("[Start] Consumer failed: %v", err)
						return
					}
				}
			}
		}()
	}

	go func() {
		for {
			select {
			case <-c.shutdownChan:
				c.logger.Info("[Start] Event processor shutting down due to context cancellation")
				return
			case e, ok := <-c.eventStream:
				if !ok {
					c.logger.Info("[Start] Event stream closed, stopping event processing")
					return
				}
				if err := c.Router.Handle(e); err != nil {
					c.logger.Errorf("[Start] Error handling event [%s]: %v", e.Topic, err)
				}
			}
		}
	}()

	// Ensure eventStream is closed once when context is done
	// go func() {
	// 	<-ctx.Done()
	// 	once.Do(func() {
	// 		close(eventStream)
	// 		c.logger.Info("[Start] Event stream closed")
	// 	})
	// }()
}

func (c *C) Shutdown(_ context.Context) error {
	var once sync.Once // To ensure eventStream is closed only once
	once.Do(func() {
		close(c.shutdownChan)
		close(c.eventStream)

		c.logger.Info("[Shutdown] Event stream closed")
	})

	return nil
}
