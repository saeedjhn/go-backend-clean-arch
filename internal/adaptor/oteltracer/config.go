package oteltracer

type Config struct {
	Endpoint   string  `mapstructure:"endpoint"`
	AppHost    string  `mapstructure:"app_host"`
	AppPort    int     `mapstructure:"app_port"`
	AppName    string  `mapstructure:"app_name"`
	AppVersion float64 `mapstructure:"app_version"`
	AppEnv     string  `mapstructure:"app_env"`
	// AppInstanceID     string  `mapstructure:"app_instance_id"`
}
