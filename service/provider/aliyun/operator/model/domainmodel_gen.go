// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	domainFieldNames          = builder.RawFieldNames(&Domain{})
	domainRows                = strings.Join(domainFieldNames, ",")
	domainRowsExpectAutoSet   = strings.Join(stringx.Remove(domainFieldNames, "`update_at`", "`updated_at`", "`update_time`", "`create_at`", "`created_at`", "`create_time`"), ",")
	domainRowsWithPlaceHolder = strings.Join(stringx.Remove(domainFieldNames, "`domain_name`", "`update_at`", "`updated_at`", "`update_time`", "`create_at`", "`created_at`", "`create_time`"), "=?,") + "=?"

	cacheDomainDomainNamePrefix = "cache:domain:domainName:"
	cacheDomainInstanceIdPrefix = "cache:domain:instanceId:"
)

type (
	domainModel interface {
		Insert(ctx context.Context, data *Domain) (sql.Result, error)
		FindOne(ctx context.Context, domainName string) (*Domain, error)
		FindOneByInstanceId(ctx context.Context, instanceId string) (*Domain, error)
		Update(ctx context.Context, data *Domain) error
		Delete(ctx context.Context, domainName string) error
	}

	defaultDomainModel struct {
		sqlc.CachedConn
		table string
	}

	Domain struct {
		InstanceId             string `db:"instance_id"`               // 实例Id
		DomainName             string `db:"domain_name"`               // 域名名称
		Belong                 string `db:"belong"`                    // 所属业务线
		DomainAuditStatus      string `db:"domain_audit_status"`       // 域名实名认证状态 FAILED：实名认证失败 SUCCEED：实名认证成功 NONAUDIT：未实名认证 AUDITING：审核中
		RegistrationTime       string `db:"registration_time"`         // 注册时间
		ExpirationDateStatus   string `db:"expiration_date_status"`    // 域名过期状态: 1：域名未过期 2：域名已过期
		ExpiredTime            string `db:"expired_time"`              // 域名到期日期
		ExpirationCurrDateDiff string `db:"expiration_curr_date_diff"` // 域名到期日和当前的时间的天数差值
		RegistrantType         string `db:"registrant_type"`           // 域名注册类型: 1：个人 2：企业
		DomainStatus           string `db:"domain_status"`             // 域名状态: 1：急需续费 2：急需赎回 3：正常
	}
)

func newDomainModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultDomainModel {
	return &defaultDomainModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`domain`",
	}
}

func (m *defaultDomainModel) Delete(ctx context.Context, domainName string) error {
	data, err := m.FindOne(ctx, domainName)
	if err != nil {
		return err
	}

	domainDomainNameKey := fmt.Sprintf("%s%v", cacheDomainDomainNamePrefix, domainName)
	domainInstanceIdKey := fmt.Sprintf("%s%v", cacheDomainInstanceIdPrefix, data.InstanceId)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `domain_name` = ?", m.table)
		return conn.ExecCtx(ctx, query, domainName)
	}, domainDomainNameKey, domainInstanceIdKey)
	return err
}

func (m *defaultDomainModel) FindOne(ctx context.Context, domainName string) (*Domain, error) {
	domainDomainNameKey := fmt.Sprintf("%s%v", cacheDomainDomainNamePrefix, domainName)
	var resp Domain
	err := m.QueryRowCtx(ctx, &resp, domainDomainNameKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `domain_name` = ? limit 1", domainRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, domainName)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultDomainModel) FindOneByInstanceId(ctx context.Context, instanceId string) (*Domain, error) {
	domainInstanceIdKey := fmt.Sprintf("%s%v", cacheDomainInstanceIdPrefix, instanceId)
	var resp Domain
	err := m.QueryRowIndexCtx(ctx, &resp, domainInstanceIdKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `instance_id` = ? limit 1", domainRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, instanceId); err != nil {
			return nil, err
		}
		return resp.DomainName, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultDomainModel) Insert(ctx context.Context, data *Domain) (sql.Result, error) {
	domainDomainNameKey := fmt.Sprintf("%s%v", cacheDomainDomainNamePrefix, data.DomainName)
	domainInstanceIdKey := fmt.Sprintf("%s%v", cacheDomainInstanceIdPrefix, data.InstanceId)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, domainRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.InstanceId, data.DomainName, data.Belong, data.DomainAuditStatus, data.RegistrationTime, data.ExpirationDateStatus, data.ExpiredTime, data.ExpirationCurrDateDiff, data.RegistrantType, data.DomainStatus)
	}, domainDomainNameKey, domainInstanceIdKey)
	return ret, err
}

func (m *defaultDomainModel) Update(ctx context.Context, newData *Domain) error {
	data, err := m.FindOne(ctx, newData.DomainName)
	if err != nil {
		return err
	}

	domainDomainNameKey := fmt.Sprintf("%s%v", cacheDomainDomainNamePrefix, data.DomainName)
	domainInstanceIdKey := fmt.Sprintf("%s%v", cacheDomainInstanceIdPrefix, data.InstanceId)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `domain_name` = ?", m.table, domainRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.InstanceId, newData.Belong, newData.DomainAuditStatus, newData.RegistrationTime, newData.ExpirationDateStatus, newData.ExpiredTime, newData.ExpirationCurrDateDiff, newData.RegistrantType, newData.DomainStatus, newData.DomainName)
	}, domainDomainNameKey, domainInstanceIdKey)
	return err
}

func (m *defaultDomainModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheDomainDomainNamePrefix, primary)
}

func (m *defaultDomainModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `domain_name` = ? limit 1", domainRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultDomainModel) tableName() string {
	return m.table
}
