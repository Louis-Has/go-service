package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"go-service/restful/art/internal/config"
	"go-service/restful/art/internal/handler"
	"go-service/restful/art/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var c config.Config

func init() {
	// onload config
	var configFile = flag.String("f", "restful/art/etc/art.yaml", "the config file")
	flag.Parse()
	conf.MustLoad(*configFile, &c)

	// set log
	logConf := logx.LogConf{
		Level:    "info",
		Encoding: "plain",
	}
	logx.MustSetup(logConf)
	logx.DisableStat()
}

func main() {

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
