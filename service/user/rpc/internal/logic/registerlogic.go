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

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterRequest) (*user.RegisterResponse, error) {
	// todo: add your logic here and delete this line
	// 判断Name是否已经注册
	_, err := l.svcCtx.UserModel.FindOneByName(l.ctx, in.Name)
	if err == nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.UsernameExistError), "数据为 err:%v,user:%+v", err, in.Name)
	}

	if err == model.ErrNotFound {
		newUser := model.User{
			Name:     in.Name,
			Gender:   in.Gender,
			Mobile:   in.Mobile,
			Email:    in.Email,
			Password: cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, in.Password),
		}

		resp, err := l.svcCtx.UserModel.Insert(l.ctx, &newUser)
		if err != nil {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "用户注册错误数据为 err:%v,user:%+v", err, in.Name)
		}

		newUser.Id, err = resp.LastInsertId()
		if err != nil {
			l.Logger.Errorf("数据库递增错误: %s", err)
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.UserDbInsertError), "数据库递增错误: err:%v,user:%+v", err, in.Name)
		}
		return &user.RegisterResponse{
			Id:     newUser.Id,
			Name:   newUser.Name,
			Gender: newUser.Gender,
			Mobile: newUser.Mobile,
			Email:  newUser.Email,
		}, nil
	}

	return nil, xerr.NewErrMsg("默认错误请联系管理员")
}
