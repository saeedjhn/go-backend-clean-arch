package redis

import "time"

type Config struct {
	Host               string        `mapstructure:"host"`
	Port               string        `mapstructure:"port"`
	Password           string        `mapstructure:"password"`
	DB                 int           `mapstructure:"db"`
	PoolSize           int           `mapstructure:"pool_size"`
	DialTimeout        time.Duration `mapstructure:"dial_timeout"`
	ReadTimeout        time.Duration `mapstructure:"read_timeout"`
	WriteTimeout       time.Duration `mapstructure:"write_timeout"`
	PoolTimeout        time.Duration `mapstructure:"pool_timeout"`
	IdleCheckFrequency time.Duration `mapstructure:"idle_check_frequency"`
}
