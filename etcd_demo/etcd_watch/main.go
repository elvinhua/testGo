package main

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("connect for failed err:", err)
	}
	defer cli.Close()
	// 创建一个watch 哨兵监听 key 的变化 （查询除外）
	watchCh := cli.Watch(context.Background(), "name")
	for wresp := range watchCh {
		for _, env := range wresp.Events {
			// 获取被修改的key
			fmt.Printf("type:%s key:%s value:%s\n", env.Type, string(env.Kv.Key), string(env.Kv.Value))
		}
	}
}
