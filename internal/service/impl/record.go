package impl

import (
	"context"

	"github.com/everywan/demo-server-go/internal/dao"
	"github.com/everywan/demo-server-go/internal/service"
)

// todo add cache, and curd form cache.
type RecordService struct {
	// cache redis.cache
	do dao.RecordDao
}

func NewRecordService(do dao.RecordDao) *RecordService {
	return &RecordService{
		do: do,
	}
}

var _ service.RecordService = new(RecordService)

func (svc *RecordService) Create(ctx context.Context, req *service.CreateRecordRequest) error {
	_, err := svc.do.Create(ctx, req)
	return err
}

func (svc *RecordService) Query(ctx context.Context, req *service.QueryRecordRequest) (*service.Record, error) {
	return svc.do.Query(ctx, req)
}

func (svc *RecordService) Delete(ctx context.Context, id uint) error {
	return svc.do.Delete(ctx, id)
}
