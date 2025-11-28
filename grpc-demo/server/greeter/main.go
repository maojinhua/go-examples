package main

import (
	"context"
	"fmt"
	"greeter/proto/greeter"
	"net"

	"google.golang.org/grpc"
)

// 1. 定义远程调用的结构体和方法
// 该结构体需要实现 greeter_gprc.pb.go 中定义的 GreeterServer 接口
type Hello struct {
	greeter.UnimplementedGreeterServer
}

func (h *Hello) SayHello(c context.Context, req *greeter.HelloReq) (*greeter.HelloRes, error) {
	fmt.Println("Hello ", req.GetName())

	return &greeter.HelloRes{
		Message: "Hello " + req.GetName(),
	}, nil
}

func main() {
	// 2. 初始化 Grpc 对象
	server := grpc.NewServer()
	// 3. 注册 Greeter 服务
	greeter.RegisterGreeterServer(server, &Hello{})
	// 4. 监听端口
	lis, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer lis.Close()

	// 5. 启动服务
	server.Serve(lis)
}
