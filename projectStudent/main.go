/*
 * @title:学生管理系统
 * @Descripttion:综合所学，写一个简单的小系统
 * @version:
 * @Author: Guoyh
 * @Date: 2022-03-07 11:30:33
 * @LastEditors: Guoyh
 * @LastEditTime: 2022-03-08 09:41:55
 */
package main

import (
	"fmt"
	"os"
)

var studentList map[int64]*student

type student struct {
	id   int64
	name string
}

func newStudent(id int64, name string) *student {
	return &student{
		id:   id,
		name: name,
	}
}

func allStudent() {
	fmt.Println("这里是所有学生")
	for k, v := range studentList {
		fmt.Printf("学号：%d,姓名：%s\n", k, v.name)
	}
}
func addStudent() {
	var id int64
	var name string
	fmt.Print("请输入学号：")
	fmt.Scanln(&id)
	fmt.Print("请输入学生姓名：")
	fmt.Scanln(&name)
	newSt := newStudent(id, name)
	studentList[id] = newSt
}
func delStudent() {
	var delId int64
	fmt.Print("请输入要删除的学生学号：")
	fmt.Scanln(&delId)
	delete(studentList, delId)
}

func main() {
	studentList = make(map[int64]*student, 10)
	for {
		fmt.Println("欢迎来到学生管理系统")
		fmt.Println(`
			1.查看所有学生
			2.添加学生
			3.删除学生
			4.退出
		`)
		var check int
		fmt.Print("请输入您的选项：")
		fmt.Scanln(&check)
		switch check {
		case 1:
			allStudent()
		case 2:
			addStudent()
		case 3:
			delStudent()
		case 4:
			os.Exit(1)

		}
	}
}
