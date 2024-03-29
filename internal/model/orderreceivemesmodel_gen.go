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
	orderReceiveMesFieldNames          = builder.RawFieldNames(&OrderReceiveMes{})
	orderReceiveMesRows                = strings.Join(orderReceiveMesFieldNames, ",")
	orderReceiveMesRowsExpectAutoSet   = strings.Join(stringx.Remove(orderReceiveMesFieldNames, "`id`", "`created_at`", "`updated_at`"), ",")
	orderReceiveMesRowsWithPlaceHolder = strings.Join(stringx.Remove(orderReceiveMesFieldNames, "`id`", "`created_at`", "`updated_at`"), "=?,") + "=?"

	cacheTestDBOrderReceiveMesIdPrefix = "cache:testDB:orderReceiveMes:id:"
)

type (
	orderReceiveMesModel interface {
		Insert(ctx context.Context, data *OrderReceiveMes) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*OrderReceiveMes, error)
		Update(ctx context.Context, data *OrderReceiveMes) error
		Delete(ctx context.Context, id int64) error
	}

	defaultOrderReceiveMesModel struct {
		sqlc.CachedConn
		table string
	}

	OrderReceiveMes struct {
		Id               int64        `db:"id"`                // 收货信息表id
		CreatedAt        time.Time    `db:"created_at"`        // 创建时间
		UpdatedAt        time.Time    `db:"updated_at"`        // 更新时间
		DeletedAt        sql.NullTime `db:"deleted_at"`        // 删除时间
		UserId           int64        `db:"user_id"`           // 用户id
		ReceiverName     string       `db:"receiver_name"`     // 收货姓名
		ReceiverPhone    string       `db:"receiver_phone"`    // 收货固定电话
		ReceiverProvince string       `db:"receiver_province"` // 省份
		ReceiverCity     string       `db:"receiver_city"`     // 城市
		ReceiverDistrict string       `db:"receiver_district"` // 区/县
		ReceiverAddress  string       `db:"receiver_address"`  // 详细地址
	}
)

func newOrderReceiveMesModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultOrderReceiveMesModel {
	return &defaultOrderReceiveMesModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`order_receive_mes`",
	}
}

func (m *defaultOrderReceiveMesModel) withSession(session sqlx.Session) *defaultOrderReceiveMesModel {
	return &defaultOrderReceiveMesModel{
		CachedConn: m.CachedConn.WithSession(session),
		table:      "`order_receive_mes`",
	}
}

func (m *defaultOrderReceiveMesModel) Delete(ctx context.Context, id int64) error {
	testDBOrderReceiveMesIdKey := fmt.Sprintf("%s%v", cacheTestDBOrderReceiveMesIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, testDBOrderReceiveMesIdKey)
	return err
}

func (m *defaultOrderReceiveMesModel) FindOne(ctx context.Context, id int64) (*OrderReceiveMes, error) {
	testDBOrderReceiveMesIdKey := fmt.Sprintf("%s%v", cacheTestDBOrderReceiveMesIdPrefix, id)
	var resp OrderReceiveMes
	err := m.QueryRowCtx(ctx, &resp, testDBOrderReceiveMesIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", orderReceiveMesRows, m.table)
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

func (m *defaultOrderReceiveMesModel) Insert(ctx context.Context, data *OrderReceiveMes) (sql.Result, error) {
	testDBOrderReceiveMesIdKey := fmt.Sprintf("%s%v", cacheTestDBOrderReceiveMesIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?)", m.table, orderReceiveMesRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.DeletedAt, data.UserId, data.ReceiverName, data.ReceiverPhone, data.ReceiverProvince, data.ReceiverCity, data.ReceiverDistrict, data.ReceiverAddress)
	}, testDBOrderReceiveMesIdKey)
	return ret, err
}

func (m *defaultOrderReceiveMesModel) Update(ctx context.Context, data *OrderReceiveMes) error {
	testDBOrderReceiveMesIdKey := fmt.Sprintf("%s%v", cacheTestDBOrderReceiveMesIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, orderReceiveMesRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.DeletedAt, data.UserId, data.ReceiverName, data.ReceiverPhone, data.ReceiverProvince, data.ReceiverCity, data.ReceiverDistrict, data.ReceiverAddress, data.Id)
	}, testDBOrderReceiveMesIdKey)
	return err
}

func (m *defaultOrderReceiveMesModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheTestDBOrderReceiveMesIdPrefix, primary)
}

func (m *defaultOrderReceiveMesModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", orderReceiveMesRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultOrderReceiveMesModel) tableName() string {
	return m.table
}
