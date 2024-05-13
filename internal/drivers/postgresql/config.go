package postgresql

import "time"

type Config struct {
	Host            string        `mapstructure:"HOST"`
	Port            string        `mapstructure:"PORT"`
	Username        string        `mapstructure:"USERNAME"`
	Password        string        `mapstructure:"PASSWORD"`
	Database        string        `mapstructure:"DATABASE"`
	SSLMode         string        `mapstructure:"SSLMODE"`
	MaxIdleConns    int           `mapstructure:"MAXIDLECONNS"`
	MaxOpenConns    int           `mapstructure:"MAXOPENCONNS"`
	ConnMaxLiftTime time.Duration `mapstructure:"CONNMAXLIFETIME"`
}
