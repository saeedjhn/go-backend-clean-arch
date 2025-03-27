package rmqpc_test

import (
	"log"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/saeedjhn/go-backend-clean-arch/internal/adaptor/rmqpc"

	setuptest "github.com/saeedjhn/go-backend-clean-arch/test/setup_test"
)

const (
	connectionName = "conn-1"

	hostMachineIP   = "0.0.0.0"
	hostMachinePort = "5672"

	name                 = "test_rabbitmq"
	repository           = "rabbitmq"
	tag                  = "4.0.5-management"
	containerExposedPort = "5672"
	containerProtocol    = "tcp"

	containerMaxWait         = 120 * time.Second
	containerExpireInSeconds = 120
)

var (
	_myRabbitMQConnectionConfig rmqpc.ConnectionConfig  //nolint:gochecknoglobals // nothing
	_myRabbitMQ                 *rmqpc.Connection       //nolint:gochecknoglobals // nothing
	_configPath                 = "testdata/config.yml" //nolint:gochecknoglobals // nothing
)

func TestMain(m *testing.M) {
	wd, err := os.Getwd()
	if err != nil {
		log.Panicf("error getting current working directory: %v", err)
	}

	_myRabbitMQConnectionConfig, err = setuptest.LoadConfig[rmqpc.ConnectionConfig](filepath.Join(wd, _configPath))
	if err != nil {
		log.Panicf("failed to load config: %v", err)
	}

	_myRabbitMQ = setuptest.NewRabbitMQ(connectionName, _myRabbitMQConnectionConfig)

	dbContainer := setuptest.NewRunContainer(setuptest.RunContainerOptions{
		Name:       name,
		Repository: repository,
		Tag:        tag,
		Env: map[string]string{
			"RABBITMQ_DEFAULT_USER": "admin",
			"RABBITMQ_DEFAULT_PASS": "password123",
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
		return _myRabbitMQ.ConnectRaw()
	}); err != nil {
		log.Panicf("failed to start container: %v", err)
	}

	defer func(dbContainer *setuptest.RunContainer, fn func() error) {
		err = dbContainer.Stop(fn)
		if err != nil {
			log.Printf("failed to stop container: %v", err)
		}
	}(dbContainer, func() error {
		return _myRabbitMQ.Close()
	})

	m.Run()
}
