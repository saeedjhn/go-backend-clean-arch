package otelcollector

import "time"

type Options struct {
	Config  Config
	AppInfo AppInfo
}

type AppInfo struct {
	Host    string
	Port    string
	Name    string
	Version string
	Env     string
	// InstanceID     string  `mapstructure:"METRIC_INSTANCE_ID"`
}

type Config struct {
	Endpoint string        `mapstructure:"endpoint"`
	Timeout  time.Duration `mapstructure:"timeout"`
}
