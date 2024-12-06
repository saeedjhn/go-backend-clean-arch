package contract

type Logger interface {
	Debug(args ...interface{})
	Debugf(msg string, args ...interface{})
	Info(args ...interface{})
	Infof(msg string, args ...interface{})
	Infow(msg string, keysAndValues ...interface{})
	Warn(args ...interface{})
	Warnf(msg string, args ...interface{})
	Error(args ...interface{})
	Errorf(msg string, args ...interface{})
	Errorw(msg string, keysAndValues ...interface{})
	DPanic(args ...interface{})
	DPanicf(msg string, args ...interface{})
	Fatal(args ...interface{})
	Fatalf(msg string, args ...interface{})
}
