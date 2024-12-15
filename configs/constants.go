package configs

const (
	AuthMiddlewareContextKey = "claims"
	PrometheusSubSytemName   = "app" // Similar to (- job_name: app) in prometheus.config.yml
	LoggerExcludePath        = "/metrics"
)
