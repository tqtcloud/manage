// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.4
// source: host/rpc/host.proto

package host

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

// HostClient is the client API for Host service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HostClient interface {
	HostSync(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*GetListResponse, error)
	HostDelete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	HostUpdate(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	HostList(ctx context.Context, in *GetListRequest, opts ...grpc.CallOption) (*GetListResponse, error)
	HostGetId(ctx context.Context, in *GetIdRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
}

type hostClient struct {
	cc grpc.ClientConnInterface
}

func NewHostClient(cc grpc.ClientConnInterface) HostClient {
	return &hostClient{cc}
}

func (c *hostClient) HostSync(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*GetListResponse, error) {
	out := new(GetListResponse)
	err := c.cc.Invoke(ctx, "/host.Host/HostSync", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hostClient) HostDelete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/host.Host/HostDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hostClient) HostUpdate(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/host.Host/HostUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hostClient) HostList(ctx context.Context, in *GetListRequest, opts ...grpc.CallOption) (*GetListResponse, error) {
	out := new(GetListResponse)
	err := c.cc.Invoke(ctx, "/host.Host/HostList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hostClient) HostGetId(ctx context.Context, in *GetIdRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/host.Host/HostGetId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HostServer is the server API for Host service.
// All implementations must embed UnimplementedHostServer
// for forward compatibility
type HostServer interface {
	HostSync(context.Context, *CreateRequest) (*GetListResponse, error)
	HostDelete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	HostUpdate(context.Context, *CreateRequest) (*CreateResponse, error)
	HostList(context.Context, *GetListRequest) (*GetListResponse, error)
	HostGetId(context.Context, *GetIdRequest) (*DeleteResponse, error)
	mustEmbedUnimplementedHostServer()
}

// UnimplementedHostServer must be embedded to have forward compatible implementations.
type UnimplementedHostServer struct {
}

func (UnimplementedHostServer) HostSync(context.Context, *CreateRequest) (*GetListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HostSync not implemented")
}
func (UnimplementedHostServer) HostDelete(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HostDelete not implemented")
}
func (UnimplementedHostServer) HostUpdate(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HostUpdate not implemented")
}
func (UnimplementedHostServer) HostList(context.Context, *GetListRequest) (*GetListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HostList not implemented")
}
func (UnimplementedHostServer) HostGetId(context.Context, *GetIdRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HostGetId not implemented")
}
func (UnimplementedHostServer) mustEmbedUnimplementedHostServer() {}

// UnsafeHostServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HostServer will
// result in compilation errors.
type UnsafeHostServer interface {
	mustEmbedUnimplementedHostServer()
}

func RegisterHostServer(s grpc.ServiceRegistrar, srv HostServer) {
	s.RegisterService(&Host_ServiceDesc, srv)
}

func _Host_HostSync_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HostServer).HostSync(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/host.Host/HostSync",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HostServer).HostSync(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Host_HostDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HostServer).HostDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/host.Host/HostDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HostServer).HostDelete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Host_HostUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HostServer).HostUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/host.Host/HostUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HostServer).HostUpdate(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Host_HostList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HostServer).HostList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/host.Host/HostList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HostServer).HostList(ctx, req.(*GetListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Host_HostGetId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HostServer).HostGetId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/host.Host/HostGetId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HostServer).HostGetId(ctx, req.(*GetIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Host_ServiceDesc is the grpc.ServiceDesc for Host service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Host_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "host.Host",
	HandlerType: (*HostServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HostSync",
			Handler:    _Host_HostSync_Handler,
		},
		{
			MethodName: "HostDelete",
			Handler:    _Host_HostDelete_Handler,
		},
		{
			MethodName: "HostUpdate",
			Handler:    _Host_HostUpdate_Handler,
		},
		{
			MethodName: "HostList",
			Handler:    _Host_HostList_Handler,
		},
		{
			MethodName: "HostGetId",
			Handler:    _Host_HostGetId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "host/rpc/host.proto",
}
