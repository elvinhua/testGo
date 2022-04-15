/*
 * @title:
 * @Descripttion:
 * @version:
 * @Author: Guoyh
 * @Date: 2022-03-24 16:14:31
 * @LastEditors: Guoyh
 * @LastEditTime: 2022-03-24 16:34:56
 */
package main

import "sync"

var c chan int
var wg sync.WaitGroup

/**
 * @name:并发读写
 * @test: test font
 * @msg:
 * @param {*}
 * @return {*}
 */
func concurrentTest() {
	// 并发读写
	m := map[string]int{}
	m1 := map[string]int{}
	// wg.Add(1)
	go func() {
		// defer wg.Done()
		for {
			m1["key1"] = 1
			m = m1
		}
	}()
	// wg.Wait()
	for {
		_ = m["key2"]
	}
}

func main() {
	concurrentTest()
	// 以下从func内拿出任意一个都会报阻塞但是与goroutine配合就不会报，只会开启两个goroutine任务
	// go func() {
	// 	for v := range c {
	// 		fmt.Println(v)
	// 	}
	// }()
	// go func() {
	// 	c = make(chan int)
	// 	c <- 10
	// 	fmt.Println("done")
	// }()
	// time.Sleep(10 * time.Second)
}
