package model

import (
	"context"
	"fmt"
	"github.com/tqtcloud/manage/service/secret/rpc/types/secret"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SecretModel = (*customSecretModel)(nil)

type (
	// SecretModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSecretModel.
	SecretModel interface {
		secretModel
		FindAllPage(ctx context.Context, req *secret.GetListRequest) ([]*Secret, error)
	}

	customSecretModel struct {
		*defaultSecretModel
	}
)

// NewSecretModel returns a model for the database table.
func NewSecretModel(conn sqlx.SqlConn, c cache.CacheConf) SecretModel {
	return &customSecretModel{
		defaultSecretModel: newSecretModel(conn, c),
	}
}

func (m *defaultSecretModel) FindAllPage(ctx context.Context, req *secret.GetListRequest) ([]*Secret, error) {
	var resp []*Secret

	query := fmt.Sprintf("select * from %s ORDER BY `id` ASC  LIMIT ? OFFSET ?", m.table)
	fmt.Println("log：", query)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, req.Limit, req.Page-1)

	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
