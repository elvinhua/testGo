/*
 * @title:chan 中的 select 多路复用
 * @Descripttion:
 * @version:
 * @Author: Guoyh
 * @Date: 2022-03-24 16:51:13
 * @LastEditors: Guoyh
 * @LastEditTime: 2022-03-24 17:52:07
 */
package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	o := make(chan bool)
	// ch := make(chan int, 1)
	// for i := 0; i < 10; i++ {
	// 	select {
	// 	case x := <-ch:
	// 		fmt.Println(x)
	// 	case ch <- i:
	// 	}
	// }

	go func() {
		for {
			select {
			case v := <-c:
				fmt.Println(v)
			case <-time.After(5 * time.Second):
				fmt.Println("timeout")
				o <- true
				// default:
				// 	fmt.Println("等待输入输出数据中……")
			}
		}
	}()
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
	}()
	<-o
	// time.Sleep(6 * time.Second)

	fmt.Println("释放通道o 解除阻塞，代码执行完毕")
}
