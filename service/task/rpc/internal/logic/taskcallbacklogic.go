package logic

import (
	"context"
	"github.com/tqtcloud/manage/common/desencryption"
	"github.com/tqtcloud/manage/service/provider/aliyun/host/rpc/types/host"
	"github.com/tqtcloud/manage/service/secret/rpc/types/secret"
	"github.com/tqtcloud/manage/service/task/model"
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
					l.Logger.Errorf("task 查询 SecretGetId %s", err)
					return nil, errorx.NewDefaultError("Secret 获取错误")
				}
				sk, _ := desencryption.Decrypt(secretData.AccessKeySecret, []byte(l.svcCtx.Config.Salt))
				l.Infof("秘钥信息为：%s", sk)

				start := time.Now()
				resp, err := l.svcCtx.HostRpc.HostSync(l.ctx, &host.CreateRequest{
					AccessKeyId:     secretData.AccessKeyId,
					AccessKeySecret: sk,
					Vendor:          in.Vendor.String(),
					Region:          in.Region,
					TaskType:        in.TaskType.String(),
				})
				if err != nil {
					newtask := model.Task{
						Id:           in.TaskId,
						Status:       task.Stage_FAILED.String(),
						Message:      err.Error(),
						StartAt:      start.UnixMilli(),  // 毫秒返回
						EndAt:        time.Now().Sub(start).Milliseconds(),
						TotalSucceed: int64(len(resp.Data)),
						TotalFailed:  0,
					}
					l.svcCtx.TaskModel.Update(l.ctx,&newtask)
					return nil, errorx.NewDefaultError(err.Error())

				} else {
					newtask := model.Task{
						Id:           in.TaskId,
						Status:       task.Stage_SUCCESS.String(),
						Message:      err.Error(),
						StartAt:      start.UnixMilli(),  // 毫秒返回
						EndAt:        time.Now().Sub(start).Milliseconds(),
						TotalSucceed: int64(len(resp.Data)),
						TotalFailed:  0,
					}
					l.svcCtx.TaskModel.Update(l.ctx,&newtask)
					return nil, errorx.NewDefaultError(err.Error())
				}

		}
	}

	return &task.CallbackResponse{}, nil
}
