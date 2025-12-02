package main

import (
	"memCache/memCache/cache"
	cacheadapter "memCache/memCache/cache-adapter"
)

// 简易实现一个内存缓存系统。
// 通过 map[string]any 进行存储，通过锁来保证并发安全
// 有以下功能：
// 1.可以设置过期时间，精确到秒级
//
//	  	通过每次 get 的时候判断 key 是否过期，过期了则直接 delete
//			启动 go 协程，ticker 每隔几秒钟检查key是否过期，过期了则直接 delete
//
// 2.支持设定最大内存，当内存超出时做出合适的处理
// 3.支持并发安全
// 4.使用 cache-adapter 适配器模式来适配不同的 Set 方法
func main() {
	cach := cache.NewMemCache()
	cach.SetMaxMemory("200MB")
	cach.Set("name", "golang", 0)
	cach.Set("name2", 2, 0)
	value, ok := cach.Get("name")
	if ok {
		println("name value", value.(string))
	}
	cach.Delete("name")
	cach.Flush()
	cach.Keys()

	cache2 := cacheadapter.NewMemCache()
	cache2.SetMaxMemory("200MB")
	cache2.Set("name", "golang")

}
