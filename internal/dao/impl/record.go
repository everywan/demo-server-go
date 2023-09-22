package impl

import (
	"context"
	"time"

	"github.com/everywan/demo-server-go/internal/dao"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type RecordDao struct {
	db *gorm.DB
}

func NewRecordDao(db *gorm.DB) *RecordDao {
	return &RecordDao{
		db: db,
	}
}

var _ dao.RecordDao = new(RecordDao)

// Create 创建数据库记录示例
func (do *RecordDao) Create(ctx context.Context, req *dao.CreateRecordRequest) error {
	record := &dao.Record{
		Name:      req.Name,
		Status:    req.Status,
		CreatedAt: time.Now().Unix(),
		CreatedBy: req.CreatedBy,
		UpdatedBy: req.CreatedBy,
	}
	err := do.db.Model(record).Create(record).Error
	if err != nil {
		return errors.Wrapf(err, "create record error. request:%+v", req)
	}
	return nil
}

// 只更新非零值字段
func (do *RecordDao) Update(ctx context.Context, req *dao.UpdateRecordRequest) error {
	if err := req.Validate(); err != nil {
		return err
	}
	record := &dao.Record{
		ID:        req.ID,
		Name:      *req.Name,
		UpdatedAt: time.Now().Unix(),
		UpdatedBy: req.UpdatedBy,
	}
	err := do.db.Model(record).Updates(record).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return err
		}
		return errors.Wrapf(err, "update record error. request:%+v", req)
	}
	return nil
}

func (do *RecordDao) UpdateStatus(ctx context.Context, req *dao.UpdateRecordStatusRequest) error {
	if err := req.Validate(); err != nil {
		return err
	}
	record := &dao.Record{
		ID:        req.ID,
		Status:    req.Status,
		UpdatedAt: time.Now().Unix(),
		UpdatedBy: req.UpdatedBy,
	}
	err := do.db.Model(record).
		Select("status", "update_at", "updated_by").
		Updates(record).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return err
		}
		return errors.Wrapf(err, "update record status error. request:%+v", req)
	}
	return nil
}

func (do *RecordDao) Get(ctx context.Context, id uint) (*dao.Record, error) {
	if id < 1 {
		return &dao.Record{}, errors.New("record id should gt 0")
	}
	record := &dao.Record{}
	err := do.db.Model(&dao.Record{}).First(record, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return &dao.Record{}, err
		}
		return &dao.Record{}, errors.Wrapf(err, "get record error. id:%d", id)
	}
	return record, nil
}

func (do *RecordDao) Query(ctx context.Context, req *dao.QueryRecordRequest) (*dao.Record, error) {
	if err := req.Validate(); err != nil {
		return &dao.Record{}, err
	}
	record := &dao.Record{}
	err := req.BuildQuery(do.db.Model(record)).First(record).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return &dao.Record{}, err
		}
		return &dao.Record{}, errors.Wrapf(err, "query record error. request:%+v", req)
	}
	return record, nil
}

func (do *RecordDao) List(ctx context.Context, req *dao.ListRecordRequest) (*dao.ListRecordResponse, error) {
	req.LoadDefault()

	// 查询 count
	var cnt int64
	err := req.BuildCountQuery(do.db.Model(&dao.Record{})).
		Count(&cnt).Error
	if err != nil {
		err = errors.Wrapf(err, "list record error on count total. requst:%+v", req)
		return &dao.ListRecordResponse{}, err
	}

	if cnt == 0 {
		return &dao.ListRecordResponse{
			Total:   int(cnt),
			Records: []*dao.Record{},
		}, nil
	}

	// 查询记录
	records := []*dao.Record{}
	err = req.BuildQuery(do.db.Model(&dao.Record{})).
		Find(&records).
		Error
	if err != nil {
		err = errors.Wrapf(err, "list record error. request:%+v", req)
		return &dao.ListRecordResponse{}, err
	}
	return &dao.ListRecordResponse{
		Total:   int(cnt),
		Records: records,
	}, nil
}

func (do *RecordDao) Delete(ctx context.Context, id uint) error {
	if id < 1 {
		return errors.New("record id should gt 0")
	}
	err := do.db.Delete(&dao.Record{ID: uint64(id)}).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		return errors.Wrapf(err, "delete record error. id:%d", id)
	}
	return nil
}
