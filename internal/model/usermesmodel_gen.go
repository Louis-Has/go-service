// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	userMesFieldNames          = builder.RawFieldNames(&UserMes{})
	userMesRows                = strings.Join(userMesFieldNames, ",")
	userMesRowsExpectAutoSet   = strings.Join(stringx.Remove(userMesFieldNames, "`id`", "`created_at`", "`updated_at`"), ",")
	userMesRowsWithPlaceHolder = strings.Join(stringx.Remove(userMesFieldNames, "`id`", "`created_at`", "`updated_at`"), "=?,") + "=?"

	cacheTestDBUserMesIdPrefix = "cache:testDB:userMes:id:"
)

type (
	userMesModel interface {
		Insert(ctx context.Context, data *UserMes) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*UserMes, error)
		Update(ctx context.Context, data *UserMes) error
		Delete(ctx context.Context, id int64) error
	}

	defaultUserMesModel struct {
		sqlc.CachedConn
		table string
	}

	UserMes struct {
		Id           int64        `db:"id"`            // 用户ID
		CreatedAt    time.Time    `db:"created_at"`    // 创建时间
		UpdatedAt    time.Time    `db:"updated_at"`    // 更新时间
		DeletedAt    sql.NullTime `db:"deleted_at"`    // 删除时间
		UserName     string       `db:"user_name"`     // 用户名
		SignedPerson int64        `db:"signed_person"` // 标记用户，0表示默认；1表示正式用户；2表示VIP
	}
)

func newUserMesModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultUserMesModel {
	return &defaultUserMesModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`user_mes`",
	}
}

func (m *defaultUserMesModel) withSession(session sqlx.Session) *defaultUserMesModel {
	return &defaultUserMesModel{
		CachedConn: m.CachedConn.WithSession(session),
		table:      "`user_mes`",
	}
}

func (m *defaultUserMesModel) Delete(ctx context.Context, id int64) error {
	testDBUserMesIdKey := fmt.Sprintf("%s%v", cacheTestDBUserMesIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, testDBUserMesIdKey)
	return err
}

func (m *defaultUserMesModel) FindOne(ctx context.Context, id int64) (*UserMes, error) {
	testDBUserMesIdKey := fmt.Sprintf("%s%v", cacheTestDBUserMesIdPrefix, id)
	var resp UserMes
	err := m.QueryRowCtx(ctx, &resp, testDBUserMesIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userMesRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
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

func (m *defaultUserMesModel) Insert(ctx context.Context, data *UserMes) (sql.Result, error) {
	testDBUserMesIdKey := fmt.Sprintf("%s%v", cacheTestDBUserMesIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, userMesRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.DeletedAt, data.UserName, data.SignedPerson)
	}, testDBUserMesIdKey)
	return ret, err
}

func (m *defaultUserMesModel) Update(ctx context.Context, data *UserMes) error {
	testDBUserMesIdKey := fmt.Sprintf("%s%v", cacheTestDBUserMesIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, userMesRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.DeletedAt, data.UserName, data.SignedPerson, data.Id)
	}, testDBUserMesIdKey)
	return err
}

func (m *defaultUserMesModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheTestDBUserMesIdPrefix, primary)
}

func (m *defaultUserMesModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userMesRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultUserMesModel) tableName() string {
	return m.table
}
