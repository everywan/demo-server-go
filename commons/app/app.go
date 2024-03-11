package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/everywan/demo-server-go/commons/logger"
)

type (
	Application interface {
		Name() string
		Run(context.Context)
		AddBundle(...IBundle)
	}
	BaseApplication struct {
		name   string
		config *appConfig
		logger logger.Logger

		// bundles
		bundles []IBundle

		// app lifecycle hooks
		beforeStart []func(ctx context.Context) error
		afterStart  []func(ctx context.Context) error
		beforeStop  []func(ctx context.Context) error
		afterStop   []func(ctx context.Context) error
	}
)

func New(_opts ...Option) *BaseApplication {
	opts := &options{}
	for _, opt := range _opts {
		opt(opts)
	}
	opts.LoadDefault()

	return &BaseApplication{
		name: opts.appName,
		config: &appConfig{
			stopTimeout: opts.stopTimeout,
			profilePort: opts.profilePort,
		},
		logger:      opts.logger,
		beforeStart: opts.beforeStart,
		afterStart:  opts.afterStart,
		beforeStop:  opts.beforeStop,
		afterStop:   opts.afterStop,
	}

}

var _ Application = new(BaseApplication)

func (app *BaseApplication) Name() string {
	return app.name
}

func (app *BaseApplication) Run(ctx context.Context) {
	app.logger.Infof(ctx, "Run application [%s]", app.name)

	// 设置 sentry

	if app.config.profilePort > 0 {
		go func() {
			err := http.ListenAndServe(fmt.Sprintf(":%d", app.config.profilePort), nil)
			if err != nil {
				app.logger.Error(ctx, "Start pprof error: ", err)
			}
		}()
	}

	app.runBeforeStart(ctx)
	bundleFinishCtx := app.runAllBundles(ctx)
	app.runAfterStart(ctx)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	select {
	case <-bundleFinishCtx.Done():
		app.logger.Info(ctx, "All bundle finished.")
	case <-quit:
		app.logger.Info(ctx, "Shutdown signal received.")
	}

	app.runBeforeStop(ctx)
	bundleStopCtx := app.stopAllBundles(ctx)
	shutdownTimeout := time.After(app.config.stopTimeout)
	select {
	case <-bundleStopCtx.Done():
		app.logger.Info(ctx, "Application stopped")
	case <-shutdownTimeout:
		app.logger.Info(ctx, "Shutdown timeout, force stop application")
	}
	app.runAfterStop(ctx)

	// 清除其他资源(如有)

	app.logger.Info(ctx, "Bye!")
}

func (app *BaseApplication) runBeforeStart(ctx context.Context) {
	for _, fn := range app.beforeStart {
		if err := fn(ctx); err != nil {
			app.logger.Fatalf(ctx, "before bundles start fn fatal:%v", err)
		}
	}
}

func (app *BaseApplication) runAfterStart(ctx context.Context) {
	for _, fn := range app.afterStart {
		if err := fn(ctx); err != nil {
			app.logger.Errorf(ctx, "after bundles start fn error:%v", err)
		}
	}
}

func (app *BaseApplication) runBeforeStop(ctx context.Context) {
	for _, fn := range app.beforeStop {
		if err := fn(ctx); err != nil {
			app.logger.Errorf(ctx, "before bundles stop fn error:%v", err)
		}
	}
}

func (app *BaseApplication) runAfterStop(ctx context.Context) {
	for _, fn := range app.afterStop {
		if err := fn(ctx); err != nil {
			app.logger.Errorf(ctx, "after bundles stop fn error:%v", err)
		}
	}
}

func (app *BaseApplication) runAllBundles(ctx context.Context) context.Context {
	finishCtx, cancel := context.WithCancel(ctx)
	wg := sync.WaitGroup{}

	for _, bundle := range app.bundles {
		bundle := bundle
		wg.Add(1)
		// 由 bundle 自己决定是否捕获异常, app 不做处理.
		go func() {
			defer wg.Done()
			bundle.Run(finishCtx)
		}()
		app.logger.Infof(ctx, "bundle %s start suceess.", bundle.GetName())
	}

	go func() {
		wg.Wait()
		cancel()
	}()

	return finishCtx
}

func (app *BaseApplication) stopAllBundles(ctx context.Context) context.Context {
	finishCtx, cancel := context.WithCancel(ctx)
	wg := sync.WaitGroup{}

	for _, bundle := range app.bundles {
		bundle := bundle
		wg.Add(1)
		// 由 bundle 自己决定是否捕获异常, app 不做处理.
		go func() {
			defer wg.Done()
			bundle.Stop(finishCtx)
		}()
		app.logger.Infof(ctx, "bundle %s stop suceess.", bundle.GetName())
	}

	go func() {
		wg.Wait()
		cancel()
		app.logger.Info(ctx, "All bundle stopped")
	}()

	return finishCtx
}

func (app *BaseApplication) AddBundle(bundles ...IBundle) {
	app.bundles = append(app.bundles, bundles...)
}
