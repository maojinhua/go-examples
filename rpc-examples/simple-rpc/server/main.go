package main

import (
	"fmt"
	"net"
	"net/rpc"
)

// 创建远程调用的函数，函数一般是放在结构体里面
type world struct {
}

type HelloReq struct {
	Id      int
	Title   string
	Price   float32
	Content string
}

type HelloRes struct {
	Success bool
	Message string
}

// 远程调用的函数
func (this *world) Helloworld(req HelloReq, res *HelloRes) error {
	fmt.Println("req ", req)
	*res = HelloRes{true, "hello world"}
	return nil
}

type GetGoodsReq struct {
	Id int
}

type GetGoodsRes struct {
	Id      int
	Title   string
	Price   float32
	Content string
}

func (this *world) GetGoods(req GetGoodsReq, res *GetGoodsRes) error {
	fmt.Println("req ", req)
	*res = GetGoodsRes{req.Id, "商品1", 15.5, "商品内容"}
	return nil
}

func main() {
	// 1. 注册RPC服务
	err := rpc.RegisterName("hello2", new(world))
	if err != nil {
		fmt.Println("注册RPC服务失败:", err)
		return
	}

	// 2. 设置监听
	listener, err := net.Listen("tcp", "127.0.0.1:8800")
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	// 3. 关闭监听
	defer listener.Close()
	fmt.Println("开始监听 127.0.0.1:8800 ...")

	// 4. 建立链接并处理请求
	for {
		// 5.接收连接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Accept() err:", err)
			continue // 修正：发生错误时不退出，继续接受新连接
		}

		// 6.. 当有新连接时，为每个连接创建goroutine处理RPC请求
		go rpc.ServeConn(conn)
	}
}
