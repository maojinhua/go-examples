package main

import (
	"context"
	"fmt"
	"log"

	"github.com/maojinhua/ginEtcd/echo"
	"github.com/maojinhua/ginEtcd/etcd"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	etcd.LoadService("echo-service1")
	addr := etcd.ServiceDiscovery("echo-service1")
	fmt.Println("addr ", addr)

	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := echo.NewEchoClient(conn)
	CallUnaryEcho(client)
}

func CallUnaryEcho(c echo.EchoClient) {
	ctx := context.Background()
	in := echo.EchoMessage{
		Message: "client say hello",
	}
	res, err := c.UnaryEcho(ctx, &in)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Client recv ", res)
}
