/*
 * @title:time 时间包
 * @Descripttion:
 * @version:
 * @Author: Guoyh
 * @Date: 2022-03-09 17:25:03
 * @LastEditors: Guoyh
 * @LastEditTime: 2022-03-24 17:22:33
 */
package main

import (
	"fmt"
	"time"
)

func main() {

	t := time.Now()
	// fmt.Println("当前时区：", )
	fmt.Println(t)
	/**
		 * @name: 当前时间加1小时 Add参数为d Duration
		 * @test: test font
		 * @msg: 	Nanosecond  Duration = 1
	    				Microsecond          = 1000 * Nanosecond
	    				Millisecond          = 1000 * Microsecond
	    				Second               = 1000 * Millisecond
	    				Minute               = 60 * Second
	    				Hour                 = 60 * Minute
		 * @param {*}
		 * @return {*}
	*/
	//
	addT := t.Add(time.Hour)
	fmt.Println(addT)
	// 当前时间加几年、几个月、几天 返回的是加后的时间
	addD := t.AddDate(2, 0, 0)
	fmt.Println(addD)
	// 判断传入的时间是否在当前时间之后 返回bool
	fmt.Println(t.After(addT))
	// 判断两个时间间隔多长时间
	fmt.Println(t.Sub(addT))
	parseTime, err := time.Parse("2006-01-02", "2019-07-08")
	if err != nil {
		fmt.Println("时间错误：", err)
		return
	}
	fmt.Println(parseTime)
	// 判断两个时间是否相同（包含时区的比较）
	isT := t.Equal(t)
	fmt.Println(t)
	fmt.Println(time.Now())
	fmt.Println(isT)
	time.Sleep(5 * time.Second)
}
