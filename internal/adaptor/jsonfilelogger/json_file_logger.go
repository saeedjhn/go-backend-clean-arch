package jsonfilelogger

import (
	"errors"
	"os"

	"github.com/saeedjhn/go-backend-clean-arch/internal/contract"
	"go.uber.org/zap"
)

var _ contract.Logger = (*L)(nil)

type L struct {
	envStrategy EnvironmentStrategy
	sugar       *zap.SugaredLogger
}

func New(envStrategy EnvironmentStrategy) *L {
	return &L{envStrategy: envStrategy}
}

func (l *L) SetStrategy(strategy EnvironmentStrategy) {
	l.envStrategy = strategy
}

func (l *L) Configure() *L {
	strategy := l.getStrategy()
	l.sugar = strategy.Sugar()

	if err := l.sugar.Sync(); err != nil && !errors.Is(err, os.ErrInvalid) {
		l.sugar.Error(err)
	}

	return l
}

func (l *L) Debug(args ...interface{}) {
	l.sugar.Debug(args...)
}

func (l *L) Debugf(msg string, args ...interface{}) {
	l.sugar.Debugf(msg, args...)
}

func (l *L) Info(args ...interface{}) {
	l.sugar.Info(args...)
}

func (l *L) Infow(msg string, keysAndValues ...interface{}) {
	l.sugar.Infow(msg, keysAndValues...)
}

func (l *L) Infof(msg string, args ...interface{}) {
	l.sugar.Infof(msg, args...)
}

func (l *L) Warn(args ...interface{}) {
	l.sugar.Warn(args...)
}

func (l *L) Warnf(msg string, args ...interface{}) {
	l.sugar.Warnf(msg, args...)
}

func (l *L) Error(args ...interface{}) {
	l.sugar.Error(args...)
}

func (l *L) Errorf(msg string, args ...interface{}) {
	l.sugar.Errorf(msg, args...)
}

func (l *L) Errorw(msg string, keysAndValues ...interface{}) {
	l.sugar.Errorw(msg, keysAndValues...)
}

func (l *L) DPanic(args ...interface{}) {
	l.sugar.DPanic(args...)
}

func (l *L) DPanicf(msg string, args ...interface{}) {
	l.sugar.DPanicf(msg, args...)
}

func (l *L) Fatal(args ...interface{}) {
	l.sugar.Fatal(args...)
}

func (l *L) Fatalf(msg string, args ...interface{}) {
	l.sugar.Fatalf(msg, args...)
}

func (l *L) getStrategy() *zap.Logger {
	return l.envStrategy.CreateLogger()
}
