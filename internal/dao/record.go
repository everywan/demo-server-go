package dao

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

//go:generate mockgen -destination mock/record.go -source=record.go
type RecordDao interface {
	Create(context.Context, *CreateRecordRequest) (id uint, err error)
	Update(context.Context, *UpdateRecordRequest) error
	UpdateStatus(context.Context, *UpdateRecordStatusRequest) error
	Get(ctx context.Context, id uint) (*Record, error)
	Query(context.Context, *QueryRecordRequest) (*Record, error)
	List(context.Context, *ListRecordRequest) (*ListRecordResponse, error)
	Delete(ctx context.Context, id uint) error
}

type RecordStatus int8

const (
	RecordStatusInit RecordStatus = iota
	RecordStatusCase1
	RecordStatusCase2
)

type Record struct {
	gorm.Model
	Name   string       `gorm:"column:name" json:"name"`
	Status RecordStatus `grom:"column:status" json:"status"`

	CreatedBy uint64 `gorm:"column:created_by" json:"created_by"`
	UpdatedBy uint64 `gorm:"column:updated_by" json:"updated_by"`
}

func (Record) TableName() string {
	return "records"
}

type CreateRecordRequest struct {
	Name      string       `json:"name"`
	Status    RecordStatus `json:"status"`
	CreatedBy uint64       `json:"created_by"`
}

type UpdateRecordRequest struct {
	ID        uint    `json:"id"`
	Name      *string `json:"name"`
	UpdatedBy uint64  `json:"updated_by"`
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

type UpdateRecordStatusRequest struct {
	ID        uint         `json:"id"`
	Status    RecordStatus `json:"status"`
	UpdatedBy uint64       `json:"updated_by"`
}

func (req *UpdateRecordStatusRequest) Validate() error {
	if req.ID == 0 {
		return errors.New("must have id")
	}
	return nil
}

type QueryRecordRequest struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
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

type (
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
