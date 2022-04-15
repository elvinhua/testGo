/*
 * @title:flag包
 * @Descripttion:
 * @version:
 * @Author: Guoyh
 * @Date: 2022-03-30 13:34:10
 * @LastEditors: Guoyh
 * @LastEditTime: 2022-03-30 17:51:55
 */
package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	/*
	 * flag.Args() // 返回命令行参数后的其他参数，以[]string类型
	 * flag.NArg() // 返回命令行参数后的其他参数个数
	 * flag.NFlag() // 返回使用的命令行参数个数
	 */
	// 创建一个标志位参数
	name := flag.String("name", "大帅", "此处要输入名字")
	age := flag.Int("age", 32, "此处要输入年龄")
	married := flag.Bool("married", true, "此处要输入是否结婚")
	marriedTime := flag.Duration("ct", time.Second, "此处要输入结婚多久")

	// 也可以使用下面的方法
	var names string
	flag.StringVar(&names, "name", "大帅", "此处要输入名字")
	// flag.IntVar()
	// flag.BoolVar()
	// ……

	// 使用flag（必须 ）
	flag.Parse()
	fmt.Println(*name, *age, *married, *marriedTime)
}
