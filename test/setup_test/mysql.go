package setuptest

import (
	"fmt"
	"sync"

	"github.com/saeedjhn/go-backend-clean-arch/pkg/persistance/db/mysql"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/persistance/db/mysql/migratormysql"
)

const (
	_migrationsDBName = "group_migrations"
	_seedDBName       = "seed_migrations"
)

var (
	once     sync.Once    //nolint:gochecknoglobals // nothing
	instance *mysql.Mysql //nolint:gochecknoglobals // nothing
)

type MySQLSeedOptions struct {
	config         mysql.Config
	HostIPOverride string
	SeedPath       string
	SeedDBName     string
}

type MySQLMigrateOptions struct {
	config          mysql.Config
	HostIPOverride  string
	MigrationPath   string
	MigrationDBName string
}

func NewMySQLDB(config mysql.Config) (*mysql.Mysql, error) {
	db := mysql.New(config)
	if err := db.ConnectTo(); err != nil {
		return nil, err
	}

	return db, nil
}

func NewMySQLDBSingleton(config mysql.Config) (*mysql.Mysql, error) {
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
// 	conn, _ := NewMySQLDBSingleton(GetDBConfig())
// }

// func GetByMobileDB(mobile string) (entity.User, error) {
//
// }

func MySQLMigrateDB(options MySQLMigrateOptions) (func() error, error) {
	if len(options.HostIPOverride) != 0 {
		options.config.Host = options.HostIPOverride
	}

	conn, err := NewMySQLDBSingleton(options.config)
	if err != nil {
		return nil, fmt.Errorf("failed to create DB connection: %w", err)
	}

	migrations := migratormysql.New(conn, migratormysql.Config{
		MigrationPath: options.MigrationPath,
		MigrationDBName: func() string {
			if len(options.MigrationDBName) == 0 {
				return _migrationsDBName
			}
			return options.MigrationDBName
		}(),
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

func SeedDB(options MySQLSeedOptions) (func() error, error) {
	if len(options.HostIPOverride) != 0 {
		options.config.Host = options.HostIPOverride
	}

	conn, err := NewMySQLDBSingleton(options.config)
	if err != nil {
		return nil, fmt.Errorf("failed to create DB connection: %w", err)
	}

	seed := migratormysql.New(conn, migratormysql.Config{
		MigrationPath: options.SeedPath,
		MigrationDBName: func() string {
			if len(options.SeedDBName) == 0 {
				return _seedDBName
			}
			return options.SeedDBName
		}(),
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
