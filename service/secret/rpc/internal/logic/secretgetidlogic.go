package logic

import (
	"context"
	"github.com/tqtcloud/manage/service/secret/model"
	"github.com/tqtcloud/manage/service/secret/rpc/internal/svc"
	"github.com/tqtcloud/manage/service/secret/rpc/types/secret"
	"github.com/tqtcloud/resp/errorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type SecretGetIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSecretGetIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SecretGetIdLogic {
	return &SecretGetIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SecretGetIdLogic) SecretGetId(in *secret.GetIdRequest) (*secret.CreateResponse, error) {
	resp, err := l.svcCtx.SecretModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, errorx.NewCodeError(1011, "ID 不存在,重新输入")
		}
		return nil, errorx.NewCodeError(1011, err.Error())
	}

	return &secret.CreateResponse{
		Id:              resp.Id,
		Vendor:          resp.Vendor,
		AccessKeyId:     resp.AccessKeyId,
		AccessKeySecret: resp.AccessKeySecret,
	}, nil
}
