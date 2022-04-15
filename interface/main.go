/*
 * @title: interface 接口类型
 * @Descripttion:
 * @version:
 * @Author: Guoyh
 * @Date: 2022-03-08 10:35:12
 * @LastEditors: Guoyh
 * @LastEditTime: 2022-03-08 13:39:20
 */
package main

import "fmt"

type anyfunc interface {
	speak()
	eat(food string)
}

type cat struct {
	food string
	feet int
}
type dog struct {
	food string
	feet int
}
type person struct {
	feet int
}

func (c cat) speak() {
	fmt.Println("喵喵喵~")
}
func (d dog) speak() {
	fmt.Println("旺旺旺~")
}
func (p person) speak() {
	fmt.Println("卧槽，卧槽，卧槽~")
}
func (c cat) eat(f string) {
	fmt.Printf("喵喵喵~吃%s", f)
}
func (d dog) eat(f string) {
	fmt.Printf("旺旺旺~吃%s", f)
}

func (p person) eat(string) {
	fmt.Println("卧槽，卧槽，卧槽~")
}

func pak(x anyfunc, f string) {
	x.speak()
	x.eat(f)
}

func main() {
	// var fc anyfunc
	// c := cat{
	// 	food: "小黄鱼",
	// 	feet: 4,
	// }
	d := dog{
		food: "肉",
		feet: 4,
	}
	// p := person{
	// 	feet: 2,
	// }
	// fc = c
	// fc.speak()
	// fc.eat(c.food)
	pak(d, d.food)
}
