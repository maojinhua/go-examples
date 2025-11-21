package main

import (
	"fmt"
	"net"
	"net/rpc"
)

// 定义类对象
type world struct {
}

// 定义类方法 - 修正了返回值处理
func (this *world) Helloworld(req string, res *string) error {
	*res = req + " 你好!"  // 需要通过指针修改值
	return nil
}

func main() {
	// 1. 注册RPC服务
	err := rpc.RegisterName("hello", new(world))
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
	defer listener.Close()
	fmt.Println("开始监听 127.0.0.1:8800 ...")
	
	// 3. 建立链接并处理请求
	for {
		// 接收连接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Accept() err:", err)
			continue // 修正：发生错误时不退出，继续接受新连接
		}
		
		// 4. 为每个连接创建goroutine处理RPC请求
		go rpc.ServeConn(conn)
	}
}