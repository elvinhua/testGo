/*
 * @Descripttion:
 * @version:
 * @Author: Guoyh
 * @Date: 2022-02-28 10:15:42
 * @LastEditors: Guoyh
 * @LastEditTime: 2022-04-01 15:59:32
 */
package main

import (
	"fmt"
	"strings"
)

const (
	a = iota
	b
	c
)

func main() {
	sa := "你好啊,请叫我王的男人"
	ss := 124
	// sb := "大宝贝儿"
	// fmt.Println(len(sa))
	// fmt.Println(sa + sb)
	// 字符串
	// sc := fmt.Sprintf("%s%s", sb, sa)
	// fmt.Printf("%s\n", sc)

	// 以逗号切割字符串
	str := strings.Split(sa, ",")
	// fmt.Println(str)
	// 字符串中包含某字符 true or false
	// fmt.Println(strings.Contains(sa, "叫"))
	// fmt.Println(strings.Contains(sa, "pei"))
	// 当前字符是否是字符串前缀
	// fmt.Println(strings.HasPrefix(sa, "你好"))
	// 当前字符是否是字符串后缀 true or false
	// fmt.Println(strings.HasSuffix(sa, "男人"))

	sd := "abcdedg"
	// 返回字符串第一次出现的位置
	fmt.Println(strings.Index(sd, "d"))
	// 返回字符串最后出现的位置
	fmt.Println(strings.LastIndex(sd, "d"))
	fmt.Println(strings.Join(str, "+"))
	fmt.Printf("%#v", ss) //加#会把原值输出 字符串：字符串+引号包裹 一般输出：你好，加#会输出："你好"
}
