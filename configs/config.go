package configs

import (
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/infrastructure/cache/redis"
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/infrastructure/db/mongo"
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/infrastructure/db/mysql"
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/infrastructure/db/postgresql"
	"time"
)

type Application struct {
	Env string `mapstructure:"env"`
}

type HTTPServer struct {
	Port                    string        `mapstructure:"port"`
	Timeout                 time.Duration `mapstructure:"timeout"`
	GracefulShutdownTimeout time.Duration `mapstructure:"graceful_shutdown_timeout"`
}

type Config struct {
	Application Application       `mapstructure:"application"`
	HTTPServer  HTTPServer        `mapstructure:"http_server"`
	Mysql       mysql.Config      `mapstructure:"mysql"`
	Postgresql  postgresql.Config `mapstructure:"postgresql"`
	Redis       redis.Config      `mapstructure:"redis"`
	Mongo       mongo.Config      `mapstructure:"mongo"`
}
