package logic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/rs/xid"
	"github.com/tqtcloud/manage/common/desencryption"
	"github.com/tqtcloud/manage/common/xerr"
	"github.com/tqtcloud/manage/service/secret/model"
	"github.com/tqtcloud/manage/service/secret/rpc/internal/svc"
	"github.com/tqtcloud/manage/service/secret/rpc/types/secret"
	"github.com/zeromicro/go-zero/core/logx"
)

type SecretCreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSecretCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SecretCreateLogic {
	return &SecretCreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SecretCreateLogic) SecretCreate(in *secret.CreateRequest) (*secret.CreateResponse, error) {
	_, err := l.svcCtx.SecretModel.FindOneByAccessKeyId(l.ctx, in.AccessKeyId)
	if err == nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.SecretIDExistError), "数据为 err:%v,Secret:%+v", err, in.AccessKeyId)
	}
	sk, _ := desencryption.Encrypt(in.AccessKeySecret, []byte(l.svcCtx.Config.Salt))
	if err == model.ErrNotFound {
		newSecret := model.Secret{
			Id:              xid.New().String(),
			Vendor:          in.Vendor,
			AccessKeyId:     in.AccessKeyId,
			AccessKeySecret: sk,
		}

		resp, err := l.svcCtx.SecretModel.Insert(l.ctx, &newSecret)
		if err != nil {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.SecretDbInsertError), "数据为 err:%v,厂商：%s,AccessKeyId：%s,AccessKeySecret：%s", err, newSecret.Vendor, newSecret.AccessKeyId, newSecret.AccessKeySecret)
		}

		rows, err := resp.RowsAffected()
		if err != nil {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.SecretDbInsertError), "数据库递增错误 err:%v,Secret:%+v", err, rows)
		}
		return &secret.CreateResponse{
			Id:              newSecret.Id,
			Vendor:          newSecret.Vendor,
			AccessKeyId:     newSecret.AccessKeyId,
			AccessKeySecret: newSecret.AccessKeySecret,
		}, nil
	}

	return nil, xerr.NewErrMsg("默认错误请联系管理员")
}
