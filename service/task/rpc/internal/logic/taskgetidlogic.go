package logic

import (
	"context"

	"github.com/tqtcloud/manage/service/task/rpc/internal/svc"
	"github.com/tqtcloud/manage/service/task/rpc/types/task"

	"github.com/zeromicro/go-zero/core/logx"
)

type TaskGetIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTaskGetIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TaskGetIdLogic {
	return &TaskGetIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TaskGetIdLogic) TaskGetId(in *task.GetIdRequest) (*task.CreateResponse, error) {
	// todo: add your logic here and delete this line

	return &task.CreateResponse{}, nil
}
