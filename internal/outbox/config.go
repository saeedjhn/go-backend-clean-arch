package outbox

import "time"

type Config struct {
	Interval       time.Duration `mapstructure:"interval"`
	BatchSize      int           `mapstructure:"batch_size"`
	RetryThreshold int           `mapstructure:"retry_threshold"`
}
