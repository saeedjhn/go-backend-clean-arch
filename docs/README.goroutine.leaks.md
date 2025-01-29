## What is a Goroutine Leak?

A goroutine leak occurs when a goroutine remains active indefinitely without completing its task or getting cleaned up.
This can happen due to various reasons, such as blocked channels, infinite loops, or forgotten goroutines. Over time,
leaked goroutines can accumulate, leading to increased memory usage and degraded application performance.

### Monitoring with runtime

The runtime the package provides a way to count active goroutines using runtime.NumGoroutine().

```go
package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	go func() {
		for {
			time.Sleep(10 * time.Second)
		}
	}()

	time.Sleep(1 * time.Second)
	fmt.Println("Number of Goroutines:", runtime.NumGoroutine())
}
```

### Profiling with pprof

The pprof package helps profile goroutines in your application. Add the following code to enable pprof.

```go
package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	// Your application code
}

```

### Common Scenarios and Solutions

#### Scenario 1: Blocked on a Channel

Problem: A goroutine is waiting indefinitely for a value from a channel that is never sent.

```go
package main

import (
	"fmt"
	"time"
)

func receive(ch chan int) {
	fmt.Println(<-ch) // Goroutine leaks if no value is sent to ch
}

func main() {
	ch := make(chan int)
	go receive(ch)
	// No value sent to ch, causing receive to block forever
}

// Solution: Ensure that the channel is either closed or always has a value sent.
func receive(ch chan int) {
	select {
	case val := <-ch:
		fmt.Println(val)
	case <-time.After(5 * time.Second):
		fmt.Println("Timeout")
	}
}

func main() {
	ch := make(chan int)
	go receive(ch)
	// Simulating some work
	time.Sleep(1 * time.Second)
	ch <- 42 // Prevents the leak by sending a value
}

```

#### Scenario 2: Forgotten Goroutine

Problem: A goroutine is started but forgotten, leading to potential leaks.

```go
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	go func() {
		for {
			time.Sleep(10 * time.Second)
			fmt.Println("Doing work")
		}
	}()
	// The goroutine runs indefinitely
}

// Solution: Use a context to manage the lifecycle of the goroutine.
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // Ensures resources are cleaned up

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Goroutine stopped")
				return
			case <-time.After(10 * time.Second):
				fmt.Println("Doing work")
			}
		}
	}(ctx)

	// Simulate main work
	time.Sleep(15 * time.Second)
	cancel() // Stop the goroutine
}
```

#### Scenario 3: Leaking in a Loop

Problem: Spawning a new goroutine inside a loop without proper termination.

```go

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("Goroutine", i)
		}()
	}
	time.Sleep(1 * time.Second) // Wait to see output
}

// Solution: Use a wait group to ensure all goroutines complete.
func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println("Goroutine", i)
		}(i)
	}
	wg.Wait() // Wait for all goroutines to finish
}
```

#### Scenario 4: Leaking through HTTP Handlers

Problem: HTTP handlers can cause leaks if they spawn goroutines that never terminate.

```go
package main

import (
	"fmt"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	go func() {
		fmt.Println("Handling request")
		// Simulate work
		time.Sleep(10 * time.Second)
	}()
	fmt.Fprintln(w, "Request received")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

// Solution: Use context with request handling.
func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	go func() {
		select {
		case <-ctx.Done():
			fmt.Println("Request canceled")
		case <-time.After(10 * time.Second):
			fmt.Println("Handled request")
		}
	}()
	fmt.Fprintln(w, "Request received")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
```

#### Scenario 5: Wait Group Misuse

Problem: A goroutine never calls Done() on the wait group, causing the Wait() to block indefinitely.

```go

package main

import (
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	if true {
		go func() {
			defer wg.Done()
			// Simulate work
			time.Sleep(10 * time.Second)
		}()
	}
	wg.Wait() // Blocks indefinitely if Done() is not called
}

// Solution:
// Condition Check Placement: The wg.Add(1) should be placed inside the if block. This ensures wg.Add(1) is only called when the condition is true.This avoids adding to the wait group count when no goroutine is launched, preventing the wg.Wait() from blocking indefinitely.
// Ensuring Done() is Called: By using defer wg.Done() inside the goroutine, it ensures Done() is called even if the goroutine panics or returns early.
// Function for Condition: I added an example function someCondition that you can define according to your actual condition logic.

func main() {
	var wg sync.WaitGroup

	if someCondition() { // Assuming `someCondition` is a function returning a boolean
		wg.Add(1)
		go func() {
			defer wg.Done()
			// Simulate work
			time.Sleep(10 * time.Second)
		}()
	}
	wg.Wait() // Correctly waits for the goroutine to finish
}

// Example function to use in the if statement
func someCondition() bool {
	// Implement your condition here
	return true
}

```
