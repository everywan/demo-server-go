package controller

import (
	"errors"

	record_pb "github.com/everywan/demo-server-go/idl/proto/record"
	"github.com/everywan/demo-server-go/internal/service"
)

// -------------------- grpc ----------------------
func ValidateIDRequest(req *record_pb.IDRequest) error {
	if req.GetId() < 1 {
		return errors.New("must have id")
	}
	return nil
}

func ToPbRecord(record *service.Record) *record_pb.Record {
	return &record_pb.Record{
		Id:        uint64(record.ID),
		Name:      record.Name,
		Status:    record_pb.RecordStatus(record.Status),
		CreatedAt: record.CreatedAt,
		UpdatedAt: record.UpdatedAt,
		CreatedBy: record.CreatedBy,
		UpdatedBy: record.UpdatedBy,
	}

}
