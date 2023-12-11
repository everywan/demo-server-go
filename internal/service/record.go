package service

import (
	"context"

	"github.com/everywan/demo-server-go/internal/dao"
)

// RecordService 假设业务领域是 Record. 注意, service 是一个接口, 我
// 们是面向接口编程的.
//
//go:generate mockgen -destination mock/record.go -source=record.go
type RecordService interface {
	Create(context.Context, *CreateRecordRequest) error
	Query(context.Context, *QueryRecordRequest) (*Record, error)
	Delete(ctx context.Context, id uint) error
}

// service 层目前没有特殊定制, 所以直接使用 dao 层的结构.
type (
	Record              = dao.Record
	CreateRecordRequest = dao.CreateRecordRequest
	QueryRecordRequest  = dao.QueryRecordRequest
)
