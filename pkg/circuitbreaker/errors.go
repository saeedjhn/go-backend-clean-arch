package circuitbreaker

import "errors"

var (
	ErrBreakerNotFound = errors.New("circuit breaker not found")
	// ErrBreakerAlreadyOpen  = errors.New("circuit breaker already open").
	// ErrBreakerAlreadyClose = errors.New("circuit breaker already closed").
)
