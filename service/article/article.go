package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"

	"go-service/service/article/internal/config"
	articleServer "go-service/service/article/internal/server/article"
	authorServer "go-service/service/article/internal/server/author"
	"go-service/service/article/internal/svc"
	"go-service/service/pb/art"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/article.yaml", "the config file")

func init() {
	// set log
	logConf := logx.LogConf{
		Level:    "info",
		Encoding: "plain",
	}
	logx.MustSetup(logConf)
	logx.DisableStat()
}

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		art.RegisterArticleServer(grpcServer, articleServer.NewArticleServer(ctx))
		art.RegisterAuthorServer(grpcServer, authorServer.NewAuthorServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
