package redis

import "time"

type Config struct {
	Host               string        `env:"host"`
	Port               string        `env:"port"`
	Password           string        `env:"password"`
	DB                 int           `env:"db"`
	PoolSize           int           `env:"pool_size"`
	DialTimeout        time.Duration `env:"dial_timeout"`
	ReadTimeout        time.Duration `env:"read_timeout"`
	WriteTimeout       time.Duration `env:"write_timeout"`
	PoolTimeout        time.Duration `env:"pool_timeout"`
	IdleCheckFrequency time.Duration `env:"idle_check_frequency"`
}
