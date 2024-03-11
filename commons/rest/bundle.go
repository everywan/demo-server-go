package rest

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/everywan/demo-server-go/commons/logger"
	"github.com/everywan/demo-server-go/commons/utils"
)

type HTTPBundle struct {
	name       string
	router     http.Handler
	httpServer http.Server

	timeout      time.Duration // 针对 client 端的限制，到时间返回
	writeTimeout time.Duration // 对应 http server 的 writeTimeout
	readTimeout  time.Duration // 对应 http server 的 readTimeout

	port int
}

func (bundle *HTTPBundle) LoadDefault() {
	if bundle.name == "" {
		bundle.name = utils.App()
	}
	if bundle.port == 0 {
		bundle.port = 8080
	}
}

func New(opts ...Option) *HTTPBundle {
	api := &HTTPBundle{}

	for _, opt := range opts {
		opt(api)
	}

	return api
}

func (bundle *HTTPBundle) GetName() string {
	return bundle.name
}

func (bundle *HTTPBundle) Run(ctx context.Context) {
	bundle.LoadDefault()

	bundle.httpServer = http.Server{
		Addr:         ":" + strconv.Itoa(bundle.port),
		Handler:      http.TimeoutHandler(bundle.router, bundle.timeout, ""),
		ReadTimeout:  bundle.readTimeout,
		WriteTimeout: bundle.writeTimeout,
	}
	// 允许取消TimeoutHandler以支持ws等协议
	if bundle.timeout == 0 {
		bundle.httpServer.Handler = bundle.router
	}

	logger.Infof(ctx, "HTTP Server is listening on %s", bundle.httpServer.Addr)
	if err := bundle.httpServer.ListenAndServe(); err != nil {
		if err != http.ErrServerClosed {
			logger.Errorf(ctx, "HTTP Server start error. err:%s", err)
			panic(err)
		}
	}
}

func (bundle *HTTPBundle) Stop(ctx context.Context) {
	logger.Info(ctx, "HTTP Server is shutdown")
	if err := bundle.httpServer.Shutdown(ctx); err != nil {
		logger.Errorf(ctx, "HTTP Server shutdown error. err:%s", err)
	}
}
