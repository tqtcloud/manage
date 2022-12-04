package task

import (
	"context"

	"github.com/tqtcloud/manage/service/task/api/internal/svc"
	"github.com/tqtcloud/manage/service/task/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetIdTaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetIdTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetIdTaskLogic {
	return &GetIdTaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetIdTaskLogic) GetIdTask(req *types.GetIdRequest) (resp *types.GetListResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
