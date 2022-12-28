package logic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/tqtcloud/manage/common/xerr"
	"github.com/tqtcloud/manage/service/provider/aliyun/operator/rpc/internal/svc"
	"github.com/tqtcloud/manage/service/provider/aliyun/operator/rpc/types/operator"

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

func (l *HostUpdateLogic) HostUpdate(in *operator.CreateHostRequest) (*operator.CreateHostResponse, error) {
	return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "云上资产不允许本地手动更新")
}
