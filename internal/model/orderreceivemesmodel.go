package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ OrderReceiveMesModel = (*customOrderReceiveMesModel)(nil)

type (
	// OrderReceiveMesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOrderReceiveMesModel.
	OrderReceiveMesModel interface {
		orderReceiveMesModel
	}

	customOrderReceiveMesModel struct {
		*defaultOrderReceiveMesModel
	}
)

// NewOrderReceiveMesModel returns a model for the database table.
func NewOrderReceiveMesModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) OrderReceiveMesModel {
	return &customOrderReceiveMesModel{
		defaultOrderReceiveMesModel: newOrderReceiveMesModel(conn, c, opts...),
	}
}
