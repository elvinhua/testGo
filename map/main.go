/*
 * @Descripttion:
 * @version:
 * @Author: Guoyh
 * @Date: 2022-03-02 16:09:01
 * @LastEditors: Guoyh
 * @LastEditTime: 2022-03-03 10:25:23
 */
/**
* 指针
*  1.& 取地址
*  2.* 根据地址取值
 */
package main

import (
	"fmt"
)

// func main() {
// 	m1 := make(map[string]string, 5)
// 	m1["你好"] = "好"
// 	m1["美女"] = "挺好"
// 	fmt.Println(m1)
// 	fmt.Println(m1["美女"])
// 	fmt.Println(m1["帅哥"])
// 	v, ok := m1["帅哥"]
// 	fmt.Println(ok)
// 	fmt.Println(v)
// }

// func main() {
// 	m1 := make(map[string]int, 5)
// 	m1["帅哥"] = 1
// 	m1["你好"] = 2
// 	fmt.Println(m1["帅哥"])
// 	v, ok := m1["帅s哥"]
// 	fmt.Println(v)
// 	fmt.Println(ok)
// }

func main() {
	// 值为切片的的map
	s1 := make(map[string][]int, 5)
	s1["A"] = []int{1, 2, 3}
	fmt.Println(s1)

	// map类型的切片
	s2 := make([]map[int]string, 1, 10)
	/* 没有对内部map做make初始化
	s2[0][100] = "A"
	*/
	s2[0] = make(map[int]string, 1)
	s2[0][10] = "大帅哥"
	fmt.Println(s2)
}
