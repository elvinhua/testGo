/*
 * @title:
 * @Descripttion:
 * @version:
 * @Author: Guoyh
 * @Date: 2022-03-23 15:44:40
 * @LastEditors: Guoyh
 * @LastEditTime: 2022-03-23 17:16:37
 */
package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func a() {
	defer wg.Done()

	for i := 0; i < 10; i++ {
		fmt.Println("A:", i)
	}
}
func b() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Println("B:", i)
	}
}

func main() {
	runtime.GOMAXPROCS(1)
	fmt.Println(runtime.NumCPU())
	wg.Add(2)
	go a()

	go b()
	time.Sleep(1 * time.Second)
	wg.Wait()

}
