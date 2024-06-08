package configs

import (
	"go-backend-clean-arch/internal/infrastructure/logger"
	"go-backend-clean-arch/internal/infrastructure/persistance/cache/redis"
	"go-backend-clean-arch/internal/infrastructure/persistance/db/mongo"
	"go-backend-clean-arch/internal/infrastructure/persistance/db/mysql"
	"go-backend-clean-arch/internal/infrastructure/persistance/db/pq"
	"go-backend-clean-arch/internal/usecase/authusecase"
	"time"
)

const (
	Development Env = "development"
	Production  Env = "production"
)

type Application struct {
	Env                     Env           `mapstructure:"env"`
	Debug                   bool          `mapstructure:"debug"`
	GracefulShutdownTimeout time.Duration `mapstructure:"graceful_shutdown_timeout"`
}

type HTTPServer struct {
	Port    string        `mapstructure:"port"`
	Timeout time.Duration `mapstructure:"timeout"`
}

type Config struct {
	Application Application        `mapstructure:"application"`
	HTTPServer  HTTPServer         `mapstructure:"http_server"`
	Logger      logger.Config      `mapstructure:"logger"`
	Mysql       mysql.Config       `mapstructure:"mysql"`
	Postgres    pq.Config          `mapstructure:"postgres"`
	Redis       redis.Config       `mapstructure:"redis"`
	Mongo       mongo.Config       `mapstructure:"mongo"`
	Auth        authusecase.Config `mapstructure:"auth"`
}
