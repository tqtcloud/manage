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

type GetListTaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetListTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetListTaskLogic {
	return &GetListTaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetListTaskLogic) GetListTask(req *types.GetTaskListRequest) (resp []*types.GetTaskListResponse, err error) {
	res, err := l.svcCtx.TaskRpc.TaskList(l.ctx, &task.GetListRequest{
		Page:  req.Page,
		Limit: req.Limit,
	})

	if err != nil {
		return nil, err
	}

	taskList := make([]*types.GetTaskListResponse, 0)

	for _, item := range res.Data {
		secretId, err := strconv.ParseInt(item.SecretId, 10, 64)
		if err != nil {
			return nil, errorx.NewCodeError(1010, "int64 类型转换错误")
		}
		taskList = append(taskList, &types.GetTaskListResponse{
			Id:           item.Id,
			TaskName:     item.TaskName,
			Vendor:       string(item.Vendor),
			TaskType:     string(item.TaskType),
			SecretId:     secretId,
			Region:       item.Region,
			TaskUser:     item.TaskUser,
			Status:       item.Status,
			Message:      item.Message,
			Start_At:     item.Start_At,
			TotalSucceed: item.TotalSucceed,
			TotalFailed:  item.TotalFailed,
			CreateTime:   item.CreateTime,
			UpdateTime:   item.UpdateTime,
		})
	}

	return taskList, nil
}
