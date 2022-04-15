/*
 * @title:atomic 原子操作
 * @Descripttion:
 * @version:
 * @Author: Guoyh
 * @Date: 2022-03-25 18:24:28
 * @LastEditors: Guoyh
 * @LastEditTime: 2022-03-28 17:51:26
 */
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup
var x int64
var y int64

func add() {
	defer wg.Done()
	//并发安全下写入数据
	atomic.AddInt64(&x, 1)
	// 并发安全下设置数据
	atomic.StoreInt64(&y, 5)
}

func main() {
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go add()
	}
	wg.Wait()
	// 并发安全下取值
	tem := atomic.LoadInt64(&y)
	fmt.Println(tem)
}
