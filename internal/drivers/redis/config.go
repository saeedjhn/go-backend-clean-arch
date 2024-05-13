package redis

import "time"

type Config struct {
	Host               string        `env:"HOST"`
	Port               string        `env:"PORT"`
	Password           string        `env:"PASSWORD"`
	DB                 int           `env:"DB"`
	PoolSize           int           `env:"POOLSIZE"`
	DialTimeout        time.Duration `env:"DIALTIMEOUT"`
	ReadTimeout        time.Duration `env:"READTIMEOUT"`
	WriteTimeout       time.Duration `env:"WRITETIMEOUT"`
	PoolTimeout        time.Duration `env:"POOLTIMEOUT"`
	IdleCheckFrequency time.Duration `env:"IDLECHECKFREQUENCY"`
}
