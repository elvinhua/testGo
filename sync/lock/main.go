/*
 * @title:互斥锁
 * @Descripttion:同一个操作，启动多个goroutine如果都从公共区域取值，就容易形成并发取（多个go取到公共区域的同一个值进行操作造成重复动作）。加互斥锁能保证最终执行完毕
 * @version:
 * @Author: Guoyh
 * @Date: 2022-03-25 16:34:46
 * @LastEditors: Guoyh
 * @LastEditTime: 2022-03-25 17:24:27
 */
package main

import (
	"fmt"
	"sync"
)

var x int = 0
var wg sync.WaitGroup
var lock sync.Mutex

func add() {
	for i := 0; i < 5000; i++ {
		lock.Lock()
		x += 1
		lock.Unlock()
	}
	wg.Done()
}
func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()

	fmt.Println(x)
}
