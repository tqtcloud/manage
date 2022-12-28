package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ DomainModel = (*customDomainModel)(nil)

type (
	// DomainModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDomainModel.
	DomainModel interface {
		domainModel
	}

	customDomainModel struct {
		*defaultDomainModel
	}
)

// NewDomainModel returns a model for the database table.
func NewDomainModel(conn sqlx.SqlConn, c cache.CacheConf) DomainModel {
	return &customDomainModel{
		defaultDomainModel: newDomainModel(conn, c),
	}
}
