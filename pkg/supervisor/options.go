package supervisor

var defaultOptions = ProcessOption{ //nolint:gochecknoglobals // nothing
	Recover:         true,
	RetryInterval:   ProcessRetryInterval,
	RecoverInterval: ProcessRecoverInterval,
	RecoverCount:    ProcessRecoverCount,
	RetryCount:      ProcessRetryCount,
	IsFatal:         true,
}
