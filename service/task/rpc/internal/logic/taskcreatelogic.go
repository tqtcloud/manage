package logic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/rs/xid"
	"github.com/tqtcloud/manage/common/xerr"
	"github.com/tqtcloud/manage/service/secret/rpc/types/secret"
	"github.com/tqtcloud/manage/service/task/model"
	"github.com/tqtcloud/manage/service/user/rpc/types/user"
	"strconv"
	"time"

	"github.com/tqtcloud/manage/service/task/rpc/internal/svc"
	"github.com/tqtcloud/manage/service/task/rpc/types/task"

	"github.com/zeromicro/go-zero/core/logx"
)

//var ErrUsernamePwdError = xerr.NewErrMsg("任务重复")

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
	secretData, err := l.svcCtx.SecretRpc.SecretGetId(l.ctx, &secret.GetIdRequest{Id: in.SecretId})
	if err != nil {
		l.Logger.Errorf("task 查询 SecretGetId %s", err)
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.SecretIDNoExistError), "task 查询 SecretGetId err:%v,Secret:%+v", err, in.SecretId)
	}
	//sk, _ := desencryption.Decrypt(secretData.AccessKeySecret, []byte(l.svcCtx.Config.Salt))
	//l.Infof("秘钥信息为：%s", sk)

	//uid, _ := l.ctx.Value("uid").(string)

	l.Logger.Infof("任务创建用户ID为：%s", in.UserId)
	userinfo, _ := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoRequest{Id: in.UserId})

	_, err = l.svcCtx.TaskModel.FindOneByTaskname(l.ctx, in.TaskName)
	if err == nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.TaskIDExistError), "查询 FindOneByTaskname err:%v,Secret:%+v", err, in.TaskName)
	}
	// 如果没有数据则创建任务
	if err == model.ErrNotFound {
		newTask := model.Task{
			Id:           xid.New().String(),
			Taskname:     in.TaskName,
			Vendor:       in.Vendor.String(),
			Tasktype:     in.TaskType.String(),
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
		l.Logger.Infof("待创建Task名称：%s,云厂商：%s, 秘钥ID: %s, 地域: %s,", newTask.Taskname, newTask.Vendor, newTask.SecretId, newTask.Region)

		resp, err := l.svcCtx.TaskModel.Insert(l.ctx, &newTask)
		if err != nil {
			l.Logger.Errorf("任务创建错误: %s", err)
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.TaskDbInsertError), "任务创建错误 err:%v,task:%+v", err, in.TaskName)
		}

		rows, err := resp.RowsAffected()
		if err != nil {
			l.Logger.Errorf("数据库递增错误: %s", err)
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "数据库递增错误 err:%v,task:%+v", err, rows)
		}

		// 回调处理任务运行的状态同步等
		//callbackClient := taskclient.NewTask(zrpc.MustNewClient(zrpc.RpcClientConf{
		//	Etcd: discov.EtcdConf{
		//		Hosts: []string{"192.168.0.102:2379"},
		//		Key:   "task.rpc",
		//	},
		//	App:   "taskapi",
		//	Token: "6jKNZbEpYGeUMAifz10gOnmoty3TV",
		//}))

		go func() {
			ctx1, cel := context.WithTimeout(context.Background(), time.Second*300)
			defer cel()
			callbackClient := NewTaskCallbackLogic(ctx1, l.svcCtx)
			defer func() {
				if err := recover(); err != nil {
					l.Errorf("panic error:%", err)
				}
			}()
			_, err := callbackClient.TaskCallback(&task.CallbackRequest{
				TaskId:   newTask.Id,
				SecretId: newTask.SecretId,
				Vendor:   in.Vendor,
				Region:   newTask.Region,
				TaskType: in.TaskType,
			})
			if err != nil {
				l.Logger.Errorf("Task 回调出现错误请注意: %s", err)
			}
		}()

		return &task.CreateResponse{
			Id:       newTask.Id,
			TaskName: newTask.Taskname,
			Vendor:   in.Vendor.String(),
			TaskType: in.TaskType.String(),
			SecretId: newTask.SecretId,
			Region:   newTask.Region,
			TaskUser: userinfo.Name,
			Status:   newTask.Status,
			Message:  newTask.Message,
			Start_At: strconv.FormatInt(newTask.StartAt, 10),
		}, nil
	}

	return nil, xerr.NewErrMsg("默认错误请联系管理员")
}
