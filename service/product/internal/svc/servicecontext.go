package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-service/internal/model"
	"go-service/service/product/internal/config"
)

type ServiceContext struct {
	Config config.Config
	model.ProductModel
	model.ProductCategoryModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := sqlx.NewSqlConn("mysql", c.Mysql.DataSource)
	return &ServiceContext{
		Config:               c,
		ProductModel:         model.NewProductModel(db, c.CacheRedis),
		ProductCategoryModel: model.NewProductCategoryModel(db, c.CacheRedis),
	}
}
