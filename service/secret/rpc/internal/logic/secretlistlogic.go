package logic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/tqtcloud/manage/common/xerr"
	"github.com/tqtcloud/manage/service/secret/model"
	"github.com/tqtcloud/manage/service/secret/rpc/internal/svc"
	"github.com/tqtcloud/manage/service/secret/rpc/types/secret"

	"github.com/zeromicro/go-zero/core/logx"
)

type SecretListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSecretListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SecretListLogic {
	return &SecretListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SecretListLogic) SecretList(in *secret.GetListRequest) (*secret.GetListResponse, error) {
	list, err := l.svcCtx.SecretModel.FindAllPage(l.ctx, in)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.SecretIDNoExistError), "数据不存在 err:%v", err)
		}
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "其他错误 err:%v", err)
	}

	secretList := make([]*secret.CreateResponse, 0)
	for _, item := range list {
		secretList = append(secretList, &secret.CreateResponse{
			Id:              item.Id,
			Vendor:          item.Vendor,
			AccessKeyId:     item.AccessKeyId,
			AccessKeySecret: item.AccessKeySecret,
		})
	}

	return &secret.GetListResponse{
		Data: secretList,
	}, nil
}
