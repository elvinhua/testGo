/*
 * @title:context 练习
 * @Descripttion: 外部 结束goroutine所使用
 * @version:
 * @Author: Guoyh
 * @Date: 2022-04-06 11:28:01
 * @LastEditors: Guoyh
 * @LastEditTime: 2022-04-06 18:11:34
 */
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// 1 定义外部变量
// var dumpGo bool
// 2 定义外部通道channel，当通道被放入数据时，停止循环跳出goroutine
// var exitChan = make(chan bool, 1)
// 3 利用context进行跳出goroutine

func f(ctx context.Context) {
	defer wg.Done()
	// 2.
	// LOOP:
	// 3
FORLOOP:
	for {
		fmt.Println("大帅哥")
		time.Sleep(500 * time.Millisecond)
		// 1.
		// if dumpGo {
		// 	break
		// }
		// 2.
		// select {
		// case <-exitChan:
		// 	break LOOP
		// default:
		// }
		// 3.
		select {
		case <-ctx.Done():
			break FORLOOP
		default:
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	wg.Add(1)
	go f(ctx)
	// 定义外部全局变量 5秒后 赋值 true退出
	time.Sleep(5 * time.Second)
	// 1.
	// dumpGo = true
	// 2.
	// exitChan <- true
	// 3
	cancel()
	wg.Wait()
	// main作为主go,如何停止子f()的goroutine

}
