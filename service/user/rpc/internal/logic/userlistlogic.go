package logic

import (
	"context"
	"github.com/tqtcloud/manage/service/user/model"
	"github.com/tqtcloud/manage/service/user/rpc/internal/svc"
	"github.com/tqtcloud/manage/service/user/rpc/types/user"
	"github.com/tqtcloud/resp/errorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserListLogic {
	return &UserListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserListLogic) UserList(in *user.UserListRequest) (*user.UserListResponse, error) {
	list, err := l.svcCtx.UserModel.FindAllPage(l.ctx, in)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, errorx.NewDefaultError(err.Error())
		}
		return nil, errorx.NewDefaultError(err.Error())
	}

	orderList := make([]*user.UserInfoResponse, 0)
	for _, item := range list {
		orderList = append(orderList, &user.UserInfoResponse{
			Id:     item.Id,
			Name:   item.Name,
			Gender: item.Gender,
			Email:  item.Email,
			Mobile: item.Mobile,
		})
	}

	return &user.UserListResponse{
		Data: orderList,
	}, nil
}
