package controller

import (
	"context"
	"net/http"

	"github.com/everywan/demo-server-go/commons/errors"
	"github.com/everywan/demo-server-go/commons/logger"
	"github.com/everywan/demo-server-go/commons/rest"
	"github.com/everywan/demo-server-go/internal/service"
	"github.com/everywan/demo-server-go/pkg/codemsg"
	"github.com/gin-gonic/gin"
)

// RecordController is GRPC controller
type RecordController struct {
	recordSvc service.RecordService
}

func NewRecordController(recordSvc service.RecordService) *RecordController {
	return &RecordController{
		recordSvc: recordSvc,
	}
}

func (ctl *RecordController) Query(c *gin.Context) {
	idReq := &rest.IDRequestUint{}
	if err := c.ShouldBindUri(idReq); err != nil {
		// todo stastd
		logger.Error(context.Background(),
			"RecordController.Query ShouldBindUri error. err:%s", err)
		c.JSON(http.StatusOK, rest.FailResponse(http.StatusBadRequest, err.Error()))
		return
	}
	// 在svc层，error不变，表示发生异常。如果要记录异常状态，则放到resp里做。
	req := &service.QueryRecordRequest{
		ID: idReq.ID,
	}
	record, err := ctl.recordSvc.Query(c, req)
	if err != nil {
		errCode, ok := err.(*errors.ErrorCode)
		if !ok {
			// todo statsd
			logger.Error(context.Background(), "RecordController.Query unknown error. err:%s", err)
			c.JSON(http.StatusInternalServerError,
				rest.FailResponse(http.StatusInternalServerError, err.Error()))
		}
		switch errCode.Code {
		case codemsg.SelfDeinfeStatsu1:
			// 假如这种场景下返回正常.
			c.JSON(http.StatusOK, rest.SucResponse(nil))
			return
		default:
			// todo statsd
			logger.Error(context.Background(),
				"RecordController.Query unknown error. err:%s", err)
			c.JSON(http.StatusInternalServerError,
				rest.FailResponse(http.StatusInternalServerError, err.Error()))
			return
		}
	}
	c.JSON(http.StatusOK, rest.SucResponse(record))
}
