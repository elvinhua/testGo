/*
 * @title:panic与recover
 * @Descripttion:go语言的错误展示与修复
 * @version:
 * @Author: Guoyh
 * @Date: 2022-03-04 10:46:56
 * @LastEditors: Guoyh
 * @LastEditTime: 2022-03-04 10:59:53
 */
package main

import "fmt"

func a() {
	fmt.Println("a")
}
func b() {
	defer func() {
		err := recover()
		fmt.Println(err)
		fmt.Println("释放资源")
	}()
	panic("触发Panic")
	fmt.Println("b")
}
func c() {
	fmt.Println("c")
}
func d() {
	fmt.Println("d")
}

func main() {
	a()
	b()
	c()
	d()
}
