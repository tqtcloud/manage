package secret

import (
	"context"
	"github.com/tqtcloud/manage/service/secret/rpc/types/secret"

	"github.com/tqtcloud/manage/service/secret/api/internal/svc"
	"github.com/tqtcloud/manage/service/secret/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetIdSecretLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetIdSecretLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetIdSecretLogic {
	return &GetIdSecretLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetIdSecretLogic) GetIdSecret(req *types.GetIdRequest) (resp *types.GetListResponse, err error) {
	res, err := l.svcCtx.SecretRpc.SecretGetId(l.ctx, &secret.GetIdRequest{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return &types.GetListResponse{
		Id:              res.Id,
		AccessKeyId:     res.AccessKeyId,
		AccessKeySecret: res.AccessKeySecret,
	}, nil
}
