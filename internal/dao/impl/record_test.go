package impl

import (
	"context"
	"database/sql/driver"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/everywan/demo-server-go/commons/tests"
	"github.com/everywan/demo-server-go/internal/dao"
	"github.com/stretchr/testify/assert"
)

type RecordTestHelper struct {
	sqlmock   sqlmock.Sqlmock
	recordDao *RecordDao

	teardown []func()
}

func NewRecordTestHelper(t *testing.T) *RecordTestHelper {
	mysqlMock := tests.NewMysqlMock(t)
	return &RecordTestHelper{
		sqlmock:   mysqlMock.Mock,
		recordDao: NewRecordDao(mysqlMock.Gdb),
		teardown: []func(){
			mysqlMock.Close,
		},
	}
}

func TestRecordCreate(t *testing.T) {
	suite := NewRecordTestHelper(t)
	ctx := context.Background()
	req := &dao.CreateRecordRequest{
		Name:      "test_name_1",
		Status:    1,
		CreatedBy: 100,
	}
	record := &dao.Record{}

	// validate create sql
	{
		sql := fmt.Sprintf("^INSERT INTO `%s` \\(`name`,`status`,`created_at`,"+
			"`updated_at`,`deleted_at`,`created_by`,`updated_by`\\) VALUES \\(\\?,"+
			"\\?,\\?,\\?,\\?,\\?,\\?\\)", record.TableName())
		suite.sqlmock.ExpectBegin()
		suite.sqlmock.ExpectExec(sql).
			WithArgs(req.Name, req.Status, sqlmock.AnyArg(), sqlmock.AnyArg(), nil,
				req.CreatedBy, req.CreatedBy).
			WillReturnResult(sqlmock.NewResult(1, 1))
		suite.sqlmock.ExpectCommit()
	}

	_, err := suite.recordDao.Create(ctx, req)
	assert.NoError(t, err, "dao.Record.create happend error")
}

func TestRecordUpdate(t *testing.T) {
	suite := NewRecordTestHelper(t)

	ctx := context.Background()
	record := &dao.Record{}
	updateRequests := []*dao.UpdateRecordRequest{
		// validate, must have name
		{
			ID:        100,
			UpdatedBy: 101,
		},
		// update
		{
			ID:        1001,
			Name:      func() *string { name := "test_update1"; return &name }(),
			UpdatedBy: 102,
		},
	}
	testcases := []struct {
		name       string
		req        *dao.UpdateRecordRequest
		expectArgs []driver.Value
		expectSql  string
		expectErr  bool
	}{
		{
			name:      "case1:validate_name",
			req:       updateRequests[0],
			expectErr: true,
		},
		{
			name: "case2:update_some_field",
			req:  updateRequests[1],
			expectArgs: []driver.Value{updateRequests[1].Name, sqlmock.AnyArg(),
				updateRequests[1].UpdatedBy, updateRequests[1].ID},
			expectSql: fmt.Sprintf("UPDATE `%s` SET `name`=\\?,`updated_at`=\\?,"+
				"`updated_by`=\\? WHERE `%s`.`deleted_at` IS NULL AND `id` = \\?",
				record.TableName(), record.TableName()),
		},
	}
	for _, tcase := range testcases {
		if tcase.expectSql != "" {
			suite.sqlmock.ExpectBegin()
			suite.sqlmock.ExpectExec(tcase.expectSql).
				WithArgs(tcase.expectArgs...).
				WillReturnResult(sqlmock.NewResult(int64(tcase.req.ID), 1))
			suite.sqlmock.ExpectCommit()
		}

		err := suite.recordDao.Update(ctx, tcase.req)
		if tcase.expectErr {
			assert.Error(t, err, "testcase [%s]: dao.record.update should error",
				tcase.name)
			continue
		}
		assert.NoError(t, err, "testcase [%s]: dao.record.update should not error",
			tcase.name)
	}
}

func TestRecordUpdateStatus(t *testing.T) {
	suite := NewRecordTestHelper(t)

	ctx := context.Background()
	record := &dao.Record{}
	testcases := []struct {
		name       string
		id         uint
		status     dao.RecordStatus
		updateBy   uint64
		expectArgs []driver.Value
		expectSql  string
		expectErr  bool
	}{
		{
			name:      "case1:validate_update_by_exist",
			id:        0,
			expectErr: true,
		},
		{
			name:       "case2:update_to_1",
			id:         1,
			status:     1,
			updateBy:   101,
			expectArgs: []driver.Value{1, sqlmock.AnyArg(), 101, 1},
			expectSql: fmt.Sprintf("UPDATE `%s` SET `status`=\\?,`updated_at`=\\?,"+
				"`updated_by`=\\? WHERE `%s`.`deleted_at` IS NULL AND `id` = \\?",
				record.TableName(), record.TableName()),
		},
		{
			name:       "case3:update_to_0",
			id:         1,
			status:     0,
			updateBy:   102,
			expectArgs: []driver.Value{0, sqlmock.AnyArg(), 102, 1},
			expectSql: fmt.Sprintf("UPDATE `%s` SET `status`=\\?,`updated_at`=\\?,"+
				"`updated_by`=\\? WHERE `%s`.`deleted_at` IS NULL AND `id` = \\?",
				record.TableName(), record.TableName()),
		},
	}
	for _, tcase := range testcases {
		if tcase.expectSql != "" {
			suite.sqlmock.ExpectBegin()
			suite.sqlmock.ExpectExec(tcase.expectSql).
				WithArgs(tcase.expectArgs...).
				WillReturnResult(sqlmock.NewResult(int64(tcase.id), 1))
			suite.sqlmock.ExpectCommit()
		}

		err := suite.recordDao.UpdateStatus(ctx, &dao.UpdateRecordStatusRequest{
			ID:        tcase.id,
			Status:    tcase.status,
			UpdatedBy: tcase.updateBy,
		})
		if tcase.expectErr {
			assert.Error(t, err, "testcase [%s]: dao.record.update_status should error",
				tcase.name)
			continue
		}
		assert.NoError(t, err, "testcase [%s]: dao.record.update_status should not error",
			tcase.name)
	}
}

func TestRecordDelete(t *testing.T) {
	suite := NewRecordTestHelper(t)

	ctx := context.Background()
	record := &dao.Record{}
	id := uint(1003)
	// validate create sql
	{
		sql := fmt.Sprintf("UPDATE `%s` SET `deleted_at`=\\? WHERE `%s`.`id` = \\? "+
			"AND `%s`.`deleted_at` IS NULL", record.TableName(), record.TableName(),
			record.TableName())
		suite.sqlmock.ExpectBegin()
		suite.sqlmock.ExpectExec(sql).
			WithArgs(sqlmock.AnyArg(), id).
			WillReturnResult(sqlmock.NewResult(int64(id), 1))
		suite.sqlmock.ExpectCommit()
	}

	err := suite.recordDao.Delete(ctx, id)
	assert.NoError(t, err, "dao.Record.delete happend error")
}

func TestRecordGet(t *testing.T) {
	suite := NewRecordTestHelper(t)

	ctx := context.Background()
	record := &dao.Record{}
	testcases := []struct {
		name       string
		id         uint
		expectName string
	}{
		{"case1:get", 1003, "test_name"},
	}
	for _, tcase := range testcases {
		{
			sql := fmt.Sprintf("^SELECT \\* FROM `%s` WHERE `%s`.`id` = \\? AND `%s`.`deleted_at` IS NULL ORDER BY `%s`.`id` LIMIT 1",
				record.TableName(), record.TableName(), record.TableName(), record.TableName())
			suite.sqlmock.ExpectQuery(sql).
				WithArgs(tcase.id).
				WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).
					FromCSVString(fmt.Sprintf("%d,%s", tcase.id, tcase.expectName)))
		}
		record, err := suite.recordDao.Get(ctx, tcase.id)
		if assert.NoError(t, err, "dao.Record.get happend error") {
			assert.Equal(t, tcase.id, record.ID, "result.id(%d)!=expect.id(%d)",
				record.ID, tcase.id)
			assert.Equal(t, tcase.expectName, record.Name, "result.name(%s)!=expect.name(%s)",
				record.Name, tcase.expectName)
		}
	}
}
