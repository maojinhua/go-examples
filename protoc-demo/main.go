package main

import (
	"fmt"
	"protoc-demo/proto/userService"

	"google.golang.org/protobuf/proto"
)

func main() {
	u := &userService.UserInfo{
		Username: "张三",
		Age:      20,
		Hobby:    []string{"吃饭", "睡觉"},
	}
	fmt.Println(u.GetUsername())
	fmt.Println(u.GetPhoneType())

	//  protoBuf 序列化
	data, err := proto.Marshal(u)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 序列化后的数据是二进制数组
	fmt.Println("序列化后的数据：",data)
	fmt.Println("序列化后的数据： string(data)", string(data))

	//  protoBuf 反序列化
	u2 := &userService.UserInfo{}
	err = proto.Unmarshal(data, u2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("反序列化的数据 %#v\n", u2)
}
