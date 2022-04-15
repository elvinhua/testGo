/*
 * @title: 学生管理系统
 * @Descripttion:用结构体方法的方式编写学生管理系统
 * @version:
 * @Author: Guoyh
 * @Date: 2022-03-07 17:51:22
 * @LastEditors: Guoyh
 * @LastEditTime: 2022-03-08 15:31:07
 */
package main

import (
	"fmt"
	"os"
)

var smr studentMgr

func showMenu() {
	fmt.Println(`
	欢迎来到学生管理系统
		1.查看所有学生
		2.添加学生
		3.删除学生
		4.退出
	`)
}

func main() {
	smr = studentMgr{
		studentList: make(map[int64]student, 50),
	}
	for {
		showMenu()
		var check int
		fmt.Print("请输入您的选项：")
		fmt.Scanln(&check)
		switch check {
		case 1:
			// smr.showStudent()
		case 2:
			// smr.addStudent()
		case 3:
			// smr.delStudent()
		case 4:
			os.Exit(1)

		}
	}
}
