package logic

import (
	"context"
	"github.com/tqtcloud/manage/service/provider/aliyun/host/rpc/internal/svc"
	"github.com/tqtcloud/manage/service/provider/aliyun/host/rpc/types/host"

	"github.com/zeromicro/go-zero/core/logx"
)

type HostDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHostDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HostDeleteLogic {
	return &HostDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *HostDeleteLogic) HostDelete(in *host.DeleteRequest) (*host.DeleteResponse, error) {
	// todo: add your logic here and delete this line

	return &host.DeleteResponse{}, nil
}
