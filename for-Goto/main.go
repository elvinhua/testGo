/*
 * @Descripttion:
 * @version:
 * @Author: Guoyh
 * @Date: 2022-03-01 11:22:53
 * @LastEditors: Guoyh
 * @LastEditTime: 2022-03-01 12:39:18
 */
package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		for j := 'A'; j < 'Z'; j++ {
			if j == 'C' {
				goto XX
			}
			// fmt.Printf("%v-%c\n", i, j)
			fmt.Printf("%v-%c \n", i, j)
		}
	}

XX:
	fmt.Println("over")
}
