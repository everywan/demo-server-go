package bootstrap

import (
	"github.com/everywan/demo-server-go/internal/controller"
)

type controllerComponent struct {
	recordCtl     *controller.RecordController
	recordGrpcCtl *controller.RecordGrpcController
}

func (b *Bootstrap) GetRecordController() *controller.RecordController {
	if b.recordCtl == nil {
		b.recordCtl = controller.NewRecordController(b.GetRecordService())
	}
	return b.recordCtl
}

func (b *Bootstrap) GetRecordGrpcController() *controller.RecordGrpcController {
	if b.recordGrpcCtl == nil {
		b.recordGrpcCtl = controller.NewRecordGrpcController(b.GetRecordService())
	}
	return b.recordGrpcCtl
}
