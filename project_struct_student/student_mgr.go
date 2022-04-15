package main

import "fmt"

type student struct {
	id   int64
	name string
}
type studentMgr struct {
	studentList map[int64]student
}

func (s studentMgr) showStudent() {
	for _, v := range s.studentList {
		fmt.Printf("学号：%d,姓名：%s\n", v.id, v.name)
	}

}
func (s studentMgr) addStudent() {
	var id int64
	var name string
	fmt.Print("请输入学号：")
	fmt.Scanln(&id)
	fmt.Print("请输入学生姓名：")
	fmt.Scanln(&name)
	newSt := student{
		id:   id,
		name: name,
	}
	s.studentList[newSt.id] = newSt

}
func (s studentMgr) editStudent() {
	var id int64
	fmt.Print("请输入要修改的学号：")
	fmt.Scanln(&id)

}
func (s studentMgr) delStudent() {

}
