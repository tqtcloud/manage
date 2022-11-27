package logic

import (
	"context"
	"github.com/tqtcloud/manage/common/desencryption"
	"github.com/tqtcloud/manage/service/secret/model"
	"github.com/tqtcloud/manage/service/secret/rpc/internal/svc"
	"github.com/tqtcloud/manage/service/secret/rpc/types/secret"
	"github.com/tqtcloud/resp/errorx"

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
		return nil, errorx.NewDefaultError("AccessKeyId已存在,请重新添加")
	}
	sk, _ := desencryption.Encrypt(in.AccessKeySecret, []byte(l.svcCtx.Config.Salt))
	if err == model.ErrNotFound {
		newSecret := model.Secret{
			Vendor:          in.Vendor,
			AccessKeyId:     in.AccessKeyId,
			AccessKeySecret: sk,
		}

		resp, err := l.svcCtx.SecretModel.Insert(l.ctx, &newSecret)
		if err != nil {
			l.Logger.Infof("厂商：%s,AccessKeyId：%s,AccessKeySecret：%s", newSecret.Vendor, newSecret.AccessKeyId, newSecret.AccessKeySecret)
			l.Logger.Errorf("AK SK 添加错误: %s", err)
			return nil, errorx.NewUserError("内部错误请联系,请联系管理员")
		}

		newSecret.Id, err = resp.LastInsertId()
		if err != nil {
			l.Logger.Errorf("数据库递增错误: %s", err)
			return nil, errorx.NewUserError("内部错误请联系,请联系管理员")
		}
		return &secret.CreateResponse{
			Id:              newSecret.Id,
			Vendor:          newSecret.Vendor,
			AccessKeyId:     newSecret.AccessKeyId,
			AccessKeySecret: newSecret.AccessKeySecret,
		}, nil
	}

	return nil, errorx.NewDefaultError("内部错误请联系,请联系管理员")
}
