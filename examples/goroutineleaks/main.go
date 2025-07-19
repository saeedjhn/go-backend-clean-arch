package main

import (
	"context"
	"log"
	"net/http"
	_ "net/http/pprof" // #nosec G108
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

// Identifying Goroutine Leaks
// Before diving into scenarios and solutions, it's crucial to know how to identify goroutine leaks.
// You can use Go's runtime and pprof packages to monitor and profile goroutines.
//
// Monitoring with runtime
// The runtime the package provides a way to count active goroutines using runtime.NumGoroutine().

// Profiling with pprof
// The pprof package helps profile goroutines in your application. Add the following code to enable pprof.
const (
	_port              = ":8081"
	_sleepDuration     = 10 * time.Second
	_readTimeout       = 10 * time.Second
	_readHeaderTimeout = 5 * time.Second
	_writeTimeout      = 10 * time.Second
	_idleTimeout       = 2 * time.Minute
)

func main() {
	// Create a context that can be cancelled
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Fixed: Add proper shutdown mechanism for the infinite loop goroutine
	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Println("Background goroutine shutting down gracefully")
				return
			default:
				time.Sleep(_sleepDuration)
			}
		}
	}()

	go func() {
		mux := http.NewServeMux()
		server := http.Server{
			Addr:              _port,
			Handler:           mux,
			ReadTimeout:       _readTimeout,
			ReadHeaderTimeout: _readHeaderTimeout,
			WriteTimeout:      _writeTimeout,
			IdleTimeout:       _idleTimeout,
		}
		log.Printf("Server.PPROF.Starting - Starting PPROF server on port: %s", _port)

		// This line is not strictly required because the `net/http/pprof` package
		// automatically registers to the default HTTP server, but it's here for clarity.
		mux.HandleFunc("/debug/pprof/", http.DefaultServeMux.ServeHTTP)

		if err := server.ListenAndServe(); err != nil {
			log.Fatalf("Server.PPROF.ListenAndServe - Failed to start PPROF server: %v", err)
		}
	}()

	time.Sleep(1 * time.Second)
	log.Println("Number of Goroutines:", runtime.NumGoroutine())

	// Fixed: Add proper signal handling for graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down gracefully...")
	cancel() // Cancel the context to stop the background goroutine
	
	// Give some time for goroutines to finish
	time.Sleep(1 * time.Second)
	log.Println("Final number of Goroutines:", runtime.NumGoroutine())
}
