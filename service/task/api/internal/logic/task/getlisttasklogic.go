package task

import (
	"context"

	"github.com/tqtcloud/manage/service/task/api/internal/svc"
	"github.com/tqtcloud/manage/service/task/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetListTaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetListTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetListTaskLogic {
	return &GetListTaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetListTaskLogic) GetListTask(req *types.GetListRequest) (resp *types.GetListResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
