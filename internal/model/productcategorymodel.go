package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ProductCategoryModel = (*customProductCategoryModel)(nil)

type (
	// ProductCategoryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customProductCategoryModel.
	ProductCategoryModel interface {
		productCategoryModel
		FindAll(ctx context.Context) ([]*ProductCategory, error)
	}

	customProductCategoryModel struct {
		*defaultProductCategoryModel
	}
)

// NewProductCategoryModel returns a model for the database table.
func NewProductCategoryModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) ProductCategoryModel {
	return &customProductCategoryModel{
		defaultProductCategoryModel: newProductCategoryModel(conn, c, opts...),
	}
}

func (c customProductCategoryModel) FindAll(ctx context.Context) ([]*ProductCategory, error) {
	var resp []*ProductCategory
	query := fmt.Sprintf("select %s from %s ", productCategoryRows, c.table)
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
