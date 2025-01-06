package user_test

import (
	"errors"
	"log"
	"testing"

	"github.com/saeedjhn/go-backend-clean-arch/internal/repository/mysql/user/user_test/setup"
)

var (
	errUserNotFound = errors.New("user not found")
	errUnexpected   = errors.New("unexpected error")
)

func TestMain(m *testing.M) {
	dbContainer := setup.NewDBContainer(setup.ContainerMaxWait, setup.ContainerExpireInSeconds)
	dbContainer.SetConfig(setup.GetDBConfig())
	if err := dbContainer.Start(); err != nil {
		log.Panicf("Failed to start MySQL dbContainer: %v", err)
	}

	defer func(testContainer *setup.DBContainer) {
		if err := testContainer.Stop(); err != nil {
			log.Printf("Failed to stop MySQL dbContainer: %v", err)
		}
	}(dbContainer)

	downMigrate, err := setup.MigrateDB()
	if err != nil {
		log.Panicf("Migration failed: %v", err)
	}
	defer func() {
		if errM := downMigrate(); errM != nil {
			log.Printf("Failed to rollback migrations: %v", errM)
		}
	}()

	downSeed, err := setup.SeedDB()
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
