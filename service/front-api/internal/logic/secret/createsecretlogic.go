package secret

import (
	"context"
	"github.com/tqtcloud/manage/service/secret/rpc/types/secret"

	"github.com/tqtcloud/manage/service/front-api/internal/svc"
	"github.com/tqtcloud/manage/service/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateSecretLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateSecretLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateSecretLogic {
	return &CreateSecretLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateSecretLogic) CreateSecret(req *types.CreateSecretRequest) (resp *types.CreateSecretResponse, err error) {
	res, err := l.svcCtx.SecretRpc.SecretCreate(l.ctx, &secret.CreateRequest{
		Vendor:          req.Vendor,
		AccessKeyId:     req.AccessKeyId,
		AccessKeySecret: req.AccessKeySecret,
	})
	if err != nil {
		return nil, err
	}

	return &types.CreateSecretResponse{
		Id:              res.Id,
		AccessKeyId:     res.AccessKeyId,
		AccessKeySecret: res.AccessKeySecret,
	}, nil
}
