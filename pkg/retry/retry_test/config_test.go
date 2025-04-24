package retry_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/saeedjhn/go-backend-clean-arch/pkg/retry"
)

//go:generate go test -v -count=1 -run TestWith ./...

func TestWithAttempts_SetValidAttempts_ConfigHasExpectedAttempts(t *testing.T) {
	cfg := retry.Config{}
	opt := retry.WithAttempts(5)
	opt(&cfg)

	assert.Equal(t, uint(5), cfg.Attempts)
}

func TestWithBaseDelay_SetValidDelay_ConfigHasExpectedBaseDelay(t *testing.T) {
	cfg := retry.Config{}
	delay := 200 * time.Millisecond
	opt := retry.WithBaseDelay(delay)
	opt(&cfg)

	assert.Equal(t, delay, cfg.BaseDelay)
}

func TestWithDelay_SetCustomDelay_ConfigHasExpectedDelay(t *testing.T) {
	cfg := retry.Config{}
	delay := 1 * time.Second
	opt := retry.WithDelay(delay)
	opt(&cfg)

	assert.Equal(t, delay, cfg.Delay)
}

func TestWithBackoff_SetLinearBackoff_ConfigHasLinearBackoff(t *testing.T) {
	cfg := retry.Config{}
	opt := retry.WithBackoff(retry.LinearBackoff)
	opt(&cfg)

	assert.Equal(t, retry.LinearBackoff, cfg.BackoffType)
}

func TestMaxJitter_SetMaxJitter_ConfigHasExpectedJitter(t *testing.T) {
	cfg := retry.Config{}
	jitter := 50 * time.Millisecond
	opt := retry.MaxJitter(jitter)
	opt(&cfg)

	assert.Equal(t, jitter, cfg.MaxJitter)
}

func TestWithRetryIf_SetShouldRetryFunc_ConfigUsesProvidedFunc(t *testing.T) {
	cfg := retry.Config{}
	called := false
	fn := func(_ error) bool {
		called = true
		return true
	}
	opt := retry.WithRetryIf(fn)
	opt(&cfg)

	cfg.ShouldRetry(nil)
	assert.True(t, called)
}

func TestWithOnRetry_SetOnRetryFunc_ConfigUsesProvidedFunc(t *testing.T) {
	cfg := retry.Config{}
	called := false
	fn := func(_ uint, _ error) {
		called = true
	}
	opt := retry.WithOnRetry(fn)
	opt(&cfg)

	cfg.OnRetry(1, nil)
	assert.True(t, called)
}

func TestWithOnSuccess_SetOnSuccessFunc_ConfigUsesProvidedFunc(t *testing.T) {
	cfg := retry.Config{}
	called := false
	fn := func(_ uint, _ time.Duration) {
		called = true
	}
	opt := retry.WithOnSuccess(fn)
	opt(&cfg)

	cfg.OnSuccess(3, 1*time.Second)
	assert.True(t, called)
}

func TestWithLastErrorOnly_SetLastErrorOnly_ConfigHasLastErrorOnlyTrue(t *testing.T) {
	cfg := retry.Config{}
	opt := retry.WithLastErrorOnly()
	opt(&cfg)

	assert.True(t, cfg.LastErrorOnly)
}
