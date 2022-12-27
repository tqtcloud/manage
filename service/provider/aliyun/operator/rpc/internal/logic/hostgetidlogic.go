package logic

import (
	"context"

	"github.com/tqtcloud/manage/service/provider/aliyun/operator/rpc/internal/svc"
	"github.com/tqtcloud/manage/service/provider/aliyun/operator/rpc/types/operator"

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

func (l *HostGetIdLogic) HostGetId(in *operator.GetIdHostRequest) (*operator.DeleteHostResponse, error) {
	// todo: add your logic here and delete this line

	return &operator.DeleteHostResponse{}, nil
}
