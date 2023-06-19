package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"go-service/internal/model"
	"go-service/restful/art/internal/config"
	"go-service/service/pb/art"
)

type ServiceContext struct {
	Config       config.Config
	ArticleModel model.ArticleModel
	ArtServer    art.ArticleClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := sqlx.NewSqlConn("mysql", "root:development@tcp(localhost:3306)/testDB?parseTime=True&loc=Asia%2FShanghai")

	// login server
	conn := zrpc.MustNewClient(zrpc.RpcClientConf{
		Target: "dns:///127.0.0.1:8080",
	})
	client := art.NewArticleClient(conn.Conn())

	return &ServiceContext{
		Config:       c,
		ArticleModel: model.NewArticleModel(db),
		ArtServer:    client,
	}
}
