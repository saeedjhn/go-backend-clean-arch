package supervisor

var defaultOptions = ProcessOption{
	Recover:         true,
	RetryInterval:   ProcessRetryInterval,
	RecoverInterval: ProcessRecoverInterval,
	RecoverCount:    ProcessRecoverCount,
	RetryCount:      ProcessRetryCount,
	IsFatal:         true,
}
