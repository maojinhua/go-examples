package etcd

import (
	"context"
	"fmt"
	"log"
	"sync"

	clientv3 "go.etcd.io/etcd/client/v3"
)

// 服务器：服务注册
func ServiceRegister(serviceName, addr string) error {
	cli, err := GetEtcdClient()
	if err != nil {
		log.Fatal("GetEtcdClient err ", err)
	}
	key := getKey(serviceName)

	ctx := context.Background()
	// 创建租约
	leaseRes, err := cli.Grant(ctx, 10)
	if err != nil {
		return err
	}
	// 向 etcd 写数据，并把租约和 key 关联
	_, err = cli.Put(ctx, key, addr, clientv3.WithLease(leaseRes.ID))
	if err != nil {
		return err
	}
	// 保持租约
	keepAliveCh, err := cli.KeepAlive(ctx, leaseRes.ID)
	if err != nil {
		return err
	}

	go func() {
		for res := range keepAliveCh {
			fmt.Printf("leaseId:%v,ttl:%v\n", res.ID, res.TTL)
		}
	}()

	return nil
}

type serviceCache struct {
	data map[string]string
	sync.RWMutex
}

// ---------------------------------- 客户端----------------------------------
var cache *serviceCache

func init() {
	cache = &serviceCache{
		data: make(map[string]string, 0),
	}
}

func getKey(serviceName string) string {
	return serviceName
}

// 服务发现
func ServiceDiscovery(serviceName string) string {
	cache.RLock()
	defer cache.RUnlock()
	return cache.data[serviceName]
}

// 第一次加载服务信息
func LoadService(serviceName string) {
	cli, err := GetEtcdClient()
	if err != nil {
		log.Fatal("GetEtcdClient ", err)
	}
	ctx := context.Background()
	key := getKey(serviceName)
	res, err := cli.Get(ctx, key)
	if err != nil {
		log.Fatal(err)
	}

	// 第一次加载服务信息
	if res.Count > 0 {
		cache.Lock()
		defer cache.Unlock()
		for _, item := range res.Kvs {
			cache.data[string(item.Key)] = string(item.Value)
		}
	}
}

// 监听服务
func WatchService(serviceName string) {
	cli, err := GetEtcdClient()
	if err != nil {
		log.Fatal("GetEtcdClient ", err)
	}
	ctx := context.Background()
	key := getKey(serviceName)

	cache.Lock()
	defer cache.Unlock()
	rch := cli.Watch(ctx, key)
	for wres := range rch {
		for _, event := range wres.Events {
			if event.Type == clientv3.EventTypeDelete {
				delete(cache.data, key)
				continue
			}
			if event.Type == clientv3.EventTypePut {
				cache.data[key] = string(event.Kv.Value)
				continue
			}
		}
	}
}
