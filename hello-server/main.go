package main

import (
	"context"
	"fmt"
	"github.com/TwiN/go-color"
	"go-service/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

// server is used to implement hello.GreeterServer.
type hello struct {
	proto.UnimplementedSayHelloServer
}

func (receiver hello) HelloL(c context.Context, req *proto.HelloRequest) (res *proto.HelloResponse, err error) {
	return &proto.HelloResponse{ResponseMes: "hello someboy " + req.RequestName}, nil
}

func main() {
	listen, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Println(color.InRed(fmt.Sprintf("Listen start Failed:%s", err)))
		return
	}
	grpcServer := grpc.NewServer()
	proto.RegisterSayHelloServer(grpcServer, &hello{})
	log.Printf("server listening at %v", "9090")

	if err := grpcServer.Serve(listen); err != nil {
		log.Println(color.InRed(fmt.Sprintf("grpcServer start Failed:%s", err)))
		return
	}
}
