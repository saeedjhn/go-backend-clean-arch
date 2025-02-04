package event

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"sync"
)

type C struct {
	chanBufferSize uint64
	Consumers      []Consumer
	Router         *Router
}

func NewEventConsumer(
	chanBufferSize uint64,
	router *Router,
	consumers ...Consumer,
) *C {
	return &C{
		chanBufferSize: chanBufferSize,
		Consumers:      consumers,
		Router:         router,
	}
}

func (c *C) Start(ctx context.Context, wg *sync.WaitGroup) error {
	eventStream := make(chan Event, c.chanBufferSize)
	g, ctx := errgroup.WithContext(ctx)

	var once sync.Once // To ensure eventStream is closed only once

	// Start consumers
	for _, consumer := range c.Consumers {
		consumer := consumer // Prevent closure issue
		g.Go(func() error {
			for {
				select {
				case <-ctx.Done():
					log.Println("[Start] Consumer shutting down due to context cancellation")
					return ctx.Err()
				default:
					if err := consumer.Consume(eventStream); err != nil {
						log.Printf("[Start] Consumer failed: %v", err)
						return fmt.Errorf("[Start] consumer execution error: %w", err)
					}
				}
			}
		})
	}

	wg.Add(1)
	g.Go(func() error {
		defer wg.Done()

		for {
			select {
			case <-ctx.Done():
				log.Println("[Start] Event processor shutting down due to context cancellation")
				return ctx.Err()
			case e, ok := <-eventStream:
				if !ok {
					log.Println("[Start] Event stream closed, stopping event processing")
					return nil
				}
				if err := c.Router.Handle(e); err != nil {
					log.Printf("[Start] Error handling event [%s]: %v", e.Topic, err)
					return fmt.Errorf("[Start] failed to handle event [%s]: %w", e.Topic, err)
				}
			}
		}
	})

	// Ensure eventStream is closed once when context is done
	go func() {
		<-ctx.Done()
		once.Do(func() {
			close(eventStream)
		})
	}()

	err := g.Wait()

	// Ensure eventStream is closed properly
	once.Do(func() {
		close(eventStream)
	})

	if err != nil {
		if errors.Is(err, context.Canceled) {
			log.Println("[Start] Event consumer stopped gracefully due to context cancellation")
			return nil
		}

		return fmt.Errorf("[Start] event consumer encountered an error: %w", err)
	}

	log.Println("[Start] Successfully completed event processing")

	return nil
}

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
