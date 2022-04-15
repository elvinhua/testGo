/*
 * @title: context
 * @Descripttion: withDeadline 延迟跳出goroutine
 * @version:
 * @Author: Guoyh
 * @Date: 2022-04-06 18:16:12
 * @LastEditors: Guoyh
 * @LastEditTime: 2022-04-06 18:25:08
 */
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// 设置过期时间
	d := time.Now().Add(50 * time.Millisecond)
	// 配置context
	ctx, cancel := context.WithDeadline(context.Background(), d)
	// 尽管context会过期，但是在任何情况下调用它的cancel函数都是很好的实践
	// 如果不这样做，可能会使上下文及其父类存活的时间超过必要的时间
	defer cancel()
	select {
	// 1秒后打印
	case <-time.After(1 * time.Second):
		fmt.Println("大帅哥")
		// 由于上文设置的是50毫秒超时，所以代码执行时会执行下方语句
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}

}
