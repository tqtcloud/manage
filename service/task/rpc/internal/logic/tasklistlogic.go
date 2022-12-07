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

type TaskListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTaskListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TaskListLogic {
	return &TaskListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TaskListLogic) TaskList(in *task.GetListRequest) (*task.GetListResponse, error) {
	list, err := l.svcCtx.TaskModel.FindAllPage(l.ctx, in)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, errorx.NewDefaultError(err.Error())
		}
		return nil, errorx.NewDefaultError(err.Error())
	}

	taskList := make([]*task.DeleteResponse, 0)
	for _, item := range list {
		taskList = append(taskList, &task.DeleteResponse{
			Id:           item.Id,
			TaskName:     item.Taskname,
			Vendor:       item.Vendor,
			TaskType:     item.Tasktype,
			SecretId:     strconv.FormatInt(item.SecretId, 10),
			Region:       item.Region,
			TaskUser:     item.Taskuser,
			Status:       item.Status,
			Message:      item.Message,
			Start_At:     strconv.FormatInt(item.StartAt, 10),
			TotalSucceed: item.TotalSucceed,
			TotalFailed:  item.TotalFailed,
			CreateTime:   item.CreateTime.Format("2006-01-02 15:04:05"),
			UpdateTime:   item.UpdateTime.Format("2006-01-02 15:04:05"),
		})
	}

	return &task.GetListResponse{
		Data: taskList,
	}, nil
}
