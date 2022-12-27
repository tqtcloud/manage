package model

import (
	"context"
	"fmt"
	"github.com/tqtcloud/manage/service/task/rpc/types/task"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ HostsModel = (*customHostsModel)(nil)

type (
	// HostsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customHostsModel.
	HostsModel interface {
		hostsModel
		FindAllPage(ctx context.Context, req *task.GetListRequest) ([]*Hosts, error)
	}

	customHostsModel struct {
		*defaultHostsModel
	}
)

// NewHostsModel returns a model for the database table.
func NewHostsModel(conn sqlx.SqlConn, c cache.CacheConf) HostsModel {
	return &customHostsModel{
		defaultHostsModel: newHostsModel(conn, c),
	}
}

func (m *defaultHostsModel) FindAllPage(ctx context.Context, req *task.GetListRequest) ([]*Hosts, error) {
	var resp []*Hosts

	query := fmt.Sprintf("select * from %s ORDER BY `id` ASC  LIMIT ? OFFSET ?", m.table)
	fmt.Println("log：", query)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, req.Limit, (req.Page-1)*req.Limit)

	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
