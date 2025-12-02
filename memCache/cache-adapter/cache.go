package cacheadapter

import (
	"memCache/memCache/cache"
	"time"
)

// cache 适配器
type cacheAdapter struct {
	memCache cache.Cache
}

func NewMemCache() *cacheAdapter {
	return &cacheAdapter{
		memCache: cache.NewMemCache(),
	}
}

// size : 1KB 100Kb 1MB 2MB 1 GB
func (ca *cacheAdapter) SetMaxMemory(size string) bool {
	return ca.memCache.SetMaxMemory(size)
}

// 将 value 写入缓存
// 使用适配器模式 将 expire 设置为可选参数
func (ca *cacheAdapter) Set(key string, val interface{}, expire ...time.Duration) bool {
	expireTs := 0*time.Second
	if len(expire) > 0 {
		expireTs = expire[0]
	}
	return ca.memCache.Set(key, val, expireTs)
}

// 根据 key 值获取 value
func (ca *cacheAdapter) Get(key string) (interface{}, bool) {
	return  ca.memCache.Get(key)
}

// 删除 key 值
func (ca *cacheAdapter) Delete(key string) bool {
	return  ca.memCache.Delete(key)
}
func (ca *cacheAdapter) Exists(key string) bool {
	return  ca.memCache.Exists(key)
}

// 清空所有 key
func (ca *cacheAdapter) Flush() bool {
	return  ca.memCache.Flush()
}

// 获取缓存中所有 key 的数量
func (ca *cacheAdapter) Keys() int64 {
	return  ca.memCache.Keys()
}
