package logic

import (
	"context"
	"github.com/tqtcloud/manage/common/desencryption"
	"github.com/tqtcloud/manage/service/secret/rpc/types/secret"
	"github.com/tqtcloud/manage/service/task/model"
	"github.com/tqtcloud/manage/service/user/rpc/types/user"
	"github.com/tqtcloud/resp/errorx"
	"strconv"
	"time"

	"github.com/tqtcloud/manage/service/task/rpc/internal/svc"
	"github.com/tqtcloud/manage/service/task/rpc/types/task"

	"github.com/zeromicro/go-zero/core/logx"
)

type TaskCreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTaskCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TaskCreateLogic {
	return &TaskCreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TaskCreateLogic) TaskCreate(in *task.CreateRequest) (*task.CreateResponse, error) {
	_, err := l.svcCtx.TaskModel.FindOneByTaskname(l.ctx, in.TaskName)
	if err == nil {
		return nil, errorx.NewDefaultError("TaskName已存在,请重新输入")
	}

	secretData, err := l.svcCtx.SecretRpc.SecretGetId(l.ctx, &secret.GetIdRequest{Id: in.SecretId})
	if err != nil {
		return nil, err
	}
	userid := l.ctx.Value("uid")
	userinfo, _ := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoRequest{Id: userid.(int64)})

	_, _ = desencryption.Decrypt(secretData.AccessKeySecret, []byte(l.svcCtx.Config.Salt))
	if err == model.ErrNotFound {
		newTask := model.Task{
			Taskname:     in.TaskName,
			Vendor:       in.Vendor,
			Tasktype:     in.TaskType,
			SecretId:     in.SecretId,
			SecretDesc:   secretData.Vendor,
			Region:       in.Region,
			Taskuser:     userinfo.Name,
			Status:       task.Stage_RUNNING.String(),
			Message:      "任务运行中",
			StartAt:      time.Now().Unix(),
			EndAt:        0,
			TotalSucceed: 0,
			TotalFailed:  0,
		}

		resp, err := l.svcCtx.TaskModel.Insert(l.ctx, &newTask)
		if err != nil {
			l.Logger.Infof("创建Task信息：%s", newTask)
			l.Logger.Errorf("任务创建错误: %s", err)
			return nil, errorx.NewUserError("任务创建错误,请联系管理员")
		}

		newTask.Id, err = resp.LastInsertId()
		if err != nil {
			l.Logger.Errorf("数据库递增错误: %s", err)
			return nil, errorx.NewUserError("内部错误请联系,请联系管理员")
		}
		return &task.CreateResponse{
			TaskName: newTask.Taskname,
			Vendor:   newTask.Vendor,
			TaskType: newTask.Tasktype,
			SecretId: strconv.FormatInt(newTask.SecretId, 10),
			Region:   newTask.Region,
			TaskUser: userinfo.Name,
			Status:   newTask.Status,
			Message:  newTask.Message,
			Start_At: strconv.FormatInt(newTask.StartAt, 10),
		}, nil
	}

	return nil, errorx.NewDefaultError("内部错误请联系,请联系管理员")
}
