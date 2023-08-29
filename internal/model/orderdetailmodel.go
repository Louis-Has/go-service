package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ OrderDetailModel = (*customOrderDetailModel)(nil)

type (
	// OrderDetailModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOrderDetailModel.
	OrderDetailModel interface {
		orderDetailModel
		FindAllByOrderId(ctx context.Context, orderId int64) ([]*OrderDetail, error)
	}

	customOrderDetailModel struct {
		*defaultOrderDetailModel
	}
)

// NewOrderDetailModel returns a model for the database table.
func NewOrderDetailModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) OrderDetailModel {
	return &customOrderDetailModel{
		defaultOrderDetailModel: newOrderDetailModel(conn, c, opts...),
	}
}

func (c customOrderDetailModel) FindAllByOrderId(ctx context.Context, orderId int64) ([]*OrderDetail, error) {
	var res []*OrderDetail
	query := fmt.Sprintf("select %s from %s where `order_id` = ? ", orderDetailRows, c.table)
	err := c.QueryRowsNoCacheCtx(ctx, &res, query, orderId)

	switch err {
	case nil:
		return res, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
