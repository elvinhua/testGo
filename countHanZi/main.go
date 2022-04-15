/*
 * @Descripttion:
 * @version:
 * @Author: Guoyh
 * @Date: 2022-02-28 15:52:03
 * @LastEditors: Guoyh
 * @LastEditTime: 2022-02-28 17:53:45
 */
package main

import (
	"fmt"
	"unicode"
)

func countHanZi(str string) (count int) {
	for _, v := range str {
		// unicode.Is配合unicode.Han,字符串 返回Boll
		if unicode.Is(unicode.Han, v) {
			count++
		}
	}
	return
}

func main() {
	// 获取字符串中的中文字数
	str := "Hello你好Golang语言学习"
	// count := 0
	// for _, v := range str {
	// 	if unicode.Is(unicode.Han, v) {
	// 		count++
	// 	}
	// }
	count := countHanZi(str)
	fmt.Printf("汉字的数量为：%v", count)

	// 九九乘法表
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d\t", i, j, i*j)
		}
		fmt.Println()
	}
	fmt.Println()
	// 倒序九九乘法表
	for i := 9; i >= 1; i-- {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d\t", i, j, i*j)
		}
		fmt.Println()
	}

}
