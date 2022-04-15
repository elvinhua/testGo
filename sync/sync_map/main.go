/*
 * @title:sync.map
 * @Descripttion: Go内置的map不是并发安全的
 * @version:
 * @Author: Guoyh
 * @Date: 2022-03-25 17:17:13
 * @LastEditors: Guoyh
 * @LastEditTime: 2022-03-29 09:23:00
 */
package main

import (
	"fmt"
	"strconv"
	"sync"
)

var m = make(map[string]int)
var lock sync.Mutex

func set(key string, v int) {
	m[key] = v
}
func get(key string) int {
	return m[key]
}

// func main() {
// 	wg := sync.WaitGroup{}
// 	for i := 0; i < 20; i++ {
// 		wg.Add(1)
// 		go func(n int) {
// 			key := strconv.Itoa(n)
// 			lock.Lock()
// 			set(key, n)
// 			lock.Unlock()
// 			fmt.Printf("key:%s,value:%d\n", key, get(key))
// 			wg.Done()
// 		}(i)
// 	}
// 	wg.Wait()
// }

var m2 sync.Map

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			m2.Store(key, n)
			// m2.LoadOrStore() //如果存在，LoadOrStore返回键的现有值。 否则，它存储并返回给定的值。 如果该值已加载，则加载结果为true，如果已存储则为false
			// m2.LoadAndDelete() //LoadAndDelete删除键的值，如果有的话返回之前的值。 加载的结果报告键是否存在。
			// m2.Range() // 循环map中的数据，内部操作逻辑
			value, _ := m2.Load(key)
			fmt.Printf("key:%s,value:%d\n", key, value)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
