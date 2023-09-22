package bootstrap

import (
	"github.com/everywan/demo-server-go/internal/service"
	"github.com/everywan/demo-server-go/internal/service/impl"
)

type serviceComponent struct {
	recordSvc service.RecordService
}

func (b *Bootstrap) GetRecordService() service.RecordService {
	if b.recordSvc == nil {
		b.recordSvc = impl.NewRecordService(b.GetRecordDao())
	}
	return b.recordSvc
}
