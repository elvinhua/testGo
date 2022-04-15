/*
 * @title:反射
 * @Descripttion:
 * @version:
 * @Author: Guoyh
 * @Date: 2022-03-18 17:32:58
 * @LastEditors: Guoyh
 * @LastEditTime: 2022-03-22 16:58:20
 */
package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

func reflectData(a interface{}) {
	v := reflect.ValueOf(a)
	fmt.Println(v.Kind())
	fmt.Println(v.Elem().Kind())
	v.Elem().SetInt(200)
}

func main() {
	str := student{
		Name:  "小帅",
		Score: 90,
	}
	t := reflect.TypeOf(str)
	fmt.Println(t.Name(), t.Kind())
	// 获取 字段的数量
	fmt.Println(t.NumField())

	var a int64 = 100
	reflectData(&a)
	fmt.Println(a)

}
