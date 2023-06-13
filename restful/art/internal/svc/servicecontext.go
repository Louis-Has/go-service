package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-service/internal/model"
	"go-service/restful/art/internal/config"
)

type ServiceContext struct {
	Config       config.Config
	ArticleModel model.ArticleModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := sqlx.NewSqlConn("mysql", "root:development@tcp(localhost:3306)/testDB?parseTime=True&loc=Asia%2FShanghai")
	return &ServiceContext{
		Config:       c,
		ArticleModel: model.NewArticleModel(db),
	}
}
