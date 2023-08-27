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
	productCategoryFieldNames          = builder.RawFieldNames(&ProductCategory{})
	productCategoryRows                = strings.Join(productCategoryFieldNames, ",")
	productCategoryRowsExpectAutoSet   = strings.Join(stringx.Remove(productCategoryFieldNames, "`id`", "`created_at`", "`updated_at`"), ",")
	productCategoryRowsWithPlaceHolder = strings.Join(stringx.Remove(productCategoryFieldNames, "`id`", "`created_at`", "`updated_at`"), "=?,") + "=?"

	cacheTestDBProductCategoryIdPrefix = "cache:testDB:productCategory:id:"
)

type (
	productCategoryModel interface {
		Insert(ctx context.Context, data *ProductCategory) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*ProductCategory, error)
		Update(ctx context.Context, data *ProductCategory) error
		Delete(ctx context.Context, id int64) error
	}

	defaultProductCategoryModel struct {
		sqlc.CachedConn
		table string
	}

	ProductCategory struct {
		Id        int64        `db:"id"`         // 分类id
		CreatedAt time.Time    `db:"created_at"` // 创建时间
		UpdatedAt time.Time    `db:"updated_at"` // 更新时间
		DeletedAt sql.NullTime `db:"deleted_at"` // 删除时间
		ParentId  int64        `db:"parent_id"`  // 父类别id当id=0时说明是根节点,一级类别
		Name      string       `db:"name"`       // 类别名称
	}
)

func newProductCategoryModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultProductCategoryModel {
	return &defaultProductCategoryModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`product_category`",
	}
}

func (m *defaultProductCategoryModel) withSession(session sqlx.Session) *defaultProductCategoryModel {
	return &defaultProductCategoryModel{
		CachedConn: m.CachedConn.WithSession(session),
		table:      "`product_category`",
	}
}

func (m *defaultProductCategoryModel) Delete(ctx context.Context, id int64) error {
	testDBProductCategoryIdKey := fmt.Sprintf("%s%v", cacheTestDBProductCategoryIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, testDBProductCategoryIdKey)
	return err
}

func (m *defaultProductCategoryModel) FindOne(ctx context.Context, id int64) (*ProductCategory, error) {
	testDBProductCategoryIdKey := fmt.Sprintf("%s%v", cacheTestDBProductCategoryIdPrefix, id)
	var resp ProductCategory
	err := m.QueryRowCtx(ctx, &resp, testDBProductCategoryIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", productCategoryRows, m.table)
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

func (m *defaultProductCategoryModel) Insert(ctx context.Context, data *ProductCategory) (sql.Result, error) {
	testDBProductCategoryIdKey := fmt.Sprintf("%s%v", cacheTestDBProductCategoryIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, productCategoryRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.DeletedAt, data.ParentId, data.Name)
	}, testDBProductCategoryIdKey)
	return ret, err
}

func (m *defaultProductCategoryModel) Update(ctx context.Context, data *ProductCategory) error {
	testDBProductCategoryIdKey := fmt.Sprintf("%s%v", cacheTestDBProductCategoryIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, productCategoryRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.DeletedAt, data.ParentId, data.Name, data.Id)
	}, testDBProductCategoryIdKey)
	return err
}

func (m *defaultProductCategoryModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheTestDBProductCategoryIdPrefix, primary)
}

func (m *defaultProductCategoryModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", productCategoryRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultProductCategoryModel) tableName() string {
	return m.table
}