package configs

import "time"

const (
	OTPChars                    = "0123456789"
	OTPExpireTime time.Duration = 2 * 60 * 1000 * 1000000 // 2 minutes

	AuthMiddlewareContextKey = "claims"
	PrometheusSubSytemName   = "app" // Similar to (- job_name: app) in prometheus.config.yml
	LoggerExcludePath        = "/metrics"
	BcryptCost               = 10

	EventBufferSize = 2048
)
