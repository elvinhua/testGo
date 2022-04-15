/*
 * @title:sync.Once 某个程序只执行一次
 * @Descripttion:
 * @version:
 * @Author: Guoyh
 * @Date: 2022-03-25 17:06:17
 * @LastEditors: Guoyh
 * @LastEditTime: 2022-03-25 17:17:57
 */
package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var once sync.Once

func f1(a chan int) {
	defer wg.Done()
	for i := 1; i < 101; i++ {
		a <- i
	}
	close(a)
}
func f2(a <-chan int, b chan<- int) {
	defer wg.Done()
	for x := range a {
		b <- x * x
	}
	// close(b)
	once.Do(func() { close(b) })
}

func main() {

	// 创建读取通道示例
	a := make(chan int, 50)
	b := make(chan int, 100)
	wg.Add(3)
	go f1(a)
	// 创建多个同样goroutine 关闭时只关闭一次就可以这里用到sync.Once，否则报panic: send on closed channel(关闭已经关闭的通道)
	go f2(a, b)
	go f2(a, b)
	wg.Wait()
	for v := range b {
		fmt.Println(v)
	}
}
