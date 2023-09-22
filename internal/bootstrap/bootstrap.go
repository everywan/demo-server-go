package bootstrap

import (
	"log"
)

type Bootstrap struct {
	// logger   *flog.Logger
	teardown func()

	configComponent
	daoComponent
	serviceComponent
	controllerComponent
}

func NewBootstrap() *Bootstrap {
	return &Bootstrap{
		teardown: func() {},
	}
}

func (b *Bootstrap) Teardown() {
	b.teardown()
}

// func (b *Bootstrap) GetLogger() *flog.Logger {
// 	if b.logger == nil {
// 		b.logger = flog.NewLogger(b.cfg.Logging, os.Stdout)
// 	}
// 	return b.logger
// }

func (b *Bootstrap) addTeardown(newTeardown func()) {
	teardown := b.teardown
	b.teardown = func() {
		teardown()
		newTeardown()
	}
}

func handleInitError(module string, err error) {
	if err == nil {
		return
	}
	log.Fatalf("init %s failed, err: %s", module, err)
}
