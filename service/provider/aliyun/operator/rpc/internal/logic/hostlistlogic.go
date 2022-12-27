package logic

import (
	"context"

	"github.com/tqtcloud/manage/service/provider/aliyun/operator/rpc/internal/svc"
	"github.com/tqtcloud/manage/service/provider/aliyun/operator/rpc/types/operator"

	"github.com/zeromicro/go-zero/core/logx"
)

type HostListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHostListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HostListLogic {
	return &HostListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *HostListLogic) HostList(in *operator.GetListRequest) (*operator.GetListResponse, error) {
	// todo: add your logic here and delete this line

	return &operator.GetListResponse{}, nil
}
