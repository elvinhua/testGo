/*
 * @Descripttion:
 * @version:
 * @Author: Guoyh
 * @Date: 2022-03-02 10:43:18
 * @LastEditors: Guoyh
 * @LastEditTime: 2022-04-01 14:39:07
 */
package main

import (
	"fmt"
	"sort"
)

/*
func main() {
	var s1 []int
	var s2 []string
	fmt.Println(s1, s2)

	a1 := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	// s3 := a1[:7]
	// fmt.Println(s3)
	// s4 := s3[:3]
	// fmt.Println(s4)

	// // 给切片追加元素
	// s2 = []string{"北京", "上海", "深圳"}
	// fmt.Printf("len(s2):%v,cap(s2):%v\n", len(s2), cap(s2))
	// s2 = append(s2, "广州", "重庆")
	// fmt.Printf("len(s2):%v,cap(s2):%v", len(s2), cap(s2))
	// fmt.Println(s2)
	// ss := []string{"武汉", "西安"}
	// // ...表示将切片拆开
	// s2 = append(s2, ss...)
	// fmt.Printf("len(s2):%v,cap(s2):%v", len(s2), cap(s2))

	as := a1[:]
	// fmt.Printf("as=%v\n", as)
	as = append(as[:2], as[4:]...)
	fmt.Println(as)
	fmt.Println(a1)
}
*/
func main() {
	a1 := [...]int{3, 5, 2, 1, 7, 4, 9, 6, 8}
	sort.Ints(a1[:])
	fmt.Println(a1)

}
