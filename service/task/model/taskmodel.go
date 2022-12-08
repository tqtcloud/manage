package model

import (
	"context"
	"fmt"
	"github.com/tqtcloud/manage/service/task/rpc/types/task"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TaskModel = (*customTaskModel)(nil)

type (
	// TaskModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTaskModel.
	TaskModel interface {
		taskModel
		FindAllPage(ctx context.Context, req *task.GetListRequest) ([]*Task, error)
	}

	customTaskModel struct {
		*defaultTaskModel
	}
)

// NewTaskModel returns a model for the database table.
func NewTaskModel(conn sqlx.SqlConn, c cache.CacheConf) TaskModel {
	return &customTaskModel{
		defaultTaskModel: newTaskModel(conn, c),
	}
}

func (m *defaultTaskModel) FindAllPage(ctx context.Context, req *task.GetListRequest) ([]*Task, error) {
	var resp []*Task

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
