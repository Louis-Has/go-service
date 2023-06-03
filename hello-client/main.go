package main

import (
	"context"
	"fmt"
	"go-service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(":9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := proto.NewSayHelloClient(conn)

	// Contact the server and print out its response.
	resp, _ := client.HelloL(context.Background(), &proto.HelloRequest{
		RequestName: "Leo",
		Age:         22,
	})

	fmt.Println(resp.GetResponseMes())
}
