package event

import (
	"context"
	"sync"

	contract2 "github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/contract"
)

type C struct {
	logger         contract2.Logger
	Consumers      []contract2.Consumer
	Router         *Router
	chanBufferSize uint64
}

func NewEventConsumer(
	chanBufferSize uint64,
	router *Router,
	consumers ...contract2.Consumer,
) *C {
	return &C{
		chanBufferSize: chanBufferSize,
		Consumers:      consumers,
		Router:         router,
	}
}

func (c *C) WithLogger(logger contract2.Logger) *C {
	c.logger = logger

	return c
}

func (c *C) Start(ctx context.Context) { //nolint:gocognit // nothing
	eventStream := make(chan contract2.Event, c.chanBufferSize)
	var once sync.Once // To ensure eventStream is closed only once

	// Start consumers
	for _, consumer := range c.Consumers {
		// consumer := consumer // Prevent closure issue
		go func() {
			for {
				select {
				case <-ctx.Done():
					c.logger.Info("[Start] Consumer shutting down due to context cancellation")
					return
				default:
					if err := consumer.Consume(eventStream); err != nil {
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
			case <-ctx.Done():
				c.logger.Info("[Start] Event processor shutting down due to context cancellation")
				return
			case e, ok := <-eventStream:
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
	go func() {
		<-ctx.Done()
		once.Do(func() {
			close(eventStream)
			c.logger.Info("[Start] Event stream closed")
		})
	}()
}

// func (c *C) Start(ctx context.Context, wg *sync.WaitGroup) {
// 	eventStream := make(chan models.Event, c.chanBufferSize)
// 	var once sync.Once // To ensure eventStream is closed only once
//
// 	// Start consumers
// 	for _, consumer := range c.Consumers {
// 		consumer := consumer // Prevent closure issue
// 		wg.Add(1)
// 		go func() {
// 			defer wg.Done()
// 			for {
// 				select {
// 				case <-ctx.Done():
// 					c.logger.Info("[Start] Consumer shutting down due to context cancellation")
// 					return
// 				default:
// 					if err := consumer.Consume(eventStream); err != nil {
// 						c.logger.Errorf("[Start] Consumer failed: %v", err)
// 						return
// 					}
// 				}
// 			}
// 		}()
// 	}
//
// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		for {
// 			select {
// 			case <-ctx.Done():
// 				c.logger.Info("[Start] Event processor shutting down due to context cancellation")
// 				return
// 			case e, ok := <-eventStream:
// 				if !ok {
// 					c.logger.Info("[Start] Event stream closed, stopping event processing")
// 					return
// 				}
// 				if err := c.Router.Handle(e); err != nil {
// 					c.logger.Errorf("[Start] Error handling event [%s]: %v", e.Topic, err)
// 				}
// 			}
// 		}
// 	}()
//
// 	// Ensure eventStream is closed once when context is done
// 	go func() {
// 		<-ctx.Done()
// 		once.Do(func() {
// 			close(eventStream)
// 			c.logger.Info("[Start] Event stream closed")
// 		})
// 	}()
// }

// func (c *C) Start(ctx context.Context, wg *sync.WaitGroup) error {
// 	eventStream := make(chan Event, c.chanBufferSize)
// 	g, ctx := errgroup.WithContext(ctx)
//
// 	var once sync.Once // To ensure eventStream is closed only once
//
// 	// Start consumers
// 	for _, consumer := range c.Consumers {
// 		consumer := consumer // Prevent closure issue
// 		g.Go(func() error {
// 			for {
// 				select {
// 				case <-ctx.Done():
// 					log.Println("[Start] Consumer shutting down due to context cancellation")
// 					return ctx.Err()
// 				default:
// 					if err := consumer.Consume(eventStream); err != nil {
// 						log.Printf("[Start] Consumer failed: %v", err)
// 						return fmt.Errorf("[Start] consumer execution error: %w", err)
// 					}
// 				}
// 			}
// 		})
// 	}
//
// 	wg.Add(1)
// 	g.Go(func() error {
// 		defer wg.Done()
//
// 		for {
// 			select {
// 			case <-ctx.Done():
// 				log.Println("[Start] Event processor shutting down due to context cancellation")
// 				return ctx.Err()
// 			case e, ok := <-eventStream:
// 				if !ok {
// 					log.Println("[Start] Event stream closed, stopping event processing")
// 					return nil
// 				}
// 				if err := c.Router.Handle(e); err != nil {
// 					log.Printf("[Start] Error handling event [%s]: %v", e.Topic, err)
// 					return fmt.Errorf("[Start] failed to handle event [%s]: %w", e.Topic, err)
// 				}
// 			}
// 		}
// 	})
//
// 	// Ensure eventStream is closed once when context is done
// 	go func() {
// 		<-ctx.Done()
// 		once.Do(func() {
// 			close(eventStream)
// 		})
// 	}()
//
// 	err := g.Wait()
//
// 	// Ensure eventStream is closed properly
// 	once.Do(func() {
// 		close(eventStream)
// 	})
//
// 	if err != nil {
// 		if errors.Is(err, context.Canceled) {
// 			log.Println("[Start] Event consumer stopped gracefully due to context cancellation")
// 			return nil
// 		}
//
// 		return fmt.Errorf("[Start] event consumer encountered an error: %w", err)
// 	}
//
// 	log.Println("[Start] Successfully completed event processing")
//
// 	return nil
// }

//
// func (c *C) Start(ctx context.Context, wg *sync.WaitGroup) error {
// 	eventStream := make(chan Event, c.chanBufferSize)
// 	defer close(eventStream)
//
// 	g, ctx := errgroup.WithContext(ctx)
//
// 	for _, consumer := range c.Consumers {
// 		// consumer := consumer
// 		g.Go(func() error {
// 			err := consumer.Consume(eventStream)
// 			if err != nil {
// 				return fmt.Errorf("consumer failed: %w", err)
// 			}
// 			return nil
// 		})
// 	}
//
// 	wg.Add(1)
// 	g.Go(func() error {
// 		defer wg.Done()
// 		for {
// 			select {
// 			case <-ctx.Done():
// 				log.Println("C shutting down...")
// 				return ctx.Err()
// 			case e, ok := <-eventStream:
// 				if !ok {
// 					return nil
// 				}
// 				if err := c.Router.Handle(e); err != nil {
// 					log.Printf("error handling event [%s]: %v", e.Topic, err)
// 					return fmt.Errorf("failed to handle event [%s]: %w", e.Topic, err)
// 				}
// 			}
// 		}
// 	})
//
// 	if err := g.Wait(); err != nil {
// 		return fmt.Errorf("event consumer encountered an error: %w", err)
// 	}
//
// 	return nil
// }
