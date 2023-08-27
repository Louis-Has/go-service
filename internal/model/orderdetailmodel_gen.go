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
	orderDetailFieldNames          = builder.RawFieldNames(&OrderDetail{})
	orderDetailRows                = strings.Join(orderDetailFieldNames, ",")
	orderDetailRowsExpectAutoSet   = strings.Join(stringx.Remove(orderDetailFieldNames, "`id`", "`created_at`", "`updated_at`"), ",")
	orderDetailRowsWithPlaceHolder = strings.Join(stringx.Remove(orderDetailFieldNames, "`id`", "`created_at`", "`updated_at`"), "=?,") + "=?"

	cacheTestDBOrderDetailIdPrefix = "cache:testDB:orderDetail:id:"
)

type (
	orderDetailModel interface {
		Insert(ctx context.Context, data *OrderDetail) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*OrderDetail, error)
		Update(ctx context.Context, data *OrderDetail) error
		Delete(ctx context.Context, id int64) error
	}

	defaultOrderDetailModel struct {
		sqlc.CachedConn
		table string
	}

	OrderDetail struct {
		Id               int64        `db:"id"`                 // 订单明细id
		CreatedAt        time.Time    `db:"created_at"`         // 创建时间
		UpdatedAt        time.Time    `db:"updated_at"`         // 更新时间
		DeletedAt        sql.NullTime `db:"deleted_at"`         // 删除时间
		OrderId          int64        `db:"order_id"`           // 订单id
		UserId           string       `db:"user_id"`            // 用户id
		ProductId        int64        `db:"product_id"`         // 商品id
		CurrentUnitPrice float64      `db:"current_unit_price"` // 生成订单时的商品单价，单位是元,保留两位小数
		ProductQuantity  int64        `db:"product_quantity"`   // 商品数量
		TotalPrice       float64      `db:"total_price"`        // 商品总价,单位是元,保留两位小数
	}
)

func newOrderDetailModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultOrderDetailModel {
	return &defaultOrderDetailModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`order_detail`",
	}
}

func (m *defaultOrderDetailModel) withSession(session sqlx.Session) *defaultOrderDetailModel {
	return &defaultOrderDetailModel{
		CachedConn: m.CachedConn.WithSession(session),
		table:      "`order_detail`",
	}
}

func (m *defaultOrderDetailModel) Delete(ctx context.Context, id int64) error {
	testDBOrderDetailIdKey := fmt.Sprintf("%s%v", cacheTestDBOrderDetailIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, testDBOrderDetailIdKey)
	return err
}

func (m *defaultOrderDetailModel) FindOne(ctx context.Context, id int64) (*OrderDetail, error) {
	testDBOrderDetailIdKey := fmt.Sprintf("%s%v", cacheTestDBOrderDetailIdPrefix, id)
	var resp OrderDetail
	err := m.QueryRowCtx(ctx, &resp, testDBOrderDetailIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", orderDetailRows, m.table)
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

func (m *defaultOrderDetailModel) Insert(ctx context.Context, data *OrderDetail) (sql.Result, error) {
	testDBOrderDetailIdKey := fmt.Sprintf("%s%v", cacheTestDBOrderDetailIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, orderDetailRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.DeletedAt, data.OrderId, data.UserId, data.ProductId, data.CurrentUnitPrice, data.ProductQuantity, data.TotalPrice)
	}, testDBOrderDetailIdKey)
	return ret, err
}

func (m *defaultOrderDetailModel) Update(ctx context.Context, data *OrderDetail) error {
	testDBOrderDetailIdKey := fmt.Sprintf("%s%v", cacheTestDBOrderDetailIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, orderDetailRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.DeletedAt, data.OrderId, data.UserId, data.ProductId, data.CurrentUnitPrice, data.ProductQuantity, data.TotalPrice, data.Id)
	}, testDBOrderDetailIdKey)
	return err
}

func (m *defaultOrderDetailModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheTestDBOrderDetailIdPrefix, primary)
}

func (m *defaultOrderDetailModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", orderDetailRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultOrderDetailModel) tableName() string {
	return m.table
}