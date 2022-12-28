package user

import (
	"context"
	"github.com/tqtcloud/manage/service/user/rpc/types/user"

	"github.com/tqtcloud/manage/service/front-api/internal/svc"
	"github.com/tqtcloud/manage/service/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserListLogic {
	return &UserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserListLogic) UserList(req *types.UserListRequest) (resp []*types.UserListResponse, err error) {
	res, err := l.svcCtx.UserRpc.UserList(l.ctx, &user.UserListRequest{
		Page:  req.Page,
		Limit: req.Limit,
	})

	if err != nil {
		return nil, err
	}

	userList := make([]*types.UserListResponse, 0)
	for _, item := range res.Data {
		userList = append(userList, &types.UserListResponse{
			Id:     item.Id,
			Name:   item.Name,
			Gender: item.Gender,
			Email:  item.Email,
			Mobile: item.Mobile,
		})
	}

	return userList, nil
}
