package model

import (
	"context"
	"fmt"
	"github.com/tqtcloud/manage/service/user/rpc/types/user"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserModel = (*customUserModel)(nil)

type (
	// UserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserModel.
	UserModel interface {
		userModel
		FindAllPage(ctx context.Context, req *user.UserListRequest) ([]*User, error)
	}

	customUserModel struct {
		*defaultUserModel
	}
)

// NewUserModel returns a model for the database table.
func NewUserModel(conn sqlx.SqlConn, c cache.CacheConf) UserModel {
	return &customUserModel{
		defaultUserModel: newUserModel(conn, c),
	}
}

func (m *defaultUserModel) FindAllPage(ctx context.Context, req *user.UserListRequest) ([]*User, error) {

	var resp []*User

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
