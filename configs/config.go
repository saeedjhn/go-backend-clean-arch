package configs

import (
	"time"

	"github.com/saeedjhn/go-backend-clean-arch/internal/adaptor/oteltracer"

	"github.com/saeedjhn/go-backend-clean-arch/pkg/logger"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/persistance/cache/redis"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/persistance/db/mongo"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/persistance/db/mysql"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/persistance/db/pq"

	"github.com/saeedjhn/go-backend-clean-arch/internal/usecase/authusecase"
)

type Env string

const (
	Development Env = "development"
	Production  Env = "production"
)

func (e Env) String() string {
	return string(e)
}

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

type CORS struct {
	AllowOrigins     []string `mapstructure:"allow_origins"`
	AllowMethods     []string `mapstructure:"allow_methods"`
	AllowHeaders     []string `mapstructure:"allow_headers"`
	AllowCredentials bool     `mapstructure:"allow_credentials"`
}

type Register struct {
	Debug bool `mapstructure:"debug"`
}

type Config struct {
	Application Application        `mapstructure:"application"`
	Pprof       Pprof              `mapstructure:"pprof"`
	HTTPServer  HTTPServer         `mapstructure:"http_server"`
	GRPCServer  GRPCServer         `mapstructure:"grpc_server"`
	CORS        CORS               `mapstructure:"cors"`
	Tracer      oteltracer.Config  `mapstructure:"tracer"`
	Logger      logger.Config      `mapstructure:"logger"`
	Mysql       mysql.Config       `mapstructure:"mysql"`
	Postgres    pq.Config          `mapstructure:"postgres"`
	Redis       redis.Config       `mapstructure:"redis"`
	Mongo       mongo.Config       `mapstructure:"mongo"`
	Auth        authusecase.Config `mapstructure:"auth"`
	Register    Register           `mapstructure:"register"`
}
