package main

import (
	"fmt"
	"net/rpc"
)

// 客户端和服务端的 Req 和 Res 结构体是一样的
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

type GetGoodsReq struct {
	Id int
}

type GetGoodsRes struct {
	Id      int
	Title   string
	Price   float32
	Content string
}

func main() {
	// 1.用 rpc 链接服务器 -- Dial()
	conn, err := rpc.Dial("tcp", "127.0.0.1:8800")
	if err != nil {
		fmt.Println("Dial err:", err)
		return
	}
	// 2. 关闭链接
	defer conn.Close()

	//  3. 调用远程函数
	rep := HelloReq{1, "标题", 15.5, "内容"}
	var reply HelloRes
	err = conn.Call("hello2.Helloworld", rep, &reply)
	if err != nil {
		fmt.Println("Call: ", err)
		return
	}
	fmt.Println("reply:", reply)

	// 4. 调用另一个远程函数
	goodsReq := GetGoodsReq{1}
	var goodsRes GetGoodsRes
	err = conn.Call("hello2.GetGoods", goodsReq, &goodsRes)
	if err != nil {
		fmt.Println("Call: ", err)
		return
	}
	fmt.Println("goodsRes:", goodsRes)
}
