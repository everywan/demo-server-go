package dao

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

//go:generate mockgen -destination mock/record.go -source=record.go
type RecordDao interface {
	Create(context.Context, *CreateRecordRequest) error
	// 只更新非零值数据. status 更新操作频繁且功能单一, 一般单独开接口实现.
	Update(context.Context, *UpdateRecordRequest) error
	// 无论 status 是否有值, 都进行更新
	UpdateStatus(context.Context, *UpdateRecordStatusRequest) error
	Get(ctx context.Context, id uint) (*Record, error)
	// 一般是一条记录有多个业务唯一ID, 用 query 查询, 返回符合的第一个.
	Query(context.Context, *QueryRecordRequest) (*Record, error)
	List(context.Context, *ListRecordRequest) (*ListRecordResponse, error)
	Delete(ctx context.Context, id uint) error
}

const (
	RecordStatusInit = iota
	RecordStatusCase1
	RecordStatusCase2
)

type (
	RecordStatus int8
	Record       struct {
		ID     uint64       `gorm:"column:id" json:"id"`
		Name   string       `gorm:"column:name" json:"name"`
		Status RecordStatus `grom:"column:status" json:"status"`

		CreatedAt int64  `gorm:"column:created_at" json:"created_at"`
		UpdatedAt int64  `gorm:"column:updated_at" json:"updated_at"`
		DeletedAt *int64 `gorm:"column:deleted_at" json:"deleted_at"`
		CreatedBy uint64 `gorm:"column:created_by" json:"created_by"`
		UpdatedBy uint64 `gorm:"column:updated_by" json:"updated_by"`
	}
	CreateRecordRequest struct {
		Name      string       `json:"name"`
		Status    RecordStatus `json:"status"`
		CreatedBy uint64       `json:"created_by"`
	}
	UpdateRecordRequest struct {
		ID        uint64  `json:"id"`
		Name      *string `json:"name"` // 假设 Name 零值有意义, 则将字段设置为指针以启用更新
		UpdatedBy uint64  `json:"updated_by"`
	}
	UpdateRecordStatusRequest struct {
		ID        uint64       `json:"id"`
		Status    RecordStatus `json:"status"`
		UpdatedBy uint64       `json:"updated_by"`
	}
	QueryRecordRequest struct {
		ID   uint64 `json:"id"`
		Name string `json:"name"`
	}
	ListRecordRequest struct {
		NameLike string        `json:"name_like"`
		Status   *RecordStatus `json:"status"`
		LastID   uint64        `json:"last_id"` // 解决深翻页
		Limit    int           `json:"limit"`
		Offset   int           `json:"offset"`
		Order    string        `json:"order"`
		OrderBy  string        `json:"order_by"`
	}
	ListRecordResponse struct {
		Total   int       `json:"total"`
		Records []*Record `json:"records"`
	}
)

func (Record) TableName() string {
	return "records"
}

func (req *UpdateRecordRequest) Validate() error {
	if req.ID == 0 {
		return errors.New("must have id")
	}
	if req.Name == nil {
		return errors.New("no data need update")
	}
	return nil
}

func (req *UpdateRecordStatusRequest) Validate() error {
	if req.ID == 0 {
		return errors.New("must have id")
	}
	return nil
}

func (req *QueryRecordRequest) Validate() error {
	if req.ID == 0 && req.Name == "" {
		return errors.New("must have a condition")
	}
	return nil
}

func (req *QueryRecordRequest) BuildQuery(db *gorm.DB) *gorm.DB {
	if req.ID != 0 {
		db = db.Where("id = ?", req.ID)
	}
	if req.Name != "" {
		db = db.Where("name = ?", req.Name)
	}
	db = db.Limit(1)
	return db
}

func (req *ListRecordRequest) LoadDefault() {
	if req.Limit < 1 {
		req.Limit = 10
	}
	if req.Offset < 1 || req.LastID != 0 {
		req.Offset = 0
	}
	if req.Order == "" {
		req.Order = "desc"
	}
	if req.OrderBy == "" {
		req.Order = "updated_at"
	}
}

func (req *ListRecordRequest) BuildQuery(db *gorm.DB) *gorm.DB {
	db = req.buildQuery(db)
	if req.LastID != 0 {
		db.Where("id > ?", req.LastID)
		// lastid 有值时, offset 设置为 0
		req.Offset = 0
	}
	db = db.Order(req.OrderBy + " " + req.Order).Offset(req.Offset).Limit(req.Limit)

	return db
}

func (req *ListRecordRequest) BuildCountQuery(db *gorm.DB) *gorm.DB {
	db = req.buildQuery(db)
	return db
}

func (req *ListRecordRequest) buildQuery(db *gorm.DB) *gorm.DB {
	if req.NameLike != "" {
		db = db.Where("name like %?%", req.NameLike)
	}
	if req.Status != nil {
		db = db.Where("status = ?", *req.Status)
	}
	return db
}
