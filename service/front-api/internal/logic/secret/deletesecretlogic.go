package secret

import (
	"context"
	"github.com/tqtcloud/manage/service/secret/rpc/types/secret"

	"github.com/tqtcloud/manage/service/front-api/internal/svc"
	"github.com/tqtcloud/manage/service/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteSecretLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteSecretLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSecretLogic {
	return &DeleteSecretLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteSecretLogic) DeleteSecret(req *types.DeleteSecretRequest) (resp *types.DeleteSecretResponse, err error) {
	res, err := l.svcCtx.SecretRpc.SecretDelete(l.ctx, &secret.DeleteRequest{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	return &types.DeleteSecretResponse{
		Id:              res.Id,
		Vendor:          res.Vendor,
		AccessKeyId:     res.AccessKeyId,
		AccessKeySecret: res.AccessKeySecret,
	}, nil
}
