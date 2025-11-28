package main

import (
	"client/proto/greeter"
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// 把服务端定义的 proto 文件生成的代码导入进来

func main() {
	// 1. 和服务器建立连接
	// credentials.NewClientTLSFromFile：从输入的证书文件中未客户端构建 TLS 凭证
	// grpc.WithTransportCredentials: 配置连接级别的安全凭证(例如，TLS/SSL),返回一个 DialOption，用于连接服务器
	// 如果不加盖参数的话会报错：
	// grpc.Dial err: grpc: no transport security set (use grpc.WithTransportCredentials(insecure.NewCredentials()) explicitly or set credentials)
	conn, err := grpc.NewClient("127.0.0.1:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("grpc.Dial err:", err)
		return
	}
	defer conn.Close()

	// 2. 创建 gRPC 客户端
	client := greeter.NewGreeterClient(conn)
	// 3. 调用远程方法
	req := &greeter.HelloReq{Name: "张三1"}
	resp, err := client.SayHello(context.Background(), req)
	if err != nil {
		fmt.Println("client.SayHello err:", err)
		return
	}
	fmt.Println("resp:", resp)
}
