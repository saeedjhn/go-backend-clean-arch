package main

import (
	"go-backend-clean-arch/configs"
	"go-backend-clean-arch/internal/scheduler"
	"log"
	"os"
	"os/signal"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

const (
	ThreadUse = 1
)

func main() {
	// Read config path from command line
	config := configs.Load(configs.Development)
	log.Printf("config: %#v\n", config)

	done := make(chan bool)

	wg.Add(ThreadUse)

	for i := 0; i < ThreadUse; i++ {
		go func() {
			defer wg.Done()

			sch := scheduler.New()
			sch.Start(done)
		}()
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt) // more SIGX (SIGINT, SIGTERM, etc)
	<-quit

	log.Println("received interrupt signal, shutting down gracefully..")

	done <- true
	time.Sleep(config.Application.GracefulShutdownTimeout)
}
