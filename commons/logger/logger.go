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

func Debug(ctx context.Context, args ...interface{}) {
	std.Debug(ctx, args...)
}

func Debugf(ctx context.Context, format string, args ...interface{}) {
	std.Debugf(ctx, format, args...)
}

func Info(ctx context.Context, args ...interface{}) {
	std.Info(ctx, args...)
}

func Infof(ctx context.Context, format string, args ...interface{}) {
	std.Infof(ctx, format, args...)
}

func Warn(ctx context.Context, args ...interface{}) {
	std.Warn(ctx, args...)
}

func Warnf(ctx context.Context, format string, args ...interface{}) {
	std.Warnf(ctx, format, args...)
}

func Error(ctx context.Context, args ...interface{}) {
	std.Error(ctx, args...)
}

func Errorf(ctx context.Context, format string, args ...interface{}) {
	std.Errorf(ctx, format, args...)
}

func Fatal(ctx context.Context, args ...interface{}) {
	std.Fatal(ctx, args...)
}

func Fatalf(ctx context.Context, format string, args ...interface{}) {
	std.Fatalf(ctx, format, args...)
}
