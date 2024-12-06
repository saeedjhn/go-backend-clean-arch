package jsonfilelogger_test

import (
	"testing"
	"time"

	"github.com/saeedjhn/go-backend-clean-arch/internal/adaptor/jsonfilelogger"
)

var _environment = "development"
var _cfg = jsonfilelogger.Config{
	MaxSize:          10,
	MaxBackups:       10,
	MaxAge:           30,
	LocalTime:        true,
	Compress:         false,
	Console:          true,
	EnableCaller:     true,
	EnableStacktrace: true,
	Level:            "debug",
}

func TestRun(t *testing.T) {
	environment := "production"

	var envStrategy jsonfilelogger.EnvironmentStrategy
	if environment == "development" {
		envStrategy = jsonfilelogger.NewDevelopmentStrategy(_cfg)
	} else {
		envStrategy = jsonfilelogger.NewProductionStrategy(_cfg)
	}

	l := jsonfilelogger.New(envStrategy)
	l.Configure()

	l.Info("-- OK --")

	l.Infow("message", "KEY", map[string]interface{}{
		"foo": "FOO",
	})

	l.Errorf("msg: %+v", map[string]string{"foo": "FOO"})
	l.Errorw("msg1", "KEY1", map[string]string{"bar": "BAR"})

}

func TestTime(t *testing.T) {
	t.Log(time.Now().Format("2006-01-02_15-04-05"))
}
