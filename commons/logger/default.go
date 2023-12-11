package logger

import (
	"context"
	"fmt"
)

type defaultLogger struct {
	level Level
}

var _ Logger = new(defaultLogger)

func New() *defaultLogger {
	return &defaultLogger{}
}

func (l *defaultLogger) SetLevel(level Level) {
	l.level = level
}

// todo logger
func (l *defaultLogger) Debug(ctx context.Context, args ...interface{}) {
	panic("")
}

func (l *defaultLogger) Debugf(ctx context.Context, format string, args ...interface{}) {
	l.Debug(ctx, fmt.Sprintf(format, args...))
}

func (l *defaultLogger) Info(ctx context.Context, args ...interface{}) {
	panic("")
}

func (l *defaultLogger) Infof(ctx context.Context, format string, args ...interface{}) {
	l.Info(ctx, fmt.Sprintf(format, args...))
}

func (l *defaultLogger) Warn(ctx context.Context, args ...interface{}) {
	panic("")
}

func (l *defaultLogger) Warnf(ctx context.Context, format string, args ...interface{}) {
	l.Warn(ctx, fmt.Sprintf(format, args...))
}

func (l *defaultLogger) Error(ctx context.Context, args ...interface{}) {
	panic("")
}

func (l *defaultLogger) Errorf(ctx context.Context, format string, args ...interface{}) {
	l.Error(ctx, fmt.Sprintf(format, args...))
}

func (l *defaultLogger) Fatal(ctx context.Context, args ...interface{}) {
	panic("")
}

func (l *defaultLogger) Fatalf(ctx context.Context, format string, args ...interface{}) {
	l.Fatal(ctx, fmt.Sprintf(format, args...))
}
