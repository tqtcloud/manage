package logic

import (
	"context"

	"github.com/tqtcloud/manage/service/task/rpc/internal/svc"
	"github.com/tqtcloud/manage/service/task/rpc/types/task"

	"github.com/zeromicro/go-zero/core/logx"
)

type TaskCallbackLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTaskCallbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TaskCallbackLogic {
	return &TaskCallbackLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TaskCallbackLogic) TaskCallback(in *task.CallbackRequest) (*task.CallbackResponse, error) {
	// todo: add your logic here and delete this line
	l.Logger.Info("TaskCallback", in.TaskId, in.SecretId)
	return &task.CallbackResponse{}, nil
}
