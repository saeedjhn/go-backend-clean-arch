package setup

import (
	"fmt"
	"sync"

	"github.com/saeedjhn/go-backend-clean-arch/pkg/persistance/db/mysql"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/persistance/db/mysql/migratormysql"
)

var (
	once     sync.Once
	instance *mysql.Mysql

	migrationPath = "./testdata/migrations"
	seedPath      = "./testdata/seed"
)

func NewDB(config mysql.Config) (*mysql.Mysql, error) {
	db := mysql.New(config)
	if err := db.ConnectTo(); err != nil {
		return nil, err
	}

	return db, nil
}

func NewDBSingleton(config mysql.Config) (*mysql.Mysql, error) {
	var initErr error

	once.Do(func() {
		db := mysql.New(config)
		if err := db.ConnectTo(); err != nil {
			initErr = err
			return
		}

		instance = db
	})

	if initErr != nil {
		return nil, initErr
	}

	return instance, nil
}

// func InsertDB() {
// 	conn, _ := NewDBSingleton(GetDBConfig())
// }

// func GetByMobileDB(mobile string) (entity.User, error) {
//
// }

func MigrateDB() (func() error, error) {
	conn, err := NewDBSingleton(GetDBConfig())
	if err != nil {
		return nil, fmt.Errorf("failed to create DB connection: %w", err)
	}

	migrations := migratormysql.New(conn, migratormysql.Config{
		MigrationPath:   migrationPath,
		MigrationDBName: "grop_migrations",
	})

	if err = migrations.Up(); err != nil {
		return nil, fmt.Errorf("failed to apply migrations: %w", err)
	}

	return func() error {
		if errD := migrations.Down(); errD != nil {
			return fmt.Errorf("failed to rollback migrations: %w", errD)
		}
		return nil
	}, nil
}

func SeedDB() (func() error, error) {
	conn, err := NewDBSingleton(GetDBConfig())
	if err != nil {
		return nil, fmt.Errorf("failed to create DB connection: %w", err)
	}

	seed := migratormysql.New(conn, migratormysql.Config{
		MigrationPath:   seedPath,
		MigrationDBName: "test_migrations",
	})

	if err = seed.Up(); err != nil {
		return nil, fmt.Errorf("failed to apply seed data: %w", err)
	}

	return func() error {
		if errD := seed.Down(); errD != nil {
			return fmt.Errorf("failed to rollback seed data: %w", errD)
		}
		return nil
	}, nil
}
