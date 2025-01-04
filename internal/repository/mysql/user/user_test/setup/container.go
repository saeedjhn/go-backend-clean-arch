package setup

import (
	"errors"
	"fmt"
	"time"

	"github.com/ory/dockertest/v3"

	_ "github.com/go-sql-driver/mysql" // Blank import without comment
	"github.com/ory/dockertest/v3/docker"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/persistance/db/mysql"
)

const (
	repository               = "mysql"
	tag                      = "9.1.0"
	exposedPort              = "3306"
	containerPortProtocol    = "3306/tcp"
	ContainerMaxWait         = 120 * time.Second
	ContainerExpireInSeconds = 120

	maxIdleConns    = 15
	maxOpenConns    = 100
	connMaxLiftTime = 5 * time.Second
)

func GetDBConfig() mysql.Config {
	return mysql.Config{
		Host:            "localhost",
		Port:            "3306",
		Username:        "admin",
		Password:        "password123",
		Database:        "your_db",
		SSLMode:         "",
		MaxIdleConns:    maxIdleConns,
		MaxOpenConns:    maxOpenConns,
		ConnMaxLiftTime: connMaxLiftTime,
	}
}

type DBContainer struct {
	dockerPool      *dockertest.Pool
	mysqlResource   *dockertest.Resource
	mysqlConfig     mysql.Config
	mysqlDBConn     *mysql.Mysql
	maxWait         time.Duration
	expiryInSeconds uint
}

func NewDBContainer(maxWait time.Duration, expireInSeconds uint) *DBContainer {
	return &DBContainer{
		maxWait:         maxWait,
		expiryInSeconds: expireInSeconds,
	}
}

func (t *DBContainer) GetConnection() (*mysql.Mysql, error) {
	if t.mysqlDBConn == nil {
		return nil, errors.New("did you forgot to start the test container")
	}

	return t.mysqlDBConn, nil
}

func (t *DBContainer) SetConfig(config mysql.Config) *DBContainer {
	t.mysqlConfig = config

	return t
}

func (t *DBContainer) GetConfig() mysql.Config {
	return t.mysqlConfig
}

func (t *DBContainer) Start() error {
	var err error

	t.dockerPool, err = dockertest.NewPool("")
	t.dockerPool.MaxWait = t.maxWait
	if err != nil {
		return fmt.Errorf("could not construct pool: %w", err)
	}

	err = t.dockerPool.Client.Ping()
	if err != nil {
		return fmt.Errorf("could not connect to Docker: %w", err)
	}

	t.mysqlResource, err = t.dockerPool.RunWithOptions(
		&dockertest.RunOptions{
			Repository: repository,
			Tag:        tag,
			Env: []string{
				"MYSQL_ROOT_PASSWORD=root",
				fmt.Sprintf("MYSQL_DATABASE=%s", t.mysqlConfig.Database),
				fmt.Sprintf("MYSQL_USER=%s", t.mysqlConfig.Username),
				fmt.Sprintf("MYSQL_PASSWORD=%s", t.mysqlConfig.Password),
			},
			ExposedPorts: []string{exposedPort},
			PortBindings: map[docker.Port][]docker.PortBinding{
				containerPortProtocol: {
					docker.PortBinding{
						HostIP:   "0.0.0.0",
						HostPort: t.mysqlConfig.Port,
					},
				},
			},
		}, func(config *docker.HostConfig) {
			config.AutoRemove = true // Ensure the container is removed after the test.
			config.RestartPolicy = docker.RestartPolicy{
				Name: "no", // Do not automatically restart the container.
			}
		},
	)
	if err != nil {
		return fmt.Errorf("could not start mysql resource: %w", err)
	}

	if err = t.mysqlResource.Expire(t.expiryInSeconds); err != nil {
		return fmt.Errorf("couldn't set mysql container expiration: %w", err)
	}

	if err = t.dockerPool.Retry(func() error {
		t.mysqlDBConn = mysql.New(t.mysqlConfig)
		if err = t.mysqlDBConn.ConnectTo(); err != nil {
			return err
		}

		return t.mysqlDBConn.Conn().Ping()
	}); err != nil {
		return fmt.Errorf("could not connect to database: %w", err)
	}

	return nil
}

func (t *DBContainer) Stop() error {
	err := t.mysqlDBConn.Conn().Close()
	if err != nil {
		return fmt.Errorf("could not close mysql db connection: %w", err)
	}

	if err = t.dockerPool.Purge(t.mysqlResource); err != nil {
		return fmt.Errorf("could not purge mysql resource: %w", err)
	}

	return nil
}
