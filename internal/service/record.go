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
	Update(context.Context, *UpdateRecordRequest) error
	UpdateStatus(context.Context, *UpdateRecordStatusRequest) error
	Get(ctx context.Context, id uint) (*Record, error)
	Query(context.Context, *QueryRecordRequest) (*Record, error)
	List(context.Context, *ListRecordRequest) (*ListRecordResponse, error)
	Delete(ctx context.Context, id uint) error
}

// service 层目前没有特殊定制, 所以直接使用 dao 层的结构.
type (
	Record                    = dao.Record
	CreateRecordRequest       = dao.CreateRecordRequest
	UpdateRecordRequest       = dao.UpdateRecordRequest
	UpdateRecordStatusRequest = dao.UpdateRecordStatusRequest
	QueryRecordRequest        = dao.QueryRecordRequest
	ListRecordRequest         = dao.ListRecordRequest
	ListRecordResponse        = dao.ListRecordResponse
)
