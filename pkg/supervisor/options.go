package supervisor

var defaultOptions = ProcessOption{ //nolint:gochecknoglobals // nothing
	Recover:         true,
	RecoverInterval: ProcessRecoverInterval,
	RecoverCount:    ProcessRecoverCount,
	RetryInterval:   ProcessRetryInterval,
	RetryCount:      ProcessRetryCount,
	IsFatal:         true,
}
