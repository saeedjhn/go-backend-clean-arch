package bootstrap

import (
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/infrastructure/persistance/db/postgresql"
	"log"
)

//type ConfigPostgresqlConnection struct {
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

func newPostgresqlConnection(config postgresql.Config) postgresql.DB {
	//func newPostgresqlConnection(config ConfigPostgresqlConnection) postgresql.DB {
	return postgresql.New(config)
	//	uri := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=Asia/Tehran",
	//		config.Host, config.Port, config.Username, config.Password,
	//		config.Database, config.SSLMode)

	//db := postgresql.New(uri)
	//
	//db.Conn().SetMaxIdleConns(config.MaxIdleConns)
	//db.Conn().SetMaxOpenConns(config.MaxOpenConns)
	//db.Conn().SetConnMaxLifetime(config.ConnMaxLiftTime * time.Second)
	//
	//return db
}

func closePostgresqlConnection(postgresqlDB postgresql.DB) {
	err := postgresqlDB.Conn().Close()

	if err != nil {
		log.Fatalf("don`t close postgresql connection: %s", err.Error())
	}
}
