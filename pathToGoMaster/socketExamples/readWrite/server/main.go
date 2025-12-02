package main

import (
	"log"
	"net"
)

// 使用 net 包无法判断数据的边界，需要自己处理读写边界
// 可以采用自定义写一个是，比如说 协议格式: [4字节长度][实际数据]
// 读取的时候，先读取4字节长度，然后根据长度读取实际数据
func handleConn(c net.Conn) {
	defer c.Close()
	// for 循环需要判断边界
	// 当客户端断开连接时，Read 会返回 EOF 错误，此时调用 c.Write 已无效
	for {
		// 从连接上读取数据
		buf := make([]byte, 10)
		log.Println("start to read from conn...")
		n, err := c.Read(buf)
		if err != nil {
			log.Println("conn read error:", err)
			return
		}
		log.Printf("read %d bytes from conn,data:%s\n", n, string(buf[:n]))

	}
	c.Write([]byte("hello"))
}

func main() {
	// 1. 监听端口
	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatal("listen error:", err)
	}

	for {
		log.Println("waiting for client...")
		// 2. 接收连接
		conn, err := listener.Accept()
		if err != nil {
			log.Println("accept error:", err)
			continue
		}
		log.Println("accept success")

		// 3. 处理连接
		go handleConn(conn)
	}
}
