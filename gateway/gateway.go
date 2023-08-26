package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/gateway"
)

var configFile = flag.String("f", "gateway/gateway.yaml", "config file")

func init() {
	// set log
	logConf := logx.LogConf{
		Level:    "info",
		Encoding: "plain",
	}
	logx.MustSetup(logConf)
	logx.SetLevel(logx.ErrorLevel)
}

func main() {
	flag.Parse()

	var c gateway.GatewayConf
	conf.MustLoad(*configFile, &c)
	gw := gateway.MustNewServer(c)
	defer gw.Stop()

	fmt.Printf("Starting gateway at %v...\n", c.Port)
	gw.Start()
}
