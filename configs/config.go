package configs

type Application struct {
	Env string `mapstructure:"ENV"`
}

type HTTPServer struct {
	Port    string `mapstructure:"PORT"`
	Timeout int    `mapstructure:"TIMEOUT"`
}

type Config struct {
	Application Application `mapstructure:"APPLICATION"`
	HTTPServer  HTTPServer  `mapstructure:"HTTPSERVER"`
}
