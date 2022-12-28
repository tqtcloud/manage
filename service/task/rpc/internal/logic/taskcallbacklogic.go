package logic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/tqtcloud/manage/common/desencryption"
	"github.com/tqtcloud/manage/common/xerr"
	alioperator "github.com/tqtcloud/manage/service/provider/aliyun/operator/rpc/types/operator"
	"github.com/tqtcloud/manage/service/secret/rpc/types/secret"
	"github.com/tqtcloud/manage/service/task/rpc/internal/svc"
	"github.com/tqtcloud/manage/service/task/rpc/types/task"
	"github.com/tqtcloud/resp/errorx"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type TaskCallbackLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTaskCallbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TaskCallbackLogic {
	return &TaskCallbackLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TaskCallbackLogic) TaskCallback(in *task.CallbackRequest) (*task.CallbackResponse, error) {
	l.Logger.Info("TaskCallback 开始回调执行", in.TaskId, in.SecretId)
	switch in.Vendor {
	case task.Vendor_AliYun:
		switch in.TaskType {
		case task.TaskType_HOST:
			secretData, err := l.svcCtx.SecretRpc.SecretGetId(l.ctx, &secret.GetIdRequest{Id: in.SecretId})
			if err != nil {
				return nil, errors.Wrapf(xerr.NewErrCode(xerr.SecretIDNoExistError), "task 查询 SecretGetId err:%v,Secret:%+v", err, in.SecretId)
			}
			sk, _ := desencryption.Decrypt(secretData.AccessKeySecret, []byte(l.svcCtx.Config.Salt))
			//l.Infof("秘钥信息为：%s", sk)

			start := time.Now()
			_, err = l.svcCtx.AliOperatorRpc.HostSync(context.Background(), &alioperator.CreateHostRequest{
				AccessKeyId:     secretData.AccessKeyId,
				AccessKeySecret: sk,
				Vendor:          in.Vendor.String(),
				Region:          in.Region,
				TaskType:        in.TaskType.String(),
			})
			if err != nil {
				updateTask, _ := l.svcCtx.TaskModel.FindOne(l.ctx, in.TaskId)
				updateTask.Status = task.Stage_FAILED.String()
				updateTask.Message = err.Error()
				updateTask.StartAt = start.UnixMilli()
				updateTask.EndAt = time.Now().Sub(start).Milliseconds()
				updateTask.TotalSucceed = 0
				updateTask.TotalFailed = 0

				//newtask := model.Task{
				//	Id:           in.TaskId,
				//	Status:       task.Stage_FAILED.String(),
				//	Message:      err.Error(),
				//	StartAt:      start.UnixMilli(), // 毫秒返回
				//	EndAt:        time.Now().Sub(start).Milliseconds(),
				//	TotalSucceed: 0,
				//	TotalFailed:  0,
				//}
				l.svcCtx.TaskModel.Update(l.ctx, updateTask)
				return nil, errorx.NewDefaultError(err.Error())

			} else {
				updateTask, _ := l.svcCtx.TaskModel.FindOne(l.ctx, in.TaskId)
				updateTask.Status = task.Stage_SUCCESS.String()
				updateTask.Message = "同步成功"
				updateTask.StartAt = start.UnixMilli()
				updateTask.EndAt = time.Now().Sub(start).Milliseconds()
				updateTask.TotalSucceed = 0
				updateTask.TotalFailed = 0
				//newtask := model.Task{
				//	Id:           in.TaskId,
				//	Status:       task.Stage_SUCCESS.String(),
				//	Message:      "同步成功",
				//	StartAt:      start.UnixMilli(), // 毫秒返回
				//	EndAt:        time.Now().Sub(start).Milliseconds(),
				//	TotalSucceed: 0,
				//	TotalFailed:  0,
				//}
				if err := l.svcCtx.TaskModel.Update(l.ctx, updateTask); err != nil {
					return nil, errors.Wrapf(xerr.NewErrCode(xerr.TaskUpdateError), "回调更新任务 err:%v, TaskID:%+v", err, in.TaskId)
				}

			}
		default:
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.TaskTypeError), "没有匹配的同步类别:%v, TaskID:%+v", in.TaskType.String(), in.TaskId)

		}
	default:
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.TaskVendorError), "没有匹配的同步云厂商 Vendor:%v, TaskID:%+v", in.Vendor.String(), in.TaskId)
	}

	return &task.CallbackResponse{}, nil
}
