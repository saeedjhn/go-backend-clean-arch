package retry

import (
	"time"

	"github.com/avast/retry-go/v4"
)

const (
	_attempts  = 10
	_baseDelay = 100 * time.Millisecond
	// _baseDelay = 1 * time.Second // for TEST.
)

type Config struct {
	Attempts      uint
	BaseDelay     time.Duration
	Delay         time.Duration
	BackoffType   BackoffType
	MaxJitter     time.Duration
	ShouldRetry   ShouldRetryFunc
	OnRetry       OnRetryFunc
	OnSuccess     OnSuccessFunc
	LastErrorOnly bool
}

type Option func(*Config)

func defaultConfig() Config {
	return Config{
		Attempts:    _attempts,
		BaseDelay:   _baseDelay,
		BackoffType: ExponentialBackoff,
	}
}

func WithAttempts(attempts uint) Option {
	return func(rc *Config) {
		rc.Attempts = attempts
	}
}

func WithBaseDelay(delay time.Duration) Option {
	return func(rc *Config) {
		rc.BaseDelay = delay
	}
}

func WithDelay(delay time.Duration) Option {
	return func(cfg *Config) {
		cfg.Delay = delay
	}
}

// func WithMaxDelay(max time.Duration) Option {
// 	return func(rc *Config) {
// 		rc.maxDelay = max
// 	}
// }

func WithBackoff(backoffType BackoffType) Option {
	return func(rc *Config) {
		rc.BackoffType = backoffType
	}
}

func MaxJitter(duration time.Duration) Option {
	return func(rc *Config) {
		rc.MaxJitter = duration
	}
}

func WithRetryIf(fn ShouldRetryFunc) Option {
	return func(rc *Config) {
		rc.ShouldRetry = fn
	}
}

func WithOnRetry(fn OnRetryFunc) Option {
	return func(rc *Config) {
		rc.OnRetry = fn
	}
}

func WithOnSuccess(fn OnSuccessFunc) Option {
	return func(rc *Config) {
		rc.OnSuccess = fn
	}
}

func WithLastErrorOnly() Option {
	return func(cfg *Config) {
		cfg.LastErrorOnly = true
	}
}

func (rc *Config) getDelayFunc() func(n uint, _ error, config *retry.Config) time.Duration {
	switch rc.BackoffType {
	case ExponentialBackoff:
		return exponentialDelay(rc.BaseDelay)
	case LinearBackoff:
		return linearDelay(rc.BaseDelay)
	case ConstantBackoff:
		return constantDelay(rc.BaseDelay)
	default:
		return exponentialDelay(rc.BaseDelay)
	}
}

func exponentialDelay(base time.Duration) func(n uint, _ error, _ *retry.Config) time.Duration {
	return func(n uint, _ error, _ *retry.Config) time.Duration {
		return base * (1 << n) // base * 2^n
	}
}

func linearDelay(baseDelay time.Duration) func(n uint, _ error, _ *retry.Config) time.Duration {
	return func(attempt uint, _ error, _ *retry.Config) time.Duration {
		return baseDelay * time.Duration(int64(attempt)) // #nosec G115
	}
}

func constantDelay(delay time.Duration) func(n uint, _ error, _ *retry.Config) time.Duration {
	return func(_ uint, _ error, _ *retry.Config) time.Duration {
		return delay
	}
}
