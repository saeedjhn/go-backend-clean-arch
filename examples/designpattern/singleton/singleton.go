package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	once     sync.Once
	instance *DriverPg
)

type single struct {
	O interface{}
}

var instantiated *single = nil

func New() *single {
	if instantiated == nil {
		instantiated = new(single)
	}

	return instantiated
}

type DriverPg struct {
	conn string
}

func Connect() *DriverPg {
	once.Do(func() {
		instance = &DriverPg{conn: "DriverConnectPostgres"}
	})
	return instance
}

func NewDBSingleton() (*DriverPg, error) {
	// var initErr error

	once.Do(func() {
		instance = &DriverPg{conn: "DriverConnectPostgres"}
		// initErr == true
	})

	// if initErr != nil {
	// 	return nil, initErr
	// }

	return instance, nil
}

func main() {
	// Simulate a delayed call to Connect.
	// go func() {
	// 	time.Sleep(time.Millisecond * 600)
	// 	fmt.Println(*Connect())
	// }()

	// Create 100 goroutines.
	for i := 0; i < 10; i++ {
		// go func(ix int) {
		time.Sleep(time.Millisecond * 60)
		fmt.Println(i, " = ", Connect().conn)
		// }(i)
	}

	fmt.Scanln()
}
