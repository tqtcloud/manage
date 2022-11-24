package svc

import (
	"github.com/tqtcloud/manage/service/user/api/internal/config"
	"github.com/tqtcloud/manage/service/user/rpc/types/user"
	"github.com/tqtcloud/manage/service/user/rpc/userclient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc user.UserClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
