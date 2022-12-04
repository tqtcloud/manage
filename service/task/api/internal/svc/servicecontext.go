package svc

import (
	"github.com/tqtcloud/manage/service/task/api/internal/config"
	"github.com/tqtcloud/manage/service/task/rpc/taskclient"
	"github.com/tqtcloud/manage/service/task/rpc/types/task"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	TaskRpc task.TaskClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		TaskRpc: taskclient.NewTask(zrpc.MustNewClient(c.TaskRpc)),
	}
}
