// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.4
// source: rpc/task.proto

package task

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

// TaskClient is the client API for Task service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TaskClient interface {
	TaskCreate(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	TaskDelete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	TaskList(ctx context.Context, in *GetListRequest, opts ...grpc.CallOption) (*GetListResponse, error)
	TaskGetId(ctx context.Context, in *GetIdRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	TaskCallback(ctx context.Context, in *CallbackRequest, opts ...grpc.CallOption) (*CallbackResponse, error)
}

type taskClient struct {
	cc grpc.ClientConnInterface
}

func NewTaskClient(cc grpc.ClientConnInterface) TaskClient {
	return &taskClient{cc}
}

func (c *taskClient) TaskCreate(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/task.Task/TaskCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskClient) TaskDelete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/task.Task/TaskDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskClient) TaskList(ctx context.Context, in *GetListRequest, opts ...grpc.CallOption) (*GetListResponse, error) {
	out := new(GetListResponse)
	err := c.cc.Invoke(ctx, "/task.Task/TaskList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskClient) TaskGetId(ctx context.Context, in *GetIdRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/task.Task/TaskGetId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskClient) TaskCallback(ctx context.Context, in *CallbackRequest, opts ...grpc.CallOption) (*CallbackResponse, error) {
	out := new(CallbackResponse)
	err := c.cc.Invoke(ctx, "/task.Task/TaskCallback", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TaskServer is the server API for Task service.
// All implementations must embed UnimplementedTaskServer
// for forward compatibility
type TaskServer interface {
	TaskCreate(context.Context, *CreateRequest) (*CreateResponse, error)
	TaskDelete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	TaskList(context.Context, *GetListRequest) (*GetListResponse, error)
	TaskGetId(context.Context, *GetIdRequest) (*DeleteResponse, error)
	TaskCallback(context.Context, *CallbackRequest) (*CallbackResponse, error)
	mustEmbedUnimplementedTaskServer()
}

// UnimplementedTaskServer must be embedded to have forward compatible implementations.
type UnimplementedTaskServer struct {
}

func (UnimplementedTaskServer) TaskCreate(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TaskCreate not implemented")
}
func (UnimplementedTaskServer) TaskDelete(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TaskDelete not implemented")
}
func (UnimplementedTaskServer) TaskList(context.Context, *GetListRequest) (*GetListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TaskList not implemented")
}
func (UnimplementedTaskServer) TaskGetId(context.Context, *GetIdRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TaskGetId not implemented")
}
func (UnimplementedTaskServer) TaskCallback(context.Context, *CallbackRequest) (*CallbackResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TaskCallback not implemented")
}
func (UnimplementedTaskServer) mustEmbedUnimplementedTaskServer() {}

// UnsafeTaskServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TaskServer will
// result in compilation errors.
type UnsafeTaskServer interface {
	mustEmbedUnimplementedTaskServer()
}

func RegisterTaskServer(s grpc.ServiceRegistrar, srv TaskServer) {
	s.RegisterService(&Task_ServiceDesc, srv)
}

func _Task_TaskCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServer).TaskCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/task.Task/TaskCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServer).TaskCreate(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Task_TaskDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServer).TaskDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/task.Task/TaskDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServer).TaskDelete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Task_TaskList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServer).TaskList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/task.Task/TaskList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServer).TaskList(ctx, req.(*GetListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Task_TaskGetId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServer).TaskGetId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/task.Task/TaskGetId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServer).TaskGetId(ctx, req.(*GetIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Task_TaskCallback_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CallbackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServer).TaskCallback(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/task.Task/TaskCallback",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServer).TaskCallback(ctx, req.(*CallbackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Task_ServiceDesc is the grpc.ServiceDesc for Task service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Task_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "task.Task",
	HandlerType: (*TaskServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TaskCreate",
			Handler:    _Task_TaskCreate_Handler,
		},
		{
			MethodName: "TaskDelete",
			Handler:    _Task_TaskDelete_Handler,
		},
		{
			MethodName: "TaskList",
			Handler:    _Task_TaskList_Handler,
		},
		{
			MethodName: "TaskGetId",
			Handler:    _Task_TaskGetId_Handler,
		},
		{
			MethodName: "TaskCallback",
			Handler:    _Task_TaskCallback_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc/task.proto",
}
