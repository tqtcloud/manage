package logic

import (
	"context"

	"github.com/tqtcloud/manage/service/task/rpc/internal/svc"
	"github.com/tqtcloud/manage/service/task/rpc/types/task"

	"github.com/zeromicro/go-zero/core/logx"
)

type TaskDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTaskDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TaskDeleteLogic {
	return &TaskDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TaskDeleteLogic) TaskDelete(in *task.DeleteRequest) (*task.DeleteResponse, error) {
	// todo: add your logic here and delete this line

	return &task.DeleteResponse{}, nil
}
