package bootstrap

import (
	"log"

	"github.com/everywan/demo-server-go/commons/logger"
)

type Bootstrap struct {
	logger    logger.Logger
	teardowns []func()

	configComponent
	daoComponent
	serviceComponent
	controllerComponent
}

func NewBootstrap() *Bootstrap {
	return &Bootstrap{
		teardowns: []func(){},
	}
}

func (b *Bootstrap) Teardown() {
	for _, fn := range b.teardowns {
		fn()
	}
}

func (b *Bootstrap) GetLogger() logger.Logger {
	if b.logger == nil {
		b.logger = logger.New()
	}
	return b.logger
}

func (b *Bootstrap) AddTeardown(teardown func()) {
	b.teardowns = append(b.teardowns, teardown)
}

func handleInitError(module string, err error) {
	if err == nil {
		return
	}
	log.Fatalf("init %s failed, err: %s", module, err)
}
