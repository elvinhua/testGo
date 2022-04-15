/*
 * @title:
 * @Descripttion:
 * @version:
 * @Author: Guoyh
 * @Date: 2022-03-30 09:51:07
 * @LastEditors: Guoyh
 * @LastEditTime: 2022-03-30 11:42:14
 */
package main

import (
	"fmt"
	"strings"
)

func main() {
	var ret []string = make([]string, 0, strings.Count("a:b:c", ":")+1)
	// 上下相等
	var rets = make([]string, 0, strings.Count("a:b:c", ":")+1)
	// index := strings.Index("a:b:c", ":")
	fmt.Println(ret)
	fmt.Println(rets)
}
