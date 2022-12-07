package logic

import (
	"context"
	"github.com/tqtcloud/manage/service/provider/aliyun/host/rpc/internal/svc"
	"github.com/tqtcloud/manage/service/provider/aliyun/host/rpc/types/host"

	"github.com/zeromicro/go-zero/core/logx"
)

type HostGetIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHostGetIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HostGetIdLogic {
	return &HostGetIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *HostGetIdLogic) HostGetId(in *host.GetIdRequest) (*host.DeleteResponse, error) {
	// todo: add your logic here and delete this line

	return &host.DeleteResponse{}, nil
}
