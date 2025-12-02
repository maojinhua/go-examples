package cache

import (
	"fmt"
	"sync"
	"time"
)

type memCache struct {
	// 最大内存
	maxMemorySize int64
	// 最大内存字符串表示
	maxMemorySizeStr string
	// 当前内存
	currentMemorySize int64

	// 缓存键值对
	values map[string]*memCacheValue
	// 锁
	locker sync.RWMutex

	// 清除过期缓存时间间隔
	clearExpiredItemInterval time.Duration
}

type memCacheValue struct {
	// value 值
	val any
	// 过期时间
	expireTime time.Time
	// 有效时长,为 0 表示永久有效
	expire time.Duration
	//  value 大小
	size int64
}

var _ Cache = (*memCache)(nil)

func NewMemCache() Cache {
	mc := &memCache{
		values: make(map[string]*memCacheValue,0),
		clearExpiredItemInterval: 5 * time.Second,
	}
	go mc.clearExpiredItem()
	return mc
}

// size : 1KB 100Kb 1MB 2MB 1 GB
func (mc *memCache) SetMaxMemory(size string) bool {
	mc.maxMemorySize, mc.maxMemorySizeStr = ParseSize(size)
	fmt.Println("SetMaxMemory ", mc.maxMemorySize, mc.maxMemorySizeStr)
	return false
}

// 将 value 写入缓存
func (mc *memCache) Set(key string, val any, expire time.Duration) bool {
	fmt.Println("Set")
	v := &memCacheValue{
		val:        val,
		expireTime: time.Now().Add(expire),
		expire:     expire,
		size:       GetValueSize(val),
	}
	mc.locker.Lock()
	defer mc.locker.Unlock()
	mc.del(key)
	mc.add(key, v)
	if mc.currentMemorySize >= mc.maxMemorySize {
		mc.del(key)
		panic(fmt.Sprintf("max memory size %s", mc.maxMemorySizeStr))
	}
	fmt.Println("set end")
	return true
}

func (mc *memCache) get(key string) (*memCacheValue, bool) {
	mc.locker.Lock()
	defer mc.locker.Unlock()
	val, ok := mc.values[key]
	if !ok {
		return nil, false
	}
	// 判断缓存是否过期
	if val.expire != 0 && val.expireTime.Before(time.Now()) {
		mc.del(key)
		return nil, false
	}

	return val, ok
}

func (mc *memCache) del(key string) {
	tmp, ok := mc.values[key]
	if !ok {
		return
	}
	if ok && tmp != nil {
		mc.currentMemorySize -= tmp.size
		delete(mc.values, key)
	}
}

func (mc *memCache) add(key string, val *memCacheValue) {
	mc.values[key] = val
	mc.currentMemorySize += val.size
}

// 根据 key 值获取 value
func (mc *memCache) Get(key string) (interface{}, bool) {
	v,e := mc.get(key)
	if e{
		return v.val,true
	}
	return nil,false
}

// 删除 key 值
func (mc *memCache) Delete(key string) bool {
	fmt.Println("Delete")
	mc.locker.Lock()
	defer mc.locker.Unlock()
	mc.del(key)
	return true
}

func (mc *memCache) Exists(key string) bool {
	mc.locker.RLock()
	defer mc.locker.RUnlock()
	_, ok := mc.values[key]
	return ok
}

// 清空所有 key
func (mc *memCache) Flush() bool {
	fmt.Println("Flush")
	mc.locker.Lock()
	defer mc.locker.Unlock()
	mc.values = make(map[string]*memCacheValue, 0)
	mc.currentMemorySize = 0
	fmt.Println("Flush end")
	return true
}

// 获取缓存中所有 key 的数量
func (mc *memCache) Keys() int64 {
	fmt.Println("Keys")
	mc.locker.RLock()
	defer mc.locker.RUnlock()
	return int64(len(mc.values))

}

func (mc *memCache) clearExpiredItem() {
	ticker := time.NewTicker(mc.clearExpiredItemInterval)
	defer ticker.Stop()
	for range ticker.C {
		for key, val := range mc.values {
			if val.expire != 0 && time.Now().After(val.expireTime) {
				mc.locker.Lock()
				mc.del(key)
				mc.locker.Unlock()
			}
		}
	}
}
