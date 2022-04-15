/*
 * @title:
 * @Descripttion:
 * @version:
 * @Author: Guoyh
 * @Date: 2022-03-31 11:02:11
 * @LastEditors: Guoyh
 * @LastEditTime: 2022-03-31 11:02:47
 */
package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// 链接数据库
	sqlData := "root:xUQm#=YSmi-i29%Oqw^N^5a,DLGl^U@tcp(127.0.0.1:3306)/cyymall"
	db, err := sql.Open("mysql", sqlData) // sql.Open 不会校验账户密码，只校验连接格式和库类型 如果配置格式出错或者数据库类型名称错误 如：myqls
	if err != nil {
		fmt.Println("sqlData 格式有问题 err:", err)
		return
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("open sql ping success error！ err:", err)
		return
	}
	fmt.Println("链接数据库成功")
}
