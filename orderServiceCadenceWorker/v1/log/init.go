package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

const (
	Development = "dev"
	Production  = "prod"
)

var (
	logger *zap.Logger
)

func InitLogger(buildEnv string, logFilePath string) zapcore.WriteSyncer {
	var (
		core zapcore.Core

		fileWriter = zapcore.AddSync(&lumberjack.Logger{
			Filename:   logFilePath,
			MaxSize:    10,
			MaxBackups: 1,
			MaxAge:     1,
		})
		consoleWriter = zapcore.AddSync(os.Stdout)
	)

	switch buildEnv {
	case Development:
		customEncoderConfig := NewCustomEncoderConfig()
		jsonEncoder := zapcore.NewJSONEncoder(customEncoderConfig)
		escapeSeqJsonEncoder := &EscapeSeqJSONEncoder{Encoder: jsonEncoder}

		core = zapcore.NewCore(
			escapeSeqJsonEncoder,
			zapcore.NewMultiWriteSyncer(consoleWriter, fileWriter),
			zap.InfoLevel,
		)

	case Production:
		core = zapcore.NewCore(
			zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig()),
			fileWriter,
			zap.InfoLevel,
		)
	}

	logger = zap.New(
		core,
		zap.AddCaller(),
		zap.AddStacktrace(zapcore.ErrorLevel),
	)

	return fileWriter
}

func NewCustomEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    "func",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}
