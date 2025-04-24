package retry

import (
	"context"
	"fmt"
	"time"

	"github.com/avast/retry-go/v4"
)

func Retry(ctx context.Context, op func() error, opts ...Option) error {
	config := defaultConfig()
	for _, opt := range opts {
		opt(&config)
	}

	options := []retry.Option{
		retry.Context(ctx),
		retry.Attempts(config.Attempts),
		retry.DelayType(config.getDelayFunc()),
		retry.LastErrorOnly(config.LastErrorOnly),
	}

	// if config.maxDelay > 0 {
	// 	options = append(options, retry.MaxDelay(config.maxDelay))
	// }

	if config.Delay > 0 {
		options = append(options, retry.Delay(config.Delay), retry.DelayType(retry.FixedDelay))
	}

	if config.MaxJitter > 0 {
		options = append(options, retry.MaxJitter(config.MaxJitter))
	}

	if config.ShouldRetry != nil {
		options = append(options, retry.RetryIf(func(err error) bool {
			return config.ShouldRetry(err)
		}))
	}

	if config.OnRetry != nil {
		options = append(options, retry.OnRetry(func(n uint, err error) {
			config.OnRetry(n+1, err)
		}))
	}

	var attempts uint
	var lastError error
	startTime := time.Now()

	err := retry.Do(
		func() error {
			attempts++
			lastError = op()
			return lastError
		},
		options...,
	)

	if err != nil {
		return fmt.Errorf("after %d Attempts: %w", attempts, lastError)
	}

	if config.OnSuccess != nil {
		config.OnSuccess(attempts, time.Since(startTime))
	}

	return nil
}
