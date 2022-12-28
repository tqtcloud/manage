package host

import (
	"context"
	"github.com/tqtcloud/manage/service/front-api/internal/svc"
	"github.com/tqtcloud/manage/service/front-api/internal/types"
	"github.com/tqtcloud/manage/service/provider/aliyun/operator/rpc/types/operator"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateHostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateHostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateHostLogic {
	return &UpdateHostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateHostLogic) UpdateHost(req *types.UpdateHostRequest) (resp *types.HostResponse, err error) {
	_, err = l.svcCtx.AliOperatorRpc.HostUpdate(l.ctx, &operator.CreateHostRequest{AccessKeyId: req.InstanceId})
	if err != nil {
		l.Logger.Errorf("AliOperatorRpc HostUpdate 更新错误 %s ", err)
		return nil, err
	}
	return nil, err
}
