package bootstrap

import (
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/infrastructure/persistance/db/mysql"
	"log"
)

//type ConfigMysqlConnection struct {
//	Host            string        `mapstructure:"host"`
//	Port            string        `mapstructure:"port"`
//	Username        string        `mapstructure:"username"`
//	Password        string        `mapstructure:"password"`
//	Database        string        `mapstructure:"database"`
//	SSLMode         string        `mapstructure:"ssl_mode"`
//	MaxIdleConns    int           `mapstructure:"max_idle_conns"`
//	MaxOpenConns    int           `mapstructure:"max_open_conns"`
//	ConnMaxLiftTime time.Duration `mapstructure:"conn_max_life_time"`
//}

// func newMysqlConnection(config ConfigMysqlConnection) mysql.DB {
func newMysqlConnection(config mysql.Config) mysql.DB {
	return mysql.New(config)
}

func closeMysqlConnection(mysqlDB mysql.DB) {
	err := mysqlDB.Conn().Close()

	if err != nil {
		log.Fatalf("don`t close mysql connection: %s", err.Error())
	}
}
