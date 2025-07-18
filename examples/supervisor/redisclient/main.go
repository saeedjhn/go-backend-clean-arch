package main

import (
	"context"
	"time"

	"github.com/saeedjhn/go-backend-clean-arch/internal/adapter/jsonfilelogger"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/persistance/cache/redis"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/supervisor"
)

func process(err error) supervisor.ProcessFunc {
	return func(ctx context.Context, processName string, terminateChannel chan<- string) error {
		for {
			select {
			case <-ctx.Done():
				return nil
			default:
				// You might want to add a sleep here to reduce CPU usage
				return err
			}
		}
	}
}

func redisClient(ctx context.Context, processName string, terminateChannel chan<- string) error {
	client := redis.New(redis.Config{
		Host:               "localhost",
		Port:               "6378",
		Password:           "password123",
		DB:                 0,
		PoolSize:           0,
		DialTimeout:        0,
		ReadTimeout:        0,
		WriteTimeout:       0,
		PoolTimeout:        0,
		IdleCheckFrequency: 0,
	})
	if err := client.ConnectTo(); err != nil {
		return err
	}

	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			return client.Client().Ping(ctx).Err()
		}
	}
}

func main() {
	loggerStrategy := jsonfilelogger.NewDevelopmentStrategy(jsonfilelogger.Config{
		FilePath:         "./logs",
		Console:          true,
		File:             false,
		EnableCaller:     true,
		EnableStacktrace: true,
		Level:            "debug",
	})
	logger := jsonfilelogger.New(loggerStrategy).Configure()

	sv := supervisor.New(10*time.Second, logger)

	r := redis.New(redis.Config{
		Host:               "",
		Port:               "",
		Password:           "",
		DB:                 0,
		PoolSize:           0,
		DialTimeout:        0,
		ReadTimeout:        0,
		WriteTimeout:       0,
		PoolTimeout:        0,
		IdleCheckFrequency: 0,
	})

	// if err := r.ConnectTo(); err != nil {
	// 	log.Panicf("redis connect unsuccessfully: %v", err)
	// }

	// sv.Register("redis-client", redisClient, nil)
	sv.Register("redis-client", process(r.ConnectTo()), nil)
	sv.Start()

	sv.WaitOnShutdownSignal()
}
