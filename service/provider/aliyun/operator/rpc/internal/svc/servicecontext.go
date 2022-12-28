package svc

import (
	"github.com/tqtcloud/manage/service/provider/aliyun/operator/model"
	"github.com/tqtcloud/manage/service/provider/aliyun/operator/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config     config.Config
	HostsModel model.HostsModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:     c,
		HostsModel: model.NewHostsModel(conn, c.CacheRedis),
	}
}
