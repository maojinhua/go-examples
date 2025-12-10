package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/maojinhua/ginEtcd/echo"
	"github.com/maojinhua/ginEtcd/echo-server/server"
	"github.com/maojinhua/ginEtcd/etcd"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("prot", 50051, "")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	echo.RegisterEchoServer(s, &server.EchoService{})
	// etcd 注册服务
	etcd.ServiceRegister("echo-service1", fmt.Sprintf("localhost:%d", *port))
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
