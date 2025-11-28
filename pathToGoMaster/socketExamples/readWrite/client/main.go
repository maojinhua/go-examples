package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("usage: go run main.go 内容")
		return
	}

	log.Println("begin dial ...")
	conn, err := net.Dial("tcp", ":8888")
	if err != nil {
		log.Println("dial error:", err)
		return
	}
	defer conn.Close()
	log.Println("dial success")

	time.Sleep(2 * time.Second)
	data := os.Args[1]
	n, err := conn.Write([]byte(data))
	if err != nil {
		log.Println("write failed:", err)
	}
	log.Println("write data success:", data, " ", n, "  bytes written")

	time.Sleep(10 * time.Second)
}
