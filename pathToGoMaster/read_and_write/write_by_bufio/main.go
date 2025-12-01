// Go 语言精进之路：57.3 使用缓存区读取和写入数据，减少磁盘 io 次数
// 当直接创建文件或打开已存在的文件的，这样后续对文件的读写都是无缓冲的，即每次读都会驱动磁盘运转来读取数据，
// 每次写（并随后调用Sync）也都会对数据进行落盘处理。这种频繁的磁盘I/O是无缓冲I/O模式性能不高的主因
// 可以使用缓冲区来解决，
// 读取数据：在读取数据时缓冲区会一次性读取指定缓冲区大小的数据，后续读取数据从缓冲区里读取(消费)
// 写入数据：缓冲区满了或者调用 sync() 时才会写入向磁盘写入数据
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 缓冲区写入和读取文件

func main() {
	file := "bufio.txt"

	f, err := os.OpenFile(file, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("open file error", err)
		return
	}
	defer func() {
		f.Sync()
		f.Close()
	}()

	data := []byte("I love golang!\n")

	// 创建带缓冲 I/O 的 Writer
	bio := bufio.NewWriterSize(f, 33)

	// 将 15 字节写入 bio 缓冲区，缓冲区缓存 15 字节，但并未写入文件
	bio.Write(data)
	// 将 15 字节写入 bio 缓冲区，缓冲区缓存 30 字节，但并未写入文件
	bio.Write(data)

	// 将 15 字节写入 bio 缓冲区，由于缓冲区满了，bufio 一次性将 33 字节写入文件
	// 缓冲区中任然缓存 15*3 - 33 字节
	bio.Write(data)

	// 强制将缓冲区剩余数据写入文件
	bio.Flush()

	f2, err := os.Open(file)
	if err != nil {
		fmt.Println("open file error", err)
		return
	}
	defer f.Close()

	// 创建带缓冲 I/O 的 Reader
	// 初始缓冲区大小为 64 字节,缓冲区中没有数据，在第一次读取时才会去文件读取数据，才会发生磁盘 io
	br := bufio.NewReaderSize(f2, 64)
	fmt.Printf("初始化状态下缓冲区缓存数据量=%d字节\n\n", br.Buffered())

	var i int = 1
	for {
		data := make([]byte, 15)
		n, err := br.Read(data)
		if err == io.EOF {
			fmt.Printf("第%d次读取到数据，读到文件末尾，程序退出\n", i)
			return
		}
		if err != nil {
			fmt.Println("read file error:", err)
			return
		}
		fmt.Printf("第%d次读取到数据：%s,长度=%d\n", i, data, n)
		fmt.Printf("当前缓冲区缓存数据量=%d字节\n\n", br.Buffered())
		i++
	}
}
