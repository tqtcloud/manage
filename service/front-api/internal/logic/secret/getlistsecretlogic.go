package secret

import (
	"context"
	"github.com/tqtcloud/manage/service/secret/rpc/types/secret"

	"github.com/tqtcloud/manage/service/front-api/internal/svc"
	"github.com/tqtcloud/manage/service/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetListSecretLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetListSecretLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetListSecretLogic {
	return &GetListSecretLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetListSecretLogic) GetListSecret(req *types.GetSecretListRequest) (resp []*types.GetSecretListResponse, err error) {
	res, err := l.svcCtx.SecretRpc.SecretList(l.ctx, &secret.GetListRequest{
		Page:  req.Page,
		Limit: req.Limit,
	})

	if err != nil {
		return nil, err
	}

	secretList := make([]*types.GetSecretListResponse, 0)
	for _, item := range res.Data {
		secretList = append(secretList, &types.GetSecretListResponse{
			Id:              item.Id,
			Vendor:          item.Vendor,
			AccessKeyId:     item.AccessKeyId,
			AccessKeySecret: item.AccessKeySecret,
		})
	}

	return secretList, nil
}
