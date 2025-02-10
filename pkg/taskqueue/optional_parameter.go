package taskqueue

import "time"

type TaskOption func(*taskOptions)

type taskOptions struct {
	queueName Queue
	retry     int
	delay     time.Duration
	at        time.Time
}

func WithQueueName(queueName Queue) TaskOption {
	return func(opts *taskOptions) {
		opts.queueName = queueName
	}
}

func WithRetry(retry int) TaskOption {
	return func(opts *taskOptions) {
		opts.retry = retry
	}
}

func WithDelay(delay time.Duration) TaskOption {
	return func(opts *taskOptions) {
		opts.delay = delay
	}
}

func WithAt(at time.Time) TaskOption {
	return func(opts *taskOptions) {
		opts.at = at
	}
}

func defaultTaskOptions() *taskOptions {
	return &taskOptions{
		queueName: QueueDefault,
		retry:     0,
		delay:     0,
		at:        time.Now(),
	}
}
