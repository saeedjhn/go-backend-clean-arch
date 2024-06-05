package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

type Logger struct {
	logger *zap.Logger
}

func New(config Config) *Logger {
	// TODO - add & check production/development config variable
	var logger, _ = zap.NewProduction()

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

	logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))

	return &Logger{logger: logger}
}

func (l *Logger) Set() *zap.Logger {
	return l.logger
}

// var Logger *zap.Logger
//var once = sync.Once{}
//func init() {
//	once.Do(func() {
//		// Add & check production/development config variable
//		//logger, _ = zap.NewProduction()
//
//		config := zap.NewProductionEncoderConfig()
//		config.EncodeTime = zapcore.ISO8601TimeEncoder
//		defaultEncoder := zapcore.NewJSONEncoder(config)
//		// Add to config
//		writer := zapcore.AddSync(&lumberjack.logger{
//			Filename:  "./logs/log.json",
//			LocalTime: false,
//			MaxSize:   10, // megabytes
//			//MaxBackups: 10,
//			MaxAge: 30, // days
//		})
//
//		stdOutWriter := zapcore.AddSync(os.Stdout)
//		defaultLogLevel := zapcore.InfoLevel
//		core := zapcore.NewTee(
//			zapcore.NewCore(defaultEncoder, writer, defaultLogLevel),
//			zapcore.NewCore(defaultEncoder, stdOutWriter, zap.InfoLevel),
//		)
//		//Logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
//	})
//}
// Use: logger.Logger.Named("main").Info("config", zap.Any("config", app.Config))
