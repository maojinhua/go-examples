package cache

import "time"

type Cache interface {
	// size : 1KB 100Kb 1MB 2MB 1 GB
	SetMaxMemory(size string) bool
	// 将 value 写入缓存
	Set(key string, val interface{}, expire time.Duration) bool
	// 根据 key 值获取 value
	Get(key string) (interface{}, bool)
	// 删除 key 值
	Delete(key string) bool
	Exists(key string) bool
	// 清空所有 key
	Flush() bool
	// 获取缓存中所有 key 的数量
	Keys() int64
}
