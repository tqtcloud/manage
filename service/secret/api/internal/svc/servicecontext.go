package svc

import (
	"github.com/tqtcloud/manage/service/secret/api/internal/config"
	"github.com/tqtcloud/manage/service/secret/rpc/secretclient"
	"github.com/tqtcloud/manage/service/secret/rpc/types/secret"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config    config.Config
	SecretRpc secret.SecretClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		SecretRpc: secretclient.NewSecret(zrpc.MustNewClient(c.SecretRpc)),
	}
}
