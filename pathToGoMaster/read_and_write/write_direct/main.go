// Go 语言精进之路：57.1 直接读写字节流
package main

import (
	"fmt"
	"io"
	"os"
)

// 通过 file 类直接将 字节数组 写入指定文件
func directWriteByteSliceToFile(path string, data []byte) (int, error) {
	// 使用 O_APPEND 追加模式，文件写入时总是会写入到末尾
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return 0, err
	}

	defer func() {
		// 将尚处于内存中的数据写入磁盘
		// 数据写入流程：应用程序 → 标准库缓冲区(如果有用 buffio) → 操作系统页缓存 → 磁盘控制器缓存 → 物理磁盘
		// f.Sync()的调用是为了确保数据的持久性，防止因系统崩溃、断电等异常情况导致已"成功写入"的数据实际上还停留在操作系统缓存中而最终丢失。
		// sync() 会阻塞等待直到数据从操作系统缓存落地到磁盘才返回，对于要求速度快可以丢失的数据(如日志数据)，可以不调用 sync()
		f.Sync()
		f.Close()
	}()

	return f.Write(data)
}

// 通过 file 类直接从文件里读取 字节数组
func directReadByteSliceFromFile(path string, data []byte) (int, error) {
	// 等价于OpenFile(name, O_RDONLY, 0)，即以只读方式打开实体文件
	f, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	// 等价于OpenFile(name, O_RDONLY, 0)，即以只读方式打开实体文件
	return f.Read(data)
}

// 文件读取结束会返回 EOF 错误
func readEof(path string, data []byte) (int, error) {
	f, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	total := 0
	for {
		n, err := f.Read(data)
		if err != nil {
			if err == io.EOF {
				fmt.Println("read meet EOF")
				return total, err
			}
			fmt.Println("read file error:", err)
			return 0, err
		}
		total += n
	}

}

func main() {
	path := "foo.txt"
	text := "Hello, Go Master!"

	n, err := directWriteByteSliceToFile(path, []byte(text))
	if err != nil {
		fmt.Println("write file error:", err)
		return
	}
	fmt.Printf("write %d bytes to file %s\n", n, path)

	buf := make([]byte, 100)
	n, err = directReadByteSliceFromFile(path, buf)
	if err != nil {
		fmt.Println("read file error:", err)
		return
	}
	fmt.Printf("read %d bytes from file %s, data: %s\n", n, path, string(buf[:n]))

	buf2 := make([]byte, 100)
	n, err = readEof(path, buf2)
	if err != nil && err != io.EOF {
		fmt.Println("read file error:", err)
		return
	}
	fmt.Printf("readEof %d bytes from file %s, data: %s\n", n, path, string(buf2[:n]))
}
