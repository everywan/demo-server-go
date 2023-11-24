package bootstrap

import (
	"github.com/everywan/demo-server-go/internal/dao"
	"github.com/everywan/demo-server-go/internal/dao/impl"
	"gorm.io/gorm"
)

type daoComponent struct {
	recordDB  *gorm.DB
	recordDao dao.RecordDao
}

func (b *Bootstrap) GetRecordDB() *gorm.DB {
	// if b.recordDB == nil {
	// 	// todo db 封装
	// 	// db, err := database.NewDatabase(b.cfg.DB)
	// 	// handleInitError("connect database", err)
	// 	// b.addTeardown(func() { db.Close() })
	// }
	return b.recordDB
}

func (b *Bootstrap) GetRecordDao() dao.RecordDao {
	if b.recordDao == nil {
		b.recordDao = impl.NewRecordDao(b.GetRecordDB())
	}
	return b.recordDao
}
