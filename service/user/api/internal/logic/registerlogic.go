package logic

import (
	"context"
	"github.com/tqtcloud/manage/service/user/api/internal/svc"
	"github.com/tqtcloud/manage/service/user/api/internal/types"
	"github.com/tqtcloud/manage/service/user/rpc/types/user"
	"github.com/tqtcloud/resp/errorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterRequest) (resp *types.RegisterResponse, err error) {
	res, err := l.svcCtx.UserRpc.Register(l.ctx, &user.RegisterRequest{
		Name:     req.Name,
		Gender:   req.Gender,
		Mobile:   req.Mobile,
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return nil, errorx.NewUserError(err.Error())
	}

	return &types.RegisterResponse{
		Id:     res.Id,
		Name:   res.Name,
		Gender: res.Gender,
		Mobile: res.Mobile,
		Email:  req.Email,
	}, nil
}
