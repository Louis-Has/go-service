package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ProductModel = (*customProductModel)(nil)

type (
	// ProductModel is an interface to be customized, add more methods here,
	// and implement the added methods in customProductModel.
	ProductModel interface {
		productModel
		FindAll(ctx context.Context) ([]*Product, error)
	}

	customProductModel struct {
		*defaultProductModel
	}
)

// NewProductModel returns a model for the database table.
func NewProductModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) ProductModel {
	return &customProductModel{
		defaultProductModel: newProductModel(conn, c, opts...),
	}
}

func (c customProductModel) FindAll(ctx context.Context) ([]*Product, error) {
	var resp []*Product
	query := fmt.Sprintf("select %s from %s ", productRows, c.table)
	err := c.QueryRowsNoCacheCtx(ctx, &resp, query)

	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
