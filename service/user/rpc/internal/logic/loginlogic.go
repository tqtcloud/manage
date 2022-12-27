package logic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/tqtcloud/manage/common/cryptx"
	"github.com/tqtcloud/manage/common/xerr"
	"github.com/tqtcloud/manage/service/user/model"
	"github.com/tqtcloud/manage/service/user/rpc/internal/svc"
	"github.com/tqtcloud/manage/service/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginRequest) (*user.LoginResponse, error) {
	resp, err := l.svcCtx.UserModel.FindOneByName(l.ctx, in.Name)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "数据为 err:%v,user:%+v", err, in.Name)
		}
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.ServerCommonError), "其他错误：%s", err)
	}

	// 判断密码是否正确
	password := cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, in.Password)
	if password != resp.Password {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.UsernamePwdError), "密码匹配出错:%s", resp.Password)
	}

	return &user.LoginResponse{
		Id:     resp.Id,
		Name:   resp.Name,
		Gender: resp.Gender,
		Mobile: resp.Mobile,
		Email:  resp.Email,
	}, nil
}
