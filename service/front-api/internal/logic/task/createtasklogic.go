package task

import (
	"context"
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
	uid, _ := l.ctx.Value("uid").(string)
	logx.Infof("创建任务用户id：%s", uid)
	res, err := l.svcCtx.TaskRpc.TaskCreate(l.ctx, &task.CreateRequest{
		UserId:   uid,
		TaskName: req.TaskName,
		Vendor:   task.Vendor(req.Vendor),
		TaskType: task.TaskType(req.TaskType),
		Region:   req.Region,
		SecretId: req.SecretId,
	})
	if err != nil {
		return nil, err
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
