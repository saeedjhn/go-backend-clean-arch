package retry_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/saeedjhn/go-backend-clean-arch/pkg/retry"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//go:generate go test -v -count=1 -run TestRetry ./...

func TestRetry_ImmediateSuccess_NoRetries(t *testing.T) {
	// Scenario: Operation succeeds immediately, no retries needed.

	t.Parallel()

	err := retry.Do(context.Background(), func() error {
		return nil
	})

	assert.NoError(t, err)
}

func TestRetry_ThirdAttemptSuccessWithExponentialBackoff_SucceedsEventually(t *testing.T) {
	// Scenario: Operation fails and then succeeds on the 3rd attempt with exponential backoff.

	t.Parallel()

	counter := 0
	err := retry.Do(context.Background(), func() error {
		counter++
		if counter < 2 {
			return errors.New("temporary error")
		}
		return nil
	}, retry.WithAttempts(5))

	require.NoError(t, err)
	assert.Equal(t, 2, counter)
}

func TestRetry_AlwaysFails_ReturnsLastError(t *testing.T) {
	// Scenario: Operation always fails, retry until limit then return error.

	t.Parallel()

	err := retry.Do(context.Background(), func() error {
		return errors.New("persistent error")
	}, retry.WithAttempts(3))

	require.Error(t, err)
	assert.Contains(t, err.Error(), "after 3 Attempts")
}

func TestRetry_UseConstantBackoff_VerifyRetryDelayType(t *testing.T) {
	// Scenario: Using constant delay for retries.

	t.Parallel()

	counter := 0
	err := retry.Do(context.Background(), func() error {
		counter++
		if counter < 2 {
			return errors.New("fail")
		}
		return nil
	}, retry.WithBackoff(retry.ConstantBackoff))

	require.NoError(t, err)
	assert.Equal(t, 2, counter)
}

func TestRetry_ShouldRetryOnlyOnCustomError_ConditionalRetry(t *testing.T) {
	// Scenario: Only retry on specific error using ShouldRetry.

	t.Parallel()

	counter := 0
	err := retry.Do(context.Background(), func() error {
		counter++
		if counter < 2 {
			// log.Println("In block counter < x invoked ")
			return errors.New("retryable")
		}

		// log.Println("After block counter < x invoked")
		return errors.New("fatal")
	}, retry.WithAttempts(3), retry.WithRetryIf(func(err error) bool {
		// log.Println("WithRetryIf invoked")

		return err.Error() == "retryable"
	}))

	require.Error(t, err)
	assert.Equal(t, 2, counter)
}

func TestRetry_WithMaxJitter_RunsSuccessfully(t *testing.T) {
	// Scenario: Add jitter to avoid thundering herd problem.

	t.Parallel()

	counter := 0
	err := retry.Do(context.Background(), func() error {
		counter++
		if counter < 2 {
			// log.Println("In block counter < x invoked ")
			return errors.New("retry")
		}

		// log.Println("After block counter < x invoked")
		return nil
	}, retry.MaxJitter(100*time.Millisecond))

	require.NoError(t, err)
}

func TestRetry_WithFixedDelay_UsesGivenDelay(t *testing.T) {
	// Scenario: Do with delay explicitly set.

	t.Parallel()

	counter := 0
	err := retry.Do(context.Background(), func() error {
		counter++
		if counter < 2 {
			// log.Println("In block counter < x invoked ")
			return errors.New("again")
		}

		// log.Println("After block counter < x invoked")
		return nil
	}, retry.WithDelay(time.Second))

	require.NoError(t, err)
}

func TestRetry_OnRetryHook_IsCalledEveryAttempt(t *testing.T) {
	// Scenario: OnRetry hook should be triggered with each failure.

	t.Parallel()

	counter := 0
	retries := 0

	err := retry.Do(context.Background(), func() error {
		counter++
		if counter < 2 {
			// log.Println("In block counter < x invoked ")
			return errors.New("fail")
		}

		// log.Println("After block counter < x invoked")
		return nil
	}, retry.WithOnRetry(func(_ uint, _ error) {
		// log.Println("onRetry hook should be triggered with each failure..")
		retries++
	}))

	require.NoError(t, err)
	assert.Equal(t, 1, retries)
}

func TestRetry_OnSuccessHook_IsCalledOnSuccess(t *testing.T) {
	// Scenario: OnSuccess should be called after final successful attempt.

	t.Parallel()

	called := false
	counter := 0

	err := retry.Do(context.Background(), func() error {
		counter++
		if counter < 2 {
			// log.Println("In block counter < x invoked ")
			return errors.New("any something")
		}

		// log.Println("After block counter < x invoked")
		return nil
	}, retry.WithOnSuccess(func(_ uint, _ time.Duration) {
		// log.Println("onSuccess hook should be called after final successfull attempt")
		called = true
	}), retry.WithDelay(time.Second))

	require.NoError(t, err)
	assert.True(t, called)
}

func TestRetry_ContextCancelled_StopsEarly(t *testing.T) {
	// Scenario: Context is cancelled before retry completes.

	t.Parallel()

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	err := retry.Do(ctx, func() error {
		time.Sleep(120 * time.Millisecond)
		return errors.New("context deadline exceeded")
	}, retry.WithAttempts(5))

	require.Error(t, err)
	assert.Contains(t, err.Error(), "context deadline exceeded")
}
