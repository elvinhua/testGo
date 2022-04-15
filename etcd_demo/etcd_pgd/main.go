/*
 * etcd 增加（Put）、获取（Get）、删除（Delete）
 */
package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/client/v3"
	"time"
)

func main() {
	// 根据配置信息创建一个客户端连接
	// Endpoints 服务节点，slice切片形式 可传多个
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Printf("connect to etcd failed err: %#v", err)
	}
	fmt.Println("connect to etcd success")
	defer cli.Close()
	// put 添加数据
	// etcd连接提供context配置，延迟关闭，如果操作Put过程超过设定时间，则通知关闭
	value := `[{"path":"D:/tmp/nginx.log","topic":"nginx_log"},{"path":"D:/tmp/redis.log","topic":"redis_log"},{"path":"D:/tmp/mysql.log","topic":"mysql_log"},{"path":"D:/tmp/web.log","topic":"web_log"}]
`
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err = cli.Put(ctx, "/logagent/173.16.30.36/collect_config", value)
	if err != nil {
		fmt.Printf("put to etcd failed, err: %#v", err)
	}
	cancel()

	// get 取数据
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, "name")
	cancel()
	if err != nil {
		fmt.Printf("get from etcd failed, err:%v\n", err)
		return
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("%s:%s\n", ev.Key, ev.Value)
	}
	// del 删除数据
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	delResponse, err := cli.Delete(ctx, "name")
	cancel()
	if err != nil {
		fmt.Printf("del from etcd failed, err:%v\n", err)
		return
	}
	fmt.Println("delete count: ", delResponse.Deleted)
}
