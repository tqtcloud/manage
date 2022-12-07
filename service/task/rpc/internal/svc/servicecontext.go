package svc

import (
	"github.com/tqtcloud/manage/service/provider/aliyun/host/rpc/hostclient"
	"github.com/tqtcloud/manage/service/provider/aliyun/host/rpc/types/host"
	"github.com/tqtcloud/manage/service/secret/rpc/secretclient"
	"github.com/tqtcloud/manage/service/secret/rpc/types/secret"
	"github.com/tqtcloud/manage/service/task/model"
	"github.com/tqtcloud/manage/service/task/rpc/internal/config"
	"github.com/tqtcloud/manage/service/user/rpc/types/user"
	"github.com/tqtcloud/manage/service/user/rpc/userclient"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config    config.Config
	TaskModel model.TaskModel
	SecretRpc secret.SecretClient
	UserRpc   user.UserClient
	HostRpc   host.HostClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:    c,
		TaskModel: model.NewTaskModel(conn, c.CacheRedis),
		SecretRpc: secretclient.NewSecret(zrpc.MustNewClient(c.SecretRpc)),
		UserRpc:   userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		HostRpc:   hostclient.NewHost(zrpc.MustNewClient(c.HostRpc)),
	}
}
