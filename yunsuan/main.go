/*
 * @Descripttion:
 * @version:
 * @Author: Guoyh
 * @Date: 2022-03-01 16:57:01
 * @LastEditors: Guoyh
 * @LastEditTime: 2022-03-02 09:23:26
 */
package main

import "fmt"

func main() {
	var (
		a = 2
		b = 3
		c = 6
	)

	/*
		a 二进制为 10;5的二进制位101
		或一一 按位或 两值比较有1则为1
		或有一说一
	*/
	fmt.Println(a | 5) //值为7
	/*
		b 二进制为 11;
		与均1 按位与 两值比较均为1则为1
		与均一一
	*/
	fmt.Println(b & 5) //值为1

	/*
		c 二进制为110
		异或1 异或 两只比较不同则为1

	*/
	fmt.Println(c ^ 5) //值为3
}
