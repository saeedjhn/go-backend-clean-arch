package supervisor

import "time"

const (
	ProcessRetryCount              = 3
	ProcessRetryInterval           = 10 * time.Second
	ProcessRecoverCount            = 10
	ProcessRecoverInterval         = 2 * time.Second
	DefaultGracefulShutdownTimeout = 5 * time.Second
	LogNSSupervisor                = "supervisor"
)
