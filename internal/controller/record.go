package controller

import (
	"net/http"

	"github.com/everywan/demo-server-go/internal/service"
	"github.com/everywan/demo-server-go/pkg/codemsg"
	"github.com/gin-gonic/gin"
)

// todo add log

// RecordController is GRPC controller
type RecordController struct {
	// logger *flog.Logger// todo
	recordSvc service.RecordService
}

func NewRecordController(recordSvc service.RecordService) *RecordController {
	return &RecordController{
		recordSvc: recordSvc,
	}
}

func (ctl *RecordController) Get(c *gin.Context) {
	idReq := &IDRequestUint{}
	if err := c.ShouldBindUri(idReq); err != nil {
		// todo log, stastd
		c.JSON(http.StatusOK, FailResponse(codemsg., err.Error()))
		return
	}
	// 在svc层，error不变，表示发生异常。如果要记录异常状态，则放到resp里做。
	record, err := ctl.recordSvc.Get(c, uint(idReq.ID))
	if err != nil {
		// todo log, stastd
		// demoCtl.logger.WithError(err).
		// 	WithField("func", "doSomething").
		// 	Error("xxx error")
		// c.JSON()
		c.JSON(http.StatusInternalServerError,
			FailResponse(codemsg.InternalErrorCode, err.Error()))
		return
	}
	c.JSON(http.StatusOK, SucResponse(record))
	return
}

func (ctl *RecordController) Create(c *gin.Context) {
	idReq := &IDRequestUint{}
	if err := c.ShouldBindUri(idReq); err != nil {
		// todo log, stastd
		c.JSON(http.StatusOK, FailResponse(codemsg.BadRequestErrorCode, err.Error()))
		return
	}
	record, err := ctl.recordSvc.Get(c, uint(idReq.ID))
	if err != nil {
		// todo log, stastd
		// demoCtl.logger.WithError(err).
		// 	WithField("func", "doSomething").
		// 	Error("xxx error")
		// c.JSON()
		c.JSON(http.StatusInternalServerError,
			FailResponse(codemsg.InternalErrorCode, err.Error()))
		return
	}
	c.JSON(http.StatusOK, SucResponse(record))
	return
}
