package task

import (
	"context"
	"github.com/tqtcloud/manage/common/errorx"
	"github.com/tqtcloud/manage/service/task/rpc/types/task"

	"github.com/tqtcloud/manage/service/front-api/internal/svc"
	"github.com/tqtcloud/manage/service/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateTaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateTaskLogic {
	return &CreateTaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateTaskLogic) CreateTask(req *types.CreateTaskRequest) (resp *types.CreateTaskResponse, err error) {
	res, err := l.svcCtx.TaskRpc.TaskCreate(l.ctx, &task.CreateRequest{
		TaskName: req.TaskName,
		Vendor:   task.Vendor(req.Vendor),
		TaskType: task.TaskType(req.TaskType),
		Region:   req.Region,
		SecretId: req.SecretId,
		UserId:   req.UserId,
	})
	if err != nil {
		return nil, errorx.NewDefaultError(err.Error())
	}

	return &types.CreateTaskResponse{
		Id:       res.Id,
		TaskName: res.TaskName,
		Vendor:   string(res.Vendor),
		TaskType: string(res.TaskType),
		SecretId: req.SecretId,
		Region:   res.Region,
		TaskUser: res.TaskUser,
		Status:   res.Status,
		Message:  res.Message,
		Start_At: res.Start_At,
	}, nil
}
