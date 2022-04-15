/*
 * @title: 闭包
 * @Descripttion: 一个函数的变量+这个函数内部的函数组合而成(套娃式)
 * @version:
 * @Author: Guoyh
 * @Date: 2022-03-03 18:38:50
 * @LastEditors: Guoyh
 * @LastEditTime: 2022-03-03 18:56:13
 */
package main

import "fmt"

func addr(x int) func(int) int {
	return func(i int) int {
		x += i
		return x
	}
}

func f1(f func()) {
	fmt.Println("this is f1")
	f()
}
func f2(x, y int) {
	fmt.Println("this is f2")
	fmt.Println(x + y)
}
func f3(f func(int, int), a, b int) func() {
	return func() {
		f(a, b)
	}
}

func main() {
	// fmt.Println(addr(1)(2)) //300

	f1(f3(f2, 1, 2))
}
