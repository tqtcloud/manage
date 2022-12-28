// Code generated by goctl. DO NOT EDIT!
// Source: operator.proto

package operatorclient

import (
	"context"

	"github.com/tqtcloud/manage/service/provider/aliyun/operator/rpc/types/operator"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CreateHostRequest  = operator.CreateHostRequest
	CreateHostResponse = operator.CreateHostResponse
	DeleteHostRequest  = operator.DeleteHostRequest
	DeleteHostResponse = operator.DeleteHostResponse
	GetIdHostRequest   = operator.GetIdHostRequest
	GetListRequest     = operator.GetListRequest
	GetListResponse    = operator.GetListResponse

	Operator interface {
		HostSync(ctx context.Context, in *CreateHostRequest, opts ...grpc.CallOption) (*GetListResponse, error)
		HostDelete(ctx context.Context, in *DeleteHostRequest, opts ...grpc.CallOption) (*DeleteHostResponse, error)
		HostUpdate(ctx context.Context, in *CreateHostRequest, opts ...grpc.CallOption) (*CreateHostResponse, error)
		HostList(ctx context.Context, in *GetListRequest, opts ...grpc.CallOption) (*GetListResponse, error)
		HostGetId(ctx context.Context, in *GetIdHostRequest, opts ...grpc.CallOption) (*DeleteHostResponse, error)
	}

	defaultOperator struct {
		cli zrpc.Client
	}
)

func NewOperator(cli zrpc.Client) Operator {
	return &defaultOperator{
		cli: cli,
	}
}

func (m *defaultOperator) HostSync(ctx context.Context, in *CreateHostRequest, opts ...grpc.CallOption) (*GetListResponse, error) {
	client := operator.NewOperatorClient(m.cli.Conn())
	return client.HostSync(ctx, in, opts...)
}

func (m *defaultOperator) HostDelete(ctx context.Context, in *DeleteHostRequest, opts ...grpc.CallOption) (*DeleteHostResponse, error) {
	client := operator.NewOperatorClient(m.cli.Conn())
	return client.HostDelete(ctx, in, opts...)
}

func (m *defaultOperator) HostUpdate(ctx context.Context, in *CreateHostRequest, opts ...grpc.CallOption) (*CreateHostResponse, error) {
	client := operator.NewOperatorClient(m.cli.Conn())
	return client.HostUpdate(ctx, in, opts...)
}

func (m *defaultOperator) HostList(ctx context.Context, in *GetListRequest, opts ...grpc.CallOption) (*GetListResponse, error) {
	client := operator.NewOperatorClient(m.cli.Conn())
	return client.HostList(ctx, in, opts...)
}

func (m *defaultOperator) HostGetId(ctx context.Context, in *GetIdHostRequest, opts ...grpc.CallOption) (*DeleteHostResponse, error) {
	client := operator.NewOperatorClient(m.cli.Conn())
	return client.HostGetId(ctx, in, opts...)
}