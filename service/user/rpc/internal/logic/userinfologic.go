package logic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/tqtcloud/manage/common/xerr"
	"github.com/tqtcloud/manage/service/user/model"
	"github.com/tqtcloud/manage/service/user/rpc/internal/svc"
	"github.com/tqtcloud/manage/service/user/rpc/types/user"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserInfoLogic) UserInfo(in *user.UserInfoRequest) (*user.UserInfoResponse, error) {
	// 查询用户是否存在
	resp, err := l.svcCtx.UserModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.UsernameNoExistError), "用户不存在错误: err:%v,user:%+v", err, in.Id)
		}
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "未知错误: err:%v,user:%+v", err, in.Id)
	}

	return &user.UserInfoResponse{
		Id:     resp.Id,
		Name:   resp.Name,
		Gender: resp.Gender,
		Mobile: resp.Mobile,
		Email:  resp.Email,
	}, nil
}
