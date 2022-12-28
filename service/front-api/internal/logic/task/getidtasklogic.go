package task

import (
	"context"
	"github.com/tqtcloud/manage/service/front-api/internal/svc"
	"github.com/tqtcloud/manage/service/front-api/internal/types"
	"github.com/tqtcloud/manage/service/task/rpc/types/task"

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

func (l *GetIdTaskLogic) GetIdTask(req *types.GetTaskIdRequest) (resp *types.GetTaskListResponse, err error) {
	res, err := l.svcCtx.TaskRpc.TaskGetId(l.ctx, &task.GetIdRequest{Id: req.Id})
	if err != nil {
		l.Logger.Errorf("TaskGetId 查询错误 %s ", err)
		return nil, err
	}

	return &types.GetTaskListResponse{
		Id:           res.Id,
		TaskName:     res.TaskName,
		Vendor:       string(res.Vendor),
		TaskType:     string(res.TaskType),
		SecretId:     res.SecretId,
		Region:       res.Region,
		TaskUser:     res.TaskUser,
		Status:       res.Status,
		Message:      res.Message,
		Start_At:     res.Start_At,
		TotalSucceed: res.TotalSucceed,
		TotalFailed:  res.TotalFailed,
		CreateTime:   res.CreateTime,
		UpdateTime:   res.UpdateTime,
	}, nil
}
