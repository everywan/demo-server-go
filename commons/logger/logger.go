package logger

import (
	"context"
)

// 考虑 Zerolog/zap 的实现

type Level int

const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
)

type Logger interface {
	SetLevel(Level)

	Debug(context.Context, ...interface{})
	Debugf(context.Context, string, ...interface{})

	Info(context.Context, ...interface{})
	Infof(context.Context, string, ...interface{})

	Warn(context.Context, ...interface{})
	Warnf(context.Context, string, ...interface{})

	Error(context.Context, ...interface{})
	Errorf(context.Context, string, ...interface{})

	Fatal(context.Context, ...interface{})
	Fatalf(context.Context, string, ...interface{})
}

var std Logger = New()

func GetLogger() Logger {
	return std
}

func SetLogger(logger Logger) {
	std = logger
}

func SetLevel(level Level) {
	std.SetLevel(level)
}
