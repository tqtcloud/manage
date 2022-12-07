package logic

import (
	"context"
	"github.com/tqtcloud/manage/service/task/model"
	"github.com/tqtcloud/resp/errorx"
	"strconv"

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
	// 查询需要删除的 ID 数据
	resp, err := l.svcCtx.TaskModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, errorx.NewCodeError(1010, "Task ID 不存在")
		}
		return nil, errorx.NewCodeError(1011, err.Error())
	}

	err = l.svcCtx.TaskModel.Delete(l.ctx, in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, errorx.NewCodeError(1010, "Task ID 不存在")
		}
		return nil, errorx.NewCodeError(1011, err.Error())
	}

	return &task.DeleteResponse{
		Id:           resp.Id,
		TaskName:     resp.Taskname,
		Vendor: task.Vendor(task.Stage_value[resp.Vendor]),
		TaskType:     task.TaskType(task.TaskType_value[resp.Tasktype]),
		SecretId:     strconv.FormatInt(resp.SecretId, 10),
		Region:       resp.Region,
		TaskUser:     resp.Taskuser,
		Status:       resp.Status,
		Message:      resp.Message,
		Start_At:     strconv.FormatInt(resp.StartAt, 10),
		TotalSucceed: resp.TotalSucceed,
		TotalFailed:  resp.TotalFailed,
		CreateTime:   resp.CreateTime.Format("2006-01-02 15:04:05"),
		UpdateTime:   resp.UpdateTime.Format("2006-01-02 15:04:05"),
	}, nil
}
