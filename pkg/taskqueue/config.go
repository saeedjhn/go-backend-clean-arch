package taskqueue

import "time"

type Config struct {
	Concurrency int
	RedisConfig RedisConfig
}

type RedisConfig struct {
	Network      string
	Host         string
	Port         string
	Password     string
	DB           int
	PoolSize     int
	DialTimeout  time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	PoolTimeout  time.Duration
}
