package event

import (
	"context"
	"sync"

	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/types"

	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/contract"
)

type C struct {
	ctx            context.Context
	logger         contract.Logger
	Consumers      []contract.Consumer
	Router         *Router
	chanBufferSize uint64
	shutdownChan   chan struct{}
	eventStream    chan types.EventStream
}

func NewEventConsumer(
	chanBufferSize uint64,
	router *Router,
	consumers ...contract.Consumer,
) *C {
	return &C{
		ctx:            context.Background(),
		chanBufferSize: chanBufferSize,
		Consumers:      consumers,
		Router:         router,
		shutdownChan:   make(chan struct{}),
		eventStream:    make(chan types.EventStream, chanBufferSize),
	}
}

func (c *C) WithLogger(logger contract.Logger) *C {
	c.logger = logger

	return c
}

func (c *C) WithContext(ctx context.Context) *C {
	c.ctx = ctx

	return c
}

func (c *C) Start() { //nolint:gocognit // nothing
	// c.eventStream = make(chan contract.EventStream, c.chanBufferSize)
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
				c.logger.Info("[Start] EventStream processor shutting down due to context cancellation")
				return
			case e, ok := <-c.eventStream:
				if !ok {
					c.logger.Info("[Start] EventStream stream closed, stopping event processing")
					return
				}
				if err := c.Router.Handle(c.ctx, e); err != nil {
					c.logger.Errorf("[Start] Error handling event [%s]: %v", e.Type, err)
				}
			}
		}
	}()

	// Ensure eventStream is closed once when context is done
	// go func() {
	// 	<-ctx.Done()
	// 	once.Do(func() {
	// 		close(eventStream)
	// 		c.logger.Info("[Start] EventStream stream closed")
	// 	})
	// }()
}

func (c *C) Shutdown(_ context.Context) error {
	var once sync.Once // To ensure eventStream is closed only once
	once.Do(func() {
		close(c.shutdownChan)
		close(c.eventStream)

		c.logger.Info("[Shutdown] EventStream stream closed")
	})

	return nil
}
