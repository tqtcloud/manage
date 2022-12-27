package logic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/tqtcloud/manage/common/xerr"
	"github.com/tqtcloud/manage/service/task/model"
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
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.TaskIDNoExistError), "task 查询 TaskID err:%v,Task:%+v", err, in.Id)
		}
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "其他错误 err:%v", err)
	}

	err = l.svcCtx.TaskModel.Delete(l.ctx, in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.TaskIDNoExistError), "Task ErrNotFound err:%v", err)
		}
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "其他错误 err:%v", err)
	}

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
