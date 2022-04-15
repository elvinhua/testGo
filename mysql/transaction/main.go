/*
 * @title:事物处理
 * @Descripttion:
 * @version:
 * @Author: Guoyh
 * @Date: 2022-03-31 18:22:02
 * @LastEditors: Guoyh
 * @LastEditTime: 2022-04-07 13:37:31
 */
package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type test struct {
	id   int
	exid int
	num  int
}

func initDB() (err error) {
	// 链接数据库
	sqlData := "root:xUQm#=YSmi-i29%Oqw^N^5a,DLGl^U@tcp(127.0.0.1:3306)/test"
	db, err = sql.Open("mysql", sqlData) // sql.Open 不会校验账户密码，只校验连接格式和库类型 如果配置格式出错或者数据库类型名称错误 如：myqls
	if err != nil {
		fmt.Println("sqlData 格式有问题 err:", err)
		return
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("open sql ping success error！ err:", err)
		return
	}
	// fmt.Println("链接数据库成功")
	// 设置数据库连接池最大链接数  如果超过最大连接数将会等待其它链接释放后才会进行操作
	db.SetMaxOpenConns(10)
	// 设置数据库连接池最大空闲数，多用于时段 如：晚上用户量少的时候可以关闭一些链接
	// db.SetMaxIdleConns(5)
	return
}

/**
 * @name: Msyql事物处理
 * @test: test font
 * @msg: 只有innordb 支持事物
 * @param {*}
 * @return {*}
 */
func transactionDemo() {
	// 开启事务
	// tx, err := db.Begin()
	// if err != nil {
	// 	fmt.Println("mysql transation faild err:", err)
	// 	return
	// }
	// tx.Exec() // 执行sql语句
	// tx.Rollback() // sql语句执行失败后调用回滚
	// tx.Commit() // sql语句执行成功提交事务
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Println("数据库初始化操作失败！err:", err)
	}
}
