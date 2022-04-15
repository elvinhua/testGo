/*
 * @Descripttion:
 * @version:
 * @Author: Guoyh
 * @Date: 2022-03-03 15:04:59
 * @LastEditors: Guoyh
 * @LastEditTime: 2022-03-03 16:45:41
 */
package main

import "fmt"

//先赋值给返回值
// 再执行defer
// 再执行返回

func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x //将x赋值给设定的返回项，然后执行defer，最后执行返回
	// 5
}

func f2() (x int) {
	fmt.Printf("defer外1：%d\n", x)
	defer func() {
		fmt.Printf("defer内1：%d\n", x)
		x++
		fmt.Printf("defer内2：%d\n", x)
	}()
	fmt.Printf("defer外2：%d\n", x)
	return 5
	// 6
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return 5

	// 5
}

func f4() (x int) {
	fmt.Printf("defer外：%d\n", x)
	defer func(x int) {
		x++
		fmt.Printf("defer内：%d\n", x)
	}(x)
	return 5
	// 5
}

// func f5() (x int) {
// 	defer func(x int) {

// 	}(x)
// 	return 5
// }

func main() {
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
}
