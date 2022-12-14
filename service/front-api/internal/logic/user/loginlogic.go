package user

import (
	"context"
	"github.com/pkg/errors"
	"github.com/tqtcloud/manage/common/jwtx"
	"github.com/tqtcloud/manage/common/xerr"
	"github.com/tqtcloud/manage/service/user/rpc/types/user"
	"time"

	"github.com/tqtcloud/manage/service/front-api/internal/svc"
	"github.com/tqtcloud/manage/service/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	res, err := l.svcCtx.UserRpc.Login(l.ctx, &user.LoginRequest{
		Name:     req.Name,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.Auth.AccessExpire

	accessToken, err := jwtx.GetToken(l.svcCtx.Config.Auth.AccessSecret, now, accessExpire, res.Id)
	if err != nil {
		l.Logger.Errorf("GetToken error :%s", err.Error())
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.TokenGenerateError), "req: %+v,rpc err:%+v", req, err)
	}

	return &types.LoginResponse{
		Name:         res.Name,
		Gender:       res.Gender,
		Mobile:       res.Mobile,
		Email:        res.Email,
		AccessToken:  accessToken,
		AccessExpire: now + accessExpire,
	}, nil
}
