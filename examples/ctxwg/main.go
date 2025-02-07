package main

import (
	"context"
	"log"
	"sync"
	"time"
)

func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			log.Println("Worker stopped")
			return
		default:
			log.Print("Worker is working...")
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		worker(ctx)
	}()

	time.Sleep(3 * time.Second)

	log.Println("Stopping all workers...")
	cancel()

	wg.Wait()

	log.Println("All workers stopped, exiting program.")
}

// func worker(ctx context.Context, id int) {
// 	for {
// 		select {
// 		case <-ctx.Done():
// 			log.Printf("Worker %d stopped\n", id)
// 			return
// 		default:
// 			log.Printf("Worker %d is working...\n", id)
// 			time.Sleep(1 * time.Second)
// 		}
// 	}
// }
//
// func main() {
// 	ctx, cancel := context.WithCancel(context.Background())
//
// 	for i := 1; i <= 3; i++ {
// 		go worker(ctx, i)
// 	}
//
// 	time.Sleep(3 * time.Second)
//
// 	log.Println("Stopping all workers...")
// 	cancel()
// }
