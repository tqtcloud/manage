package secret

import (
	"context"
	"github.com/tqtcloud/manage/service/secret/rpc/types/secret"

	"github.com/tqtcloud/manage/service/front-api/internal/svc"
	"github.com/tqtcloud/manage/service/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSecretLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateSecretLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSecretLogic {
	return &UpdateSecretLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateSecretLogic) UpdateSecret(req *types.UpdateSecretRequest) (resp *types.UpdateSecretResponse, err error) {
	res, err := l.svcCtx.SecretRpc.SecretUpdate(l.ctx, &secret.UpdateRequest{
		Id:              req.Id,
		Vendor:          req.Vendor,
		AccessKeyId:     req.AccessKeyId,
		AccessKeySecret: req.AccessKeySecret,
	})
	if err != nil {
		return nil, err
	}
	return &types.UpdateSecretResponse{
		Id:              res.Id,
		Vendor:          res.Vendor,
		AccessKeyId:     res.AccessKeyId,
		AccessKeySecret: res.AccessKeySecret,
	}, nil
}
