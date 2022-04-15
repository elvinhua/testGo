/*
 * @title:结构体
 * @Descripttion:一个结构体，一个构造方法（逻辑处理后返回赋值后的结构体），一个对应结构体的接收者方法，结构体可以直接调用
 * @version:
 * @Author: Guoyh
 * @Date: 2022-03-04 17:04:24
 * @LastEditors: Guoyh
 * @LastEditTime: 2022-03-07 10:36:23
 */
package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string, age int) *person {
	return &person{
		name: name,
		age:  age,
	}
}

func (p *person) hello() {
	fmt.Printf("你好%v岁的%v", p.age, p.name)
}

func main() {
	// p := person{
	// 	"帅哥",
	// 	18,
	// }
	p := newPerson("小帅哥", 18)
	// ap := &p
	// fmt.Printf("%v\n", p)
	// fmt.Println(ap)
	p.hello()
}
