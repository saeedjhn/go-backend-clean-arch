package jsonfilelogger

import (
	"fmt"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type DevelopmentStrategy struct {
	cfg Config
}

func NewDevelopmentStrategy(cfg Config) *DevelopmentStrategy {
	return &DevelopmentStrategy{cfg: cfg}
}

func (d *DevelopmentStrategy) CreateLogger() *zap.Logger {
	fileName := d.generateLogFilename()

	writer := d.createLogWriter(fileName)

	encoderCfg := d.createEncoderConfig()

	encoder := d.CreateEncoder(encoderCfg)

	zapCore := d.createCore(encoder, writer)

	zapOption := d.createOption()

	logger := zap.New(zapCore, zapOption...)

	return logger
}

func (d *DevelopmentStrategy) generateLogFilename() string {
	fileName := fmt.Sprintf(
		"log_%s_%d.log",
		time.Now().Format("2006-01-02_15-04-05"),
		time.Now().Unix(),
	)

	return fileName
}

func (d *DevelopmentStrategy) createLogWriter(fileName string) zapcore.WriteSyncer {
	writer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    d.cfg.MaxSize,    // megabytes
		MaxBackups: d.cfg.MaxBackups, // megabytes
		MaxAge:     d.cfg.MaxAge,     // days
		LocalTime:  d.cfg.LocalTime,  // T/F
		Compress:   d.cfg.Compress,   // T/F
	})

	return writer
}

func (d *DevelopmentStrategy) createEncoderConfig() zapcore.EncoderConfig {
	encoderCfg := zap.NewDevelopmentEncoderConfig()

	encoderCfg.LevelKey = "LEVEL"
	encoderCfg.CallerKey = "CALLER"
	encoderCfg.TimeKey = "TS"
	encoderCfg.NameKey = "NAME"
	encoderCfg.MessageKey = "MESSAGE"

	return encoderCfg
}

func (d *DevelopmentStrategy) CreateEncoder(encoderCfg zapcore.EncoderConfig) zapcore.Encoder {
	return zapcore.NewJSONEncoder(encoderCfg)
}

func (d *DevelopmentStrategy) createCore(
	defaultEncoder zapcore.Encoder,
	writer zapcore.WriteSyncer,
) zapcore.Core {
	var zapCore []zapcore.Core
	zapCore = append(zapCore,
		zapcore.NewCore(
			defaultEncoder, writer, zap.NewAtomicLevelAt(d.getLoggerLevel()),
		),
	)

	if d.cfg.Console {
		zapCore = append(zapCore,
			zapcore.NewCore(
				defaultEncoder, zapcore.AddSync(os.Stdout), zap.NewAtomicLevelAt(d.getLoggerLevel()),
			),
		)
	}

	core := zapcore.NewTee(zapCore...)

	return core
}

func (d *DevelopmentStrategy) createOption() []zap.Option {
	var zapOption []zap.Option

	if d.cfg.EnableCaller {
		zapOption = append(zapOption, zap.AddCaller(), zap.AddCallerSkip(1))
	}
	if d.cfg.EnableStacktrace {
		zapOption = append(zapOption, zap.AddStacktrace(zapcore.ErrorLevel))
	}

	return zapOption
}

func (d *DevelopmentStrategy) getLoggerLevel() zapcore.Level {
	var loggerLevelMap = map[string]zapcore.Level{
		"debug":  zapcore.DebugLevel,
		"info":   zapcore.InfoLevel,
		"warn":   zapcore.WarnLevel,
		"error":  zapcore.ErrorLevel,
		"dpanic": zapcore.DPanicLevel,
		"panic":  zapcore.PanicLevel,
		"fatal":  zapcore.FatalLevel,
	}

	level, exist := loggerLevelMap[d.cfg.Level]
	if !exist {
		return zapcore.DebugLevel
	}

	return level
}
