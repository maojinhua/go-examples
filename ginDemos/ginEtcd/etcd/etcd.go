package etcd

import clientv3 "go.etcd.io/etcd/client/v3"

func getEtcdEndpoints() []string {
	return []string{"127.0.0.1:2379"}
}

func GetEtcdClient() (*clientv3.Client, error) {
	return clientv3.New(clientv3.Config{
		Endpoints: getEtcdEndpoints(),
	})
}
