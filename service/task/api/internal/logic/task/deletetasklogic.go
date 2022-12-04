package task

import (
	"context"

	"github.com/tqtcloud/manage/service/task/api/internal/svc"
	"github.com/tqtcloud/manage/service/task/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteTaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteTaskLogic {
	return &DeleteTaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteTaskLogic) DeleteTask(req *types.DeleteRequest) (resp *types.DeleteResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
