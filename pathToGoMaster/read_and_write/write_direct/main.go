package main

import (
	"fmt"
	"io"
	"os"
)

func directWriteByteSliceToFile(path string, data []byte) (int, error) {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return 0, err
	}

	defer func() {
		// 将尚处于内存中的数据写入磁盘
		f.Sync()
		f.Close()
	}()

	return f.Write(data)
}

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
