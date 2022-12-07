package logic

import (
	"context"
	"github.com/tqtcloud/manage/service/provider/aliyun/host/rpc/internal/svc"
	"github.com/tqtcloud/manage/service/provider/aliyun/host/rpc/types/host"

	"github.com/zeromicro/go-zero/core/logx"
)

type HostUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHostUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HostUpdateLogic {
	return &HostUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *HostUpdateLogic) HostUpdate(in *host.CreateRequest) (*host.CreateResponse, error) {
	// todo: add your logic here and delete this line

	return &host.CreateResponse{}, nil
}
