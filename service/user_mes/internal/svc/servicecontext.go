package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-service/internal/model"
	"go-service/service/user_mes/internal/config"
)

type ServiceContext struct {
	Config config.Config
	model.UserMesModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := sqlx.NewSqlConn("mysql", c.Mysql.DataSource)
	return &ServiceContext{
		Config:       c,
		UserMesModel: model.NewUserMesModel(db, c.CacheRedis),
	}
}
