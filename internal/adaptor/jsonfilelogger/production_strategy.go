package jsonfilelogger

import (
	"fmt"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type ProductionStrategy struct {
	cfg Config
}

func NewProductionStrategy(cfg Config) *ProductionStrategy {
	return &ProductionStrategy{cfg: cfg}
}

func (d *ProductionStrategy) CreateLogger() *zap.Logger {
	fileName := d.generateLogFilename()

	writer := d.createLogWriter(fileName)

	encoderCfg := d.createEncoderConfig()

	encoder := d.CreateEncoder(encoderCfg)

	zapCore := d.createCore(encoder, writer)

	zapOption := d.createOption()

	logger := zap.New(zapCore, zapOption...)

	return logger
}

func (d *ProductionStrategy) generateLogFilename() string {
	fileName := fmt.Sprintf(
		"log_%s_%d.log",
		time.Now().Format("2006-01-02_15-04-05"),
		time.Now().Unix(),
	)

	return fileName
}

func (d *ProductionStrategy) createLogWriter(fileName string) zapcore.WriteSyncer {
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

func (d *ProductionStrategy) createEncoderConfig() zapcore.EncoderConfig {
	return zap.NewProductionEncoderConfig()
}

func (d *ProductionStrategy) CreateEncoder(encoderCfg zapcore.EncoderConfig) zapcore.Encoder {
	return zapcore.NewJSONEncoder(encoderCfg)
}

func (d *ProductionStrategy) createCore(
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

func (d *ProductionStrategy) createOption() []zap.Option {
	var zapOption []zap.Option

	if d.cfg.EnableCaller {
		zapOption = append(zapOption, zap.AddCaller(), zap.AddCallerSkip(1))
	}
	if d.cfg.EnableStacktrace {
		zapOption = append(zapOption, zap.AddStacktrace(zapcore.ErrorLevel))
	}

	return zapOption
}

func (d *ProductionStrategy) getLoggerLevel() zapcore.Level {
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
