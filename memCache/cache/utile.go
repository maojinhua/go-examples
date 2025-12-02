package cache

import (
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

// 每个单位都差 2 的10次方倍，采用左移运算符实现
const (
	B = 1 << (iota * 10)
	KB
	MB
	GB
	TB
	PB
)

func ParseSize(size string) (int64, string) {
	// 默认大小为 100 MB
	// 正则表达式匹配数字部分
	re, _ := regexp.Compile("[0-9]+")
	// 将正则表达式匹配到的数字部分全部转换成 空字符串 就得到了单位部分
	unit := string(re.ReplaceAll([]byte(size), []byte("")))
	// 获取 num 部分,将 size 字符串中的 unit 单位替换成 空字符串 就得到了 num 部分
	num, err := strconv.ParseInt(strings.Replace(size, unit, "", 1), 10, 64)
	if err != nil {
		fmt.Println("parse size error:", err)
		return 100 * MB, "MB"
	}
	unit = strings.ToUpper(unit)
	var byteNum int64 = 0
	switch unit {
	case "B":
		byteNum = num
	case "KB":
		byteNum = num * KB
	case "MB":
		byteNum = num * MB
	case "GB":
		byteNum = num * GB
	case "TB":
		byteNum = num * TB
	case "PB":
		byteNum = num * PB
	default:
		log.Println("ParseSize 仅支持 B、KB、Mb、GB、TB、PB")
		return 100 * MB, "MB"
	}

	sizeStr := fmt.Sprintf("%d %s", num, unit)
	return byteNum, sizeStr
}

func GetValueSize(val any) int64{
	bytes ,_:=json.Marshal(val)
	size := int64(len(bytes))
	fmt.Println("val ",val,"size ",size)
	return int64(size)
}