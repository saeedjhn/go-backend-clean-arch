package retry

import "time"

type (
	ShouldRetryFunc func(error) bool
	OnRetryFunc     func(attempt uint, err error)
	OnSuccessFunc   func(attempts uint, totalDuration time.Duration)
)
