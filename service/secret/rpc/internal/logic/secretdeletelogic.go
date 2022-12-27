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
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.SecretIDNoExistError), "数据为 err:%v,Secret:%+v", err, in.Id)
		}
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "其他错误 err:%v,Secret:%+v", err, in.Id)
	}

	err = l.svcCtx.SecretModel.Delete(l.ctx, in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.SecretIDNoExistError), "数据为 err:%v,Secret:%+v", err, in.Id)
		}
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "数据为 err:%v,Secret:%+v", err, in.Id)
	}

	return &secret.DeleteResponse{
		Id:              resp.Id,
		Vendor:          resp.Vendor,
		AccessKeyId:     resp.AccessKeyId,
		AccessKeySecret: resp.AccessKeySecret,
	}, nil
}
