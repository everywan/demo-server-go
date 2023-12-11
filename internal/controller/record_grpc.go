package controller

import (
	"context"

	"github.com/everywan/demo-server-go/commons/errors"
	"github.com/everywan/demo-server-go/commons/logger"
	record_pb "github.com/everywan/demo-server-go/idl/proto/record"
	"github.com/everywan/demo-server-go/internal/service"
	"github.com/everywan/demo-server-go/pkg/codemsg"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RecordGrpcController struct {
	recordSvc service.RecordService
}

var _ record_pb.RecordServiceServer = new(RecordGrpcController)

func NewRecordGrpcController(recordSvc service.RecordService) *RecordGrpcController {
	return &RecordGrpcController{
		recordSvc: recordSvc,
	}
}

func (cl *RecordGrpcController) Get(ctx context.Context, req *record_pb.IDRequest,
) (resp *record_pb.Record, err error) {
	if err := ValidateIDRequest(req); err != nil {
		logger.Error(context.Background(),
			"RecordController.Query ShouldBindUri error. err:%s", err)
		return &record_pb.Record{}, status.Errorf(codes.InvalidArgument, err.Error())
	}
	svcReq := &service.QueryRecordRequest{
		ID: req.Id,
	}
	record, err := cl.recordSvc.Query(ctx, svcReq)
	if err != nil {
		errCode, ok := err.(*errors.ErrorCode)
		if !ok {
			// todo statsd
			logger.Error(context.Background(), "RecordController.Query unknown error. err:%s", err)
			return &record_pb.Record{}, status.Errorf(codes.Internal, err.Error())
		}
		switch errCode.Code {
		case codemsg.SelfDeinfeStatsu1:
			// 假如这种场景下返回正常.
			return &record_pb.Record{}, nil
		default:
			// todo statsd
			logger.Error(context.Background(),
				"RecordController.Query unknown error. err:%s", err)
			return &record_pb.Record{}, status.Errorf(codes.Internal, err.Error())
		}
	}
	return ToPbRecord(record), nil
}
