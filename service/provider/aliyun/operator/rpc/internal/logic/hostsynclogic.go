package logic

import (
	"context"

	"github.com/tqtcloud/manage/service/provider/aliyun/operator/rpc/internal/svc"
	"github.com/tqtcloud/manage/service/provider/aliyun/operator/rpc/types/operator"

	"github.com/zeromicro/go-zero/core/logx"
)

type HostSyncLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHostSyncLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HostSyncLogic {
	return &HostSyncLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *HostSyncLogic) HostSync(in *operator.CreateHostRequest) (*operator.GetListResponse, error) {
	// todo: add your logic here and delete this line

	return &operator.GetListResponse{}, nil
}
