package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Logger struct {
	logger *zap.Logger
}

func New(config Config) *Logger {
	// TODO - add & check production/development config variable
	// var logger, _ = zap.NewProduction()

	configZap := zap.NewProductionEncoderConfig()
	configZap.EncodeTime = zapcore.ISO8601TimeEncoder
	defaultEncoder := zapcore.NewJSONEncoder(configZap)

	writer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   config.Filename,   // "./logs/log.json"
		MaxSize:    config.MaxSize,    // megabytes
		MaxBackups: config.MaxBackups, // megabytes
		MaxAge:     config.MaxAge,     // days
		LocalTime:  config.LocalTime,  // T/F
		Compress:   config.Compress,   // T/F
	})

	stdOutWriter := zapcore.AddSync(os.Stdout)
	defaultLogLevel := zapcore.InfoLevel
	core := zapcore.NewTee(
		zapcore.NewCore(defaultEncoder, writer, defaultLogLevel),
		zapcore.NewCore(defaultEncoder, stdOutWriter, zap.DebugLevel),
	)

	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))

	return &Logger{logger: logger}
}

func (l *Logger) Set() *zap.Logger {
	return l.logger
}
