package user_test

import (
	"errors"
	"log"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/saeedjhn/go-backend-clean-arch/pkg/persistance/db/mysql"
	setuptest "github.com/saeedjhn/go-backend-clean-arch/test/setup_test"
)

const (
	hostMachineIP   = "0.0.0.0"
	hostMachinePort = "3306"

	name                 = "test_mysql"
	repository           = "mysql"
	tag                  = "9.1.0"
	containerExposedPort = "3306"
	containerProtocol    = "tcp"

	containerMaxWait         = 120 * time.Second
	containerExpireInSeconds = 120
)

var (
	_myDBConfig mysql.Config            //nolint:gochecknoglobals // nothing
	_myDB       *mysql.Mysql            //nolint:gochecknoglobals // nothing
	_configPath = "testdata/config.yml" //nolint:gochecknoglobals // nothing

	migrationPath = "./testdata/migrations" //nolint:gochecknoglobals // nothing
	seedPath      = "./testdata/seed"       //nolint:gochecknoglobals // nothing

	errUserNotFound = errors.New("user not found")
	errUnexpected   = errors.New("unexpected error")
)

func TestMain(m *testing.M) {
	wd, err := os.Getwd()
	if err != nil {
		log.Panicf("error getting current working directory: %v", err)
	}

	_myDBConfig, err = setuptest.LoadConfig[mysql.Config](filepath.Join(wd, _configPath))
	if err != nil {
		log.Panicf("failed to load config: %v", err)
	}

	_myDB, err = setuptest.NewMySQLDBSingleton(_myDBConfig)
	if err != nil {
		log.Panicf("failed to new mysql connection: %v", err)
	}

	dbContainer := setuptest.NewRunContainer(setuptest.RunContainerOptions{
		Name:       name,
		Repository: repository,
		Tag:        tag,
		Env: map[string]string{
			"MYSQL_ROOT_PASSWORD": "root",
			"MYSQL_DATABASE":      "test_db",
			"MYSQL_USER":          "admin",
			"MYSQL_PASSWORD":      "password123",
		},
		Exposed: setuptest.Exposed{
			Protocol: containerProtocol,
			Port:     containerExposedPort,
		},
		PortBinding: setuptest.PortBinding{
			HostIP:   hostMachineIP,
			HostPort: hostMachinePort,
		},
		MaxWaitRetry:    containerMaxWait,
		ExpireInSeconds: containerExpireInSeconds,
	})

	if err = dbContainer.Start(func() error {
		return _myDB.Conn().Ping()
	}); err != nil {
		log.Panicf("failed to start container: %v", err)
	}

	defer func(dbContainer *setuptest.RunContainer, fn func() error) {
		err = dbContainer.Stop(fn)
		if err != nil {
			log.Printf("failed to stop container: %v", err)
		}
	}(dbContainer, func() error {
		return _myDB.Conn().Close()
	})

	downMigrate, err := setuptest.MySQLMigrateDB(setuptest.MySQLMigrateOptions{
		MigrationPath: migrationPath,
	})
	if err != nil {
		log.Panicf("migration failed: %v", err)
	}
	defer func() {
		if errM := downMigrate(); errM != nil {
			log.Printf("failed to rollback migrations: %v", errM)
		}
	}()

	downSeed, err := setuptest.SeedDB(setuptest.MySQLSeedOptions{
		SeedPath: seedPath,
	})
	if err != nil {
		log.Panicf("Seeding failed: %v", err)
	}
	defer func() {
		if errD := downSeed(); errD != nil {
			log.Printf("Failed to rollback seed data: %v", errD)
		}
	}()

	m.Run()
}
