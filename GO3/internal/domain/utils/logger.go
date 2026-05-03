package utils

import (
	"container/list"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const maxMessage = 5

type Logger interface {
	Debug(msg string, fields ...zap.Field)
	Info(msg string, fields ...zap.Field)
	Warn(msg string, fields ...zap.Field)
	Error(msg string, fields ...zap.Field)
	Fatal(msg string, fields ...zap.Field)
	GameMessage(msg string)
	GetMessages() []string
	Sync() error
}

type gameLogger struct {
	*zap.Logger
	cache *list.List
}

func (gl *gameLogger) GameMessage(msg string) {
	gl.cache.PushFront(msg)
	for gl.cache.Len() > maxMessage {
		gl.cache.Remove(gl.cache.Back())
	}
}

func (gl *gameLogger) GetMessages() []string {
	res := make([]string, 0, gl.cache.Len())
	for e := gl.cache.Front(); e != nil; e = e.Next() {
		str := e.Value.(string)
		res = append(res, str)
	}
	return res
}

func NewFileLogger(filePath string) (Logger, error) {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}

	config := zap.Config{
		OutputPaths:      []string{file.Name()},
		ErrorOutputPaths: []string{"stderr"},
		Encoding:         "json",
		Level:            zap.NewAtomicLevelAt(zap.WarnLevel),
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:     "message",
			TimeKey:        "timestamp",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			FunctionKey:    zapcore.OmitKey,
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.CapitalLevelEncoder,
			EncodeTime:     zapcore.EpochTimeEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
	}

	logger, err := config.Build()
	if err != nil {
		return nil, err
	}

	return &gameLogger{Logger: logger, cache: list.New()}, nil
}

type NoopLogger struct{}

func (*NoopLogger) Debug(msg string, fields ...zap.Field) {}
func (*NoopLogger) Info(msg string, fields ...zap.Field)  {}
func (*NoopLogger) Warn(msg string, fields ...zap.Field)  {}
func (*NoopLogger) Error(msg string, fields ...zap.Field) {}

func (*NoopLogger) GameMessage(msg string)                {}
func (*NoopLogger) GetMessages() []string                 { return nil }
func (*NoopLogger) Fatal(msg string, fields ...zap.Field) { panic(msg) }
func (*NoopLogger) Sync() error                           { return nil }

func NewNoopLogger() *NoopLogger {
	return &NoopLogger{}
}
