package configs

import (
	"time"

	"github.com/saeedjhn/go-backend-clean-arch/internal/infrastructure/logger"
	"github.com/saeedjhn/go-backend-clean-arch/internal/infrastructure/persistance/cache/redis"
	"github.com/saeedjhn/go-backend-clean-arch/internal/infrastructure/persistance/db/mongo"
	"github.com/saeedjhn/go-backend-clean-arch/internal/infrastructure/persistance/db/mysql"
	"github.com/saeedjhn/go-backend-clean-arch/internal/infrastructure/persistance/db/pq"
	"github.com/saeedjhn/go-backend-clean-arch/internal/service/authservice"
)

const (
	Development Env = "development"
	Production  Env = "production"
)

type Application struct {
	Env                     Env           `mapstructure:"env"`
	Debug                   bool          `mapstructure:"debug"`
	EntropyPassword         float64       `mapstructure:"entropy_password"`
	GracefulShutdownTimeout time.Duration `mapstructure:"graceful_shutdown_timeout"`
}

type Pprof struct {
	Port              string        `mapstructure:"port"`
	ReadTimeout       time.Duration `mapstructure:"read_timeout"`
	ReadHeaderTimeout time.Duration `mapstructure:"read_header_timeout"`
	WriteTimeout      time.Duration `mapstructure:"write_timeout"`
	IdleTimeout       time.Duration `mapstructure:"idle_timeout"`
}

type HTTPServer struct {
	Port    string        `mapstructure:"port"`
	Timeout time.Duration `mapstructure:"timeout"`
}

type GRPCServer struct {
	Network string        `mapstructure:"network"`
	Port    string        `mapstructure:"port"`
	Timeout time.Duration `mapstructure:"timeout"`
}

type Config struct {
	Application Application        `mapstructure:"application"`
	Pprof       Pprof              `mapstructure:"pprof"`
	HTTPServer  HTTPServer         `mapstructure:"http_server"`
	GRPCServer  GRPCServer         `mapstructure:"grpc_server"`
	Logger      logger.Config      `mapstructure:"logger"`
	Mysql       mysql.Config       `mapstructure:"mysql"`
	Postgres    pq.Config          `mapstructure:"postgres"`
	Redis       redis.Config       `mapstructure:"redis"`
	Mongo       mongo.Config       `mapstructure:"mongo"`
	Auth        authservice.Config `mapstructure:"auth"`
}
