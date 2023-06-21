package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-service/internal/model"
	"go-service/service/article/internal/config"
)

type ServiceContext struct {
	Config config.Config
	model.ArticleModel
	model.AuthorMesModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := sqlx.NewSqlConn("mysql", "root:development@tcp(localhost:3306)/testDB?parseTime=True&loc=Asia%2FShanghai")
	return &ServiceContext{
		Config:         c,
		ArticleModel:   model.NewArticleModel(db),
		AuthorMesModel: model.NewAuthorMesModel(db),
	}
}
