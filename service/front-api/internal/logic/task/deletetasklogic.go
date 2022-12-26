package task

import (
	"context"
	"github.com/tqtcloud/manage/common/errorx"
	"github.com/tqtcloud/manage/service/task/rpc/types/task"
	"strconv"

	"github.com/tqtcloud/manage/service/front-api/internal/svc"
	"github.com/tqtcloud/manage/service/front-api/internal/types"

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

func (l *DeleteTaskLogic) DeleteTask(req *types.DeleteTaskRequest) (resp *types.DeleteTaskResponse, err error) {
	res, err := l.svcCtx.TaskRpc.TaskDelete(l.ctx, &task.DeleteRequest{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	secretId, err := strconv.ParseInt(res.SecretId, 10, 64)
	if err != nil {
		return nil, errorx.NewCodeError(1010, "int64 类型转换错误")
	}
	return &types.DeleteTaskResponse{
		Id:           res.Id,
		TaskName:     res.TaskName,
		Vendor:       string(res.Vendor),
		TaskType:     string(res.TaskType),
		SecretId:     secretId,
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
