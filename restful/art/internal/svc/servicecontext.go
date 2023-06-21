package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"go-service/restful/art/internal/config"
	"go-service/service/pb/art"
)

type ServiceContext struct {
	Config config.Config
	art.ArticleClient
	art.AuthorClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	// login server
	conn := zrpc.MustNewClient(zrpc.RpcClientConf{
		Target: "dns:///127.0.0.1:8080",
	})

	return &ServiceContext{
		Config:        c,
		ArticleClient: art.NewArticleClient(conn.Conn()),
		AuthorClient:  art.NewAuthorClient(conn.Conn()),
	}
}
