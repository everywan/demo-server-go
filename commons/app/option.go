package app

import (
	"context"
	"time"

	"github.com/everywan/demo-server-go/commons/logger"
	"github.com/everywan/demo-server-go/commons/utils"
)

type options struct {
	appName      string        // app name
	logger       logger.Logger // 日志
	profilePort  int           // profile port
	stopTimeout  time.Duration // app stop timeout
	includePaths []string      // sentry path

	// lifecycle hooks
	beforeStart []func(ctx context.Context) error
	afterStart  []func(ctx context.Context) error
	beforeStop  []func(ctx context.Context) error
	afterStop   []func(ctx context.Context) error
}

func (opts *options) LoadDefault() {
	if opts.appName == "" {
		opts.appName = utils.App()
	}
	if opts.logger == nil {
		opts.logger = logger.GetLogger()
	}
	if opts.stopTimeout < 0 {
		opts.stopTimeout = DefaultStopTimeout
	}
}

type Option func(*options)

func Name(name string) Option {
	return func(o *options) {
		o.appName = name
	}
}
func WithLogger(logger logger.Logger) Option {
	return func(o *options) {
		o.logger = logger
	}
}

func WithProfiler(profilePort int) Option {
	return func(o *options) {
		o.profilePort = profilePort
	}
}

func StopTimeout(timeout time.Duration) Option {
	if timeout < 0 {
		timeout = 0
	}
	return func(o *options) {
		o.stopTimeout = timeout
	}
}

func SentryIncludePaths(paths ...string) Option {
	return func(o *options) {
		o.includePaths = paths
	}
}

func BeforeStart(fn func(ctx context.Context) error) Option {
	return func(opts *options) {
		opts.beforeStart = append(opts.beforeStart, fn)
	}
}

func AfterStart(fn func(ctx context.Context) error) Option {
	return func(opts *options) {
		opts.afterStart = append(opts.afterStart, fn)
	}
}

func BeforeStop(fn func(ctx context.Context) error) Option {
	return func(opts *options) {
		opts.beforeStop = append(opts.beforeStop, fn)
	}
}

func AfterStop(fn func(ctx context.Context) error) Option {
	return func(opts *options) {
		opts.afterStop = append(opts.afterStop, fn)
	}
}
