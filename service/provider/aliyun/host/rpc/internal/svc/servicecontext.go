package svc

import (
	"github.com/tqtcloud/manage/service/provider/aliyun/host/rpc/internal/config"
	"github.com/tqtcloud/manage/service/provider/aliyun/operator/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config      config.Config
	HostsModel  model.HostsModel
	DomainModel model.DomainModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:      c,
		HostsModel:  model.NewHostsModel(conn, c.CacheRedis),
		DomainModel: model.NewDomainModel(conn, c.CacheRedis),
	}
}
