package bootstrap

import (
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/infrastructure/persistance/db/pq"
	"log"
)

func newPostgresConnection(config pq.Config) pq.DB {
	//func newPostgresConnection(config ConfigPostgresqlConnection) postgresql.DB {
	return pq.New(config)
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

func closePostgresConnection(postgresDB pq.DB) {
	err := postgresDB.Conn().Close()

	if err != nil {
		log.Fatalf("don`t close postgresql connection: %s", err.Error())
	}
}
