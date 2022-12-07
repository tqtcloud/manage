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

func (l *TaskGetIdLogic) TaskGetId(in *task.GetIdRequest) (*task.DeleteResponse, error) {
	resp, err := l.svcCtx.TaskModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, errorx.NewCodeError(1011, "ID 不存在,重新输入")
		}
		return nil, errorx.NewCodeError(1011, err.Error())
	}

	l.Infof("厂商：%s, 类型：%s", resp.Vendor, resp.Tasktype)
	return &task.DeleteResponse{
		Id:           resp.Id,
		TaskName:     resp.Taskname,
		Vendor:       resp.Vendor,
		TaskType:     resp.Tasktype,
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
