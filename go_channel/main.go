/*
 * @title:goroutine与channel结合使用
 * @Descripttion:
 * @version:
 * @Author: Guoyh
 * @Date: 2022-03-24 14:17:35
 * @LastEditors: Guoyh
 * @LastEditTime: 2022-03-24 16:08:20
 */
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

type Job struct {
	x int64
}
type Result struct {
	Job    *Job
	result int64
}

// var jobChan chan *Job
var jobChan = make(chan *Job, 100)
var resultChan = make(chan *Result, 100)

func randNum(j chan<- *Job) {
	defer wg.Done()
	for {
		v := rand.Int63()
		newJob := &Job{
			x: v,
		}
		j <- newJob
		time.Sleep(time.Millisecond * 500)
	}
}

func comNum(j <-chan *Job, res chan<- *Result) {

	for {
		job := <-j
		sum := int64(0)
		n := job.x
		for n > 0 {
			sum += n % 10
			n = n / 10
		}
		newResule := &Result{
			Job:    job,
			result: sum,
		}
		res <- newResule
	}
}

func main() {
	wg.Add(1)
	go randNum(jobChan)
	wg.Add(24)
	for i := 0; i < 24; i++ {
		go comNum(jobChan, resultChan)
	}
	for v := range resultChan {
		fmt.Println(v.Job.x, v.result)
	}
	wg.Wait()
}
