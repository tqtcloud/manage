package logic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/tqtcloud/manage/common/desencryption"
	"github.com/tqtcloud/manage/common/xerr"
	"github.com/tqtcloud/manage/service/secret/rpc/internal/svc"
	"github.com/tqtcloud/manage/service/secret/rpc/types/secret"
	"github.com/tqtcloud/manage/service/user/model"
	"github.com/zeromicro/go-zero/core/logx"
)

type SecretUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSecretUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SecretUpdateLogic {
	return &SecretUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SecretUpdateLogic) SecretUpdate(in *secret.UpdateRequest) (*secret.UpdateResponse, error) {
	res, err := l.svcCtx.SecretModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.SecretIDNoExistError), "数据为 err:%v,Secret:%+v", err, in.Id)
		}
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "其他错误 err:%v", err)
	}

	if in.Vendor != "" {
		res.Vendor = in.Vendor
	}
	if in.AccessKeyId != "" {
		res.AccessKeyId = in.AccessKeyId
	}
	if in.AccessKeySecret != "" {
		sk, _ := desencryption.Encrypt(in.AccessKeySecret, []byte(l.svcCtx.Config.Salt))
		res.AccessKeySecret = sk
	}

	err = l.svcCtx.SecretModel.Update(l.ctx, res)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.SecretUpdateError), "更新失败 err:%v,Secret:%+v", err, in.Id)
	}

	return &secret.UpdateResponse{
		Id:              in.Id,
		Vendor:          res.Vendor,
		AccessKeyId:     res.AccessKeyId,
		AccessKeySecret: res.AccessKeySecret,
	}, nil
}
