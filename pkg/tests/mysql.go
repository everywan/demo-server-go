package tests

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlMock struct {
	DB   *sql.DB
	Gdb  *gorm.DB
	Mock sqlmock.Sqlmock
}

func (mock *MysqlMock) Close() {
	_ = mock.DB.Close()
}

func NewMysqlMock(t *testing.T) *MysqlMock {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err, "get sqlmock error")
	gdb, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	assert.NoError(t, err, "create gorm.db error")
	return &MysqlMock{
		DB:   db,
		Gdb:  gdb,
		Mock: mock,
	}
}
