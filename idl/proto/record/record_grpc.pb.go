// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.2
// source: idl/proto/record.proto

package record_pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	RecordService_Create_FullMethodName = "/everywan.demo_server_go.RecordService/Create"
	RecordService_Get_FullMethodName    = "/everywan.demo_server_go.RecordService/Get"
)

// RecordServiceClient is the client API for RecordService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RecordServiceClient interface {
	Create(ctx context.Context, in *CreateRequset, opts ...grpc.CallOption) (*EmptyRespose, error)
	Get(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*Record, error)
}

type recordServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRecordServiceClient(cc grpc.ClientConnInterface) RecordServiceClient {
	return &recordServiceClient{cc}
}

func (c *recordServiceClient) Create(ctx context.Context, in *CreateRequset, opts ...grpc.CallOption) (*EmptyRespose, error) {
	out := new(EmptyRespose)
	err := c.cc.Invoke(ctx, RecordService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recordServiceClient) Get(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*Record, error) {
	out := new(Record)
	err := c.cc.Invoke(ctx, RecordService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RecordServiceServer is the server API for RecordService service.
// All implementations should embed UnimplementedRecordServiceServer
// for forward compatibility
type RecordServiceServer interface {
	Create(context.Context, *CreateRequset) (*EmptyRespose, error)
	Get(context.Context, *IDRequest) (*Record, error)
}

// UnimplementedRecordServiceServer should be embedded to have forward compatible implementations.
type UnimplementedRecordServiceServer struct {
}

func (UnimplementedRecordServiceServer) Create(context.Context, *CreateRequset) (*EmptyRespose, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedRecordServiceServer) Get(context.Context, *IDRequest) (*Record, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}

// UnsafeRecordServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RecordServiceServer will
// result in compilation errors.
type UnsafeRecordServiceServer interface {
	mustEmbedUnimplementedRecordServiceServer()
}

func RegisterRecordServiceServer(s grpc.ServiceRegistrar, srv RecordServiceServer) {
	s.RegisterService(&RecordService_ServiceDesc, srv)
}

func _RecordService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequset)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecordServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RecordService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecordServiceServer).Create(ctx, req.(*CreateRequset))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecordService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecordServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RecordService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecordServiceServer).Get(ctx, req.(*IDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RecordService_ServiceDesc is the grpc.ServiceDesc for RecordService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RecordService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "everywan.demo_server_go.RecordService",
	HandlerType: (*RecordServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _RecordService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _RecordService_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "idl/proto/record.proto",
}