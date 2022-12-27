package logic

import (
	"context"

	"github.com/tqtcloud/manage/service/provider/aliyun/operator/rpc/internal/svc"
	"github.com/tqtcloud/manage/service/provider/aliyun/operator/rpc/types/operator"

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

func (l *HostDeleteLogic) HostDelete(in *operator.DeleteHostRequest) (*operator.DeleteHostResponse, error) {
	// todo: add your logic here and delete this line

	return &operator.DeleteHostResponse{}, nil
}
