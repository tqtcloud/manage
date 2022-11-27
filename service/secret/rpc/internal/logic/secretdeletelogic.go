package logic

import (
	"context"
	"github.com/tqtcloud/manage/service/secret/model"
	"github.com/tqtcloud/manage/service/secret/rpc/internal/svc"
	"github.com/tqtcloud/manage/service/secret/rpc/types/secret"
	"github.com/tqtcloud/resp/errorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type SecretDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSecretDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SecretDeleteLogic {
	return &SecretDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SecretDeleteLogic) SecretDelete(in *secret.DeleteRequest) (*secret.DeleteResponse, error) {
	// 查询需要删除的 ID 数据
	resp, err := l.svcCtx.SecretModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, errorx.NewCodeError(1010, "Secret 不存在")
		}
		return nil, errorx.NewCodeError(1011, err.Error())
	}

	err = l.svcCtx.SecretModel.Delete(l.ctx, in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, errorx.NewCodeError(1010, "Secret 不存在")
		}
		return nil, errorx.NewCodeError(1011, err.Error())
	}

	return &secret.DeleteResponse{
		Id:              resp.Id,
		Vendor:          resp.Vendor,
		AccessKeyId:     resp.AccessKeyId,
		AccessKeySecret: resp.AccessKeySecret,
	}, nil
}
