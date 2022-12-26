package svc

import (
	"github.com/tqtcloud/manage/service/front-api/internal/config"
	"github.com/tqtcloud/manage/service/secret/rpc/secretclient"
	"github.com/tqtcloud/manage/service/secret/rpc/types/secret"
	"github.com/tqtcloud/manage/service/task/rpc/taskclient"
	"github.com/tqtcloud/manage/service/task/rpc/types/task"
	"github.com/tqtcloud/manage/service/user/rpc/types/user"
	"github.com/tqtcloud/manage/service/user/rpc/userclient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config    config.Config
	UserRpc   user.UserClient
	TaskRpc   task.TaskClient
	SecretRpc secret.SecretClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		UserRpc:   userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		TaskRpc:   taskclient.NewTask(zrpc.MustNewClient(c.TaskRpc)),
		SecretRpc: secretclient.NewSecret(zrpc.MustNewClient(c.SecretRpc)),
	}
}
