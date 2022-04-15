/*
 * @title:runtime
 * @Descripttion:
 * @version:
 * @Author: Guoyh
 * @Date: 2022-03-10 16:15:48
 * @LastEditors: Guoyh
 * @LastEditTime: 2022-03-10 16:43:35
 */
package main

import (
	"fmt"
	"runtime"
)

func main() {

	// fmt.Println(runtime.CPUProfile())
	pc, file, line, ok := runtime.Caller(0)
	funcNme := runtime.FuncForPC(pc).Name()
	fmt.Println(pc)
	fmt.Println(funcNme)
	fmt.Println(file)
	fmt.Println(line)
	fmt.Println(ok)
	// runtime.Callers()
}
