// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: scheduled.proto

package proto_scheduled

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
	ScheduledService_GetScheduleds_FullMethodName   = "/ScheduledService/GetScheduleds"
	ScheduledService_GetScheduled_FullMethodName    = "/ScheduledService/GetScheduled"
	ScheduledService_CreateScheduled_FullMethodName = "/ScheduledService/CreateScheduled"
)

// ScheduledServiceClient is the client API for ScheduledService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ScheduledServiceClient interface {
	GetScheduleds(ctx context.Context, in *GetScheduledsReq, opts ...grpc.CallOption) (*GetScheduledsRes, error)
	GetScheduled(ctx context.Context, in *GetScheduledReq, opts ...grpc.CallOption) (*GetScheduledRes, error)
	CreateScheduled(ctx context.Context, in *CreateScheduledReq, opts ...grpc.CallOption) (*CreateScheduledRes, error)
}

type scheduledServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewScheduledServiceClient(cc grpc.ClientConnInterface) ScheduledServiceClient {
	return &scheduledServiceClient{cc}
}

func (c *scheduledServiceClient) GetScheduleds(ctx context.Context, in *GetScheduledsReq, opts ...grpc.CallOption) (*GetScheduledsRes, error) {
	out := new(GetScheduledsRes)
	err := c.cc.Invoke(ctx, ScheduledService_GetScheduleds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scheduledServiceClient) GetScheduled(ctx context.Context, in *GetScheduledReq, opts ...grpc.CallOption) (*GetScheduledRes, error) {
	out := new(GetScheduledRes)
	err := c.cc.Invoke(ctx, ScheduledService_GetScheduled_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scheduledServiceClient) CreateScheduled(ctx context.Context, in *CreateScheduledReq, opts ...grpc.CallOption) (*CreateScheduledRes, error) {
	out := new(CreateScheduledRes)
	err := c.cc.Invoke(ctx, ScheduledService_CreateScheduled_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ScheduledServiceServer is the server API for ScheduledService service.
// All implementations must embed UnimplementedScheduledServiceServer
// for forward compatibility
type ScheduledServiceServer interface {
	GetScheduleds(context.Context, *GetScheduledsReq) (*GetScheduledsRes, error)
	GetScheduled(context.Context, *GetScheduledReq) (*GetScheduledRes, error)
	CreateScheduled(context.Context, *CreateScheduledReq) (*CreateScheduledRes, error)
	mustEmbedUnimplementedScheduledServiceServer()
}

// UnimplementedScheduledServiceServer must be embedded to have forward compatible implementations.
type UnimplementedScheduledServiceServer struct {
}

func (UnimplementedScheduledServiceServer) GetScheduleds(context.Context, *GetScheduledsReq) (*GetScheduledsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetScheduleds not implemented")
}
func (UnimplementedScheduledServiceServer) GetScheduled(context.Context, *GetScheduledReq) (*GetScheduledRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetScheduled not implemented")
}
func (UnimplementedScheduledServiceServer) CreateScheduled(context.Context, *CreateScheduledReq) (*CreateScheduledRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateScheduled not implemented")
}
func (UnimplementedScheduledServiceServer) mustEmbedUnimplementedScheduledServiceServer() {}

// UnsafeScheduledServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ScheduledServiceServer will
// result in compilation errors.
type UnsafeScheduledServiceServer interface {
	mustEmbedUnimplementedScheduledServiceServer()
}

func RegisterScheduledServiceServer(s grpc.ServiceRegistrar, srv ScheduledServiceServer) {
	s.RegisterService(&ScheduledService_ServiceDesc, srv)
}

func _ScheduledService_GetScheduleds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetScheduledsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScheduledServiceServer).GetScheduleds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ScheduledService_GetScheduleds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScheduledServiceServer).GetScheduleds(ctx, req.(*GetScheduledsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScheduledService_GetScheduled_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetScheduledReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScheduledServiceServer).GetScheduled(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ScheduledService_GetScheduled_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScheduledServiceServer).GetScheduled(ctx, req.(*GetScheduledReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScheduledService_CreateScheduled_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateScheduledReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScheduledServiceServer).CreateScheduled(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ScheduledService_CreateScheduled_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScheduledServiceServer).CreateScheduled(ctx, req.(*CreateScheduledReq))
	}
	return interceptor(ctx, in, info, handler)
}

// ScheduledService_ServiceDesc is the grpc.ServiceDesc for ScheduledService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ScheduledService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ScheduledService",
	HandlerType: (*ScheduledServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetScheduleds",
			Handler:    _ScheduledService_GetScheduleds_Handler,
		},
		{
			MethodName: "GetScheduled",
			Handler:    _ScheduledService_GetScheduled_Handler,
		},
		{
			MethodName: "CreateScheduled",
			Handler:    _ScheduledService_CreateScheduled_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "scheduled.proto",
}
