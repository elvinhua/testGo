/*
 * @title:通道
 * @Descripttion:
 * @version:
 * @Author: Guoyh
 * @Date: 2022-03-23 17:20:09
 * @LastEditors: Guoyh
 * @LastEditTime: 2022-03-25 17:17:16
 */
package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var once sync.Once

// 有缓冲区的通道 （有make空间）
func bufChannel() {
	a := make(chan int, 1)
	a <- 10
	fmt.Println("10放入a通道中")
	x := <-a
	fmt.Println("从通道中取值为：", x)
	close(a)
}

// 无缓冲区通道
func noBufChannel() {
	b := make(chan int)
	b <- 10
	fmt.Println("10放入b通道中")
}

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
	// for {
	// 	x, ok := <-a
	// 	if !ok {
	// 		break
	// 	}
	// 	b <- x * x
	// }
	// close(b)
	once.Do(func() { close(b) })
}

func main() {
	// 带缓冲区
	// bufChannel()
	// 不带缓冲区
	// noBufChannel()
	// 创建读取通道示例
	/*
		a := make(chan int, 50)
		b := make(chan int, 100)
		wg.Add(3)
		go f1(a)
		go f2(a, b)
		go f2(a, b)
		wg.Wait()
		for v := range b {
			fmt.Println(v)
		}
	*/
	/*
	   channel 异常总结
	*/
	// 1.未make过的channel
	// var noMakeStr chan int
	// fmt.Println("此时打印值是 nil", noMakeStr)
	// noMakeStr <- 10
	// fmt.Println("发送值此时打印值是 阻塞睡眠状态  err：all goroutines are asleep", noMakeStr)
	// tmp := <-noMakeStr
	// fmt.Println("数据发送不过去发送值此时打印值是 阻塞睡眠状态  err：all goroutines are asleep", tmp)
	// close(noMakeStr) // 关闭未make的会报panic

	// 2.make通道以后
	// var makeStr chan int
	// makeStr = make(chan int)
	// fmt.Println("此时打印值是 空间地址", makeStr)
	// tem2 := <-makeStr
	// fmt.Println("当tem2接收空的通道传输时 变为 阻塞睡眠状态 err：all goroutines are asleep", tem2)
	makeStr2 := make(chan int, 1)
	// 接空值
	// tem3 := <-makeStr2
	// fmt.Println("接空值 阻塞", tem3)
	// 发送值
	makeStr2 <- 10
	fmt.Println("发送到通道中", makeStr2)
	// 继续发送（此时已塞满）
	// makeStr2 <- 20
	fmt.Println("通道塞满以后打印报阻塞", makeStr2)

}
