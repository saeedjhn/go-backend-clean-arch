package setuptest

import (
	"fmt"
	"time"

	"github.com/ory/dockertest/v3"

	"github.com/ory/dockertest/v3/docker"
)

type BuildAndRunContainerOptions struct {
	Name            string
	ContextDir      string
	Dockerfile      string
	Env             DBEnv
	Exposed         Exposed
	PortBinding     PortBinding
	MaxWaitRetry    time.Duration
	ExpireInSeconds uint
}

type BuildAndRunContainer struct {
	options    BuildAndRunContainerOptions
	dockerPool *dockertest.Pool
	resource   *dockertest.Resource
	network    *dockertest.Network
}

func NewBuildAndRunContainer(options BuildAndRunContainerOptions) *BuildAndRunContainer {
	return &BuildAndRunContainer{
		options: options,
	}
}

func (c *BuildAndRunContainer) GetOptions() BuildAndRunContainerOptions {
	return c.options
}

func (c *BuildAndRunContainer) SetNetwork(network *dockertest.Network) *BuildAndRunContainer {
	c.network = network

	return c
}

func (c *BuildAndRunContainer) Start(fn func() error) error {
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

	c.resource, err = c.dockerPool.BuildAndRunWithBuildOptions(
		&dockertest.BuildOptions{
			ContextDir: c.options.ContextDir,
			Dockerfile: c.options.Dockerfile,
		},
		&dockertest.RunOptions{
			Name:         c.options.Name,
			Env:          c.options.Env.ToSlice(),
			ExposedPorts: []string{c.options.Exposed.Port},
			Networks: func() []*dockertest.Network {
				if c.network != nil {
					return []*dockertest.Network{c.network}
				}
				return nil
			}(),
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

func (c *BuildAndRunContainer) Stop(fn func() error) error {
	if err := fn(); err != nil {
		return fmt.Errorf("could not close %s connection: %w", c.options.Name, err)
	}

	if err := c.dockerPool.Purge(c.resource); err != nil {
		return fmt.Errorf("could not purge %s resource: %w", c.options.Name, err)
	}

	return nil
}
