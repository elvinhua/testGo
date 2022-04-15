/*
 * @title:字符串
 * @Descripttion:
 * @version:
 * @Author: Guoyh
 * @Date: 2022-03-22 18:19:27
 * @LastEditors: Guoyh
 * @LastEditTime: 2022-03-23 15:44:24
 */
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a string = "100"
	// strconv.ParseInt()
	// 字符串转int
	intV, _ := strconv.Atoi(a)
	// int转字符串
	strV := strconv.Itoa(intV)
	fmt.Printf("Type:%T,Value:%v\n", intV, intV)
	fmt.Printf("Type:%T,Value:%v", strV, strV)
}
