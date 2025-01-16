package setuptest

import (
	"fmt"
	"time"

	"github.com/ory/dockertest/v3"

	"github.com/ory/dockertest/v3/docker"
)

type RunContainerOptions struct {
	Name            string
	Repository      string
	Tag             string
	Env             DBEnv
	Exposed         Exposed
	PortBinding     PortBinding
	Volumes         []string
	MaxWaitRetry    time.Duration
	ExpireInSeconds uint
}

type RunContainer struct {
	options    RunContainerOptions
	dockerPool *dockertest.Pool
	resource   *dockertest.Resource
	network    *dockertest.Network
}

func NewRunContainer(options RunContainerOptions) *RunContainer {
	return &RunContainer{
		options: options,
	}
}

func (c *RunContainer) SetNetwork(network *dockertest.Network) *RunContainer {
	c.network = network

	return c
}

func (c *RunContainer) Start(fn func() error) error {
	var err error

	c.dockerPool, err = dockertest.NewPool("")
	c.dockerPool.MaxWait = c.options.MaxWaitRetry

	if err != nil {
		return fmt.Errorf("could not construct pool: %w", err)
	}

	err = c.dockerPool.Client.Ping()
	if err != nil {
		return fmt.Errorf("could not connect to Docker: %w", err)
	}

	c.resource, err = c.dockerPool.RunWithOptions(
		&dockertest.RunOptions{
			Name:       c.options.Name,
			Repository: c.options.Repository,
			Tag:        c.options.Tag,
			Env:        c.options.Env.ToSlice(),
			Networks: func() []*dockertest.Network {
				if c.network != nil {
					return []*dockertest.Network{c.network}
				}
				return nil
			}(),
			ExposedPorts: []string{c.options.Exposed.Port},
			PortBindings: map[docker.Port][]docker.PortBinding{
				docker.Port(fmt.Sprintf("%s/%s", c.options.Exposed.Port, c.options.Exposed.Protocol)): {
					docker.PortBinding{
						HostIP:   c.options.PortBinding.HostIP,
						HostPort: c.options.PortBinding.HostPort,
					},
				},
			},
		}, func(config *docker.HostConfig) {
			config.AutoRemove = true // Ensure the container is removed after the test.
			config.RestartPolicy = docker.RestartPolicy{
				Name: "no", // Do not automatically restart the container.
			}

			if len(c.options.Volumes) != 0 {
				config.Binds = c.options.Volumes
			}
		},
	)
	if err != nil {
		return fmt.Errorf("could not start %s resource: %w", c.options.Name, err)
	}

	if err = c.resource.Expire(c.options.ExpireInSeconds); err != nil {
		return fmt.Errorf("couldn'c set %s container expiration: %w", c.options.Name, err)
	}

	if err = c.dockerPool.Retry(fn); err != nil {
		return fmt.Errorf("could not connect to %s container: %w", c.options.Name, err)
	}

	return nil
}

func (c *RunContainer) Stop(fn func() error) error {
	if err := fn(); err != nil {
		return fmt.Errorf("could not close %s connection: %w", c.options.Name, err)
	}

	if err := c.dockerPool.Purge(c.resource); err != nil {
		return fmt.Errorf("could not purge %s resource: %w", c.options.Name, err)
	}

	return nil
}
