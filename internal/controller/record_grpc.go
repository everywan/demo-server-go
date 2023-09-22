package controller

import (
	"context"

	record_pb "github.com/everywan/demo-server-go/idl/proto/record"
	"github.com/everywan/demo-server-go/internal/dao"
	"github.com/everywan/demo-server-go/internal/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RecordGrpcController struct {
	// logger *flog.Logger// todo
	recordSvc service.RecordService
}

var _ record_pb.RecordServiceServer = new(RecordGrpcController)

func NewRecordGrpcController(recordSvc service.RecordService) *RecordGrpcController {
	return &RecordGrpcController{
		recordSvc: recordSvc,
	}
}

func (cl *RecordGrpcController) Create(ctx context.Context, req *record_pb.CreateRequset,
) (*record_pb.EmptyRespose, error) {
	svcRequest := &service.CreateRecordRequest{
		Name:      req.Name,
		Status:    dao.RecordStatus(req.Status),
		CreatedBy: req.CreatedBy,
	}
	err := cl.recordSvc.Create(ctx, svcRequest)
	if err != nil {
		// todo log error
	}
	return &record_pb.EmptyRespose{}, nil
}

func (cl *RecordGrpcController) Get(ctx context.Context, req *record_pb.IDRequest,
) (resp *record_pb.Record, err error) {
	if err := ValidateIDRequest(req); err != nil {
		return &record_pb.Record{}, status.Errorf(codes.InvalidArgument, err.Error())
	}
	record, err := cl.recordSvc.Get(ctx, uint(req.Id))
	if err != nil {
		// todo log error
	}
	return ToPbRecord(record), nil
}
