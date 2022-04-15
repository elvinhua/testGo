/*
 * @title:sync包的锁机制
 * @Descripttion:当有读写机制时候，读写加互斥锁比较慢，读不加锁要比加锁要快，当读加Rlock读锁时候，比其它都快。写是肯定要加锁的
 * @version:
 * @Author: Guoyh
 * @Date: 2022-03-25 15:13:06
 * @LastEditors: Guoyh
 * @LastEditTime: 2022-03-25 16:31:16
 */
package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	x      = 0
	wg     sync.WaitGroup
	lock   sync.Mutex
	rwlock sync.RWMutex
)

func read() {
	defer wg.Done()
	// rwlock.RLock()
	lock.Lock()
	fmt.Println(x)
	time.Sleep(time.Millisecond)
	// rwlock.RUnlock()
	lock.Unlock()
}
func write() {
	defer wg.Done()
	lock.Lock()
	x += 1
	time.Sleep(time.Millisecond * 5)
	lock.Unlock()
}

func main() {
	startTime := time.Now()
	for i := 0; i < 10; i++ {
		go write()
		wg.Add(1)
	}
	time.Sleep(time.Second)
	for i := 0; i < 1000; i++ {
		go read()
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(time.Now().Sub(startTime))
}
