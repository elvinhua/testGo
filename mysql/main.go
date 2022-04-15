/*
 * @title:
 * @Descripttion:
 * @version:
 * @Author: Guoyh
 * @Date: 2022-03-31 09:54:04
 * @LastEditors: Guoyh
 * @LastEditTime: 2022-04-02 16:23:59
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
 * @name:单条查询
 * @test: test font
 * @msg:
 * @param {int} id
 * @return {*}
 */
func queryOne(id int) {
	// 单条查询 “?”占位符
	var t test
	sqlStr := "select id,exid,num from test_nums where id=?;"
	// 传入地址 Scan是用指针接收的
	err := db.QueryRow(sqlStr, id).Scan(&t.id, &t.exid, &t.num)
	if err != nil {
		fmt.Println("select row error err:", err)
	}
	fmt.Printf("t:%#v\n", t)
}

/**
 * @name: 多条查询
 * @test: test font
 * @msg:
 * @param {int} n
 * @return {*}
 */
func queryMore(n int) {
	sqlStr := "select id,exid,num from test_nums where id>?"
	rows, err := db.Query(sqlStr, n)
	if err != nil {
		fmt.Println("Query Rows error err:", err)
		return
	}
	defer rows.Close()
	// rows.next （如果有值就返回bool）循环取值
	for rows.Next() {
		var t test
		err = rows.Scan(&t.id, &t.exid, &t.num)
		if err != nil {
			fmt.Println("select rows error err:", err)
			return
		}
		fmt.Printf("t:%#v\n", t)
	}
}

/**
 * @name: 插入数据
 * @test: test font
 * @msg:
 * @param {*}
 * @return {*}
 */
func insertData() {
	// INSERT INTO 表名称 VALUES (值1, 值2,....)
	// sqlStr := "insert into test_nums values (5,6)"
	// INSERT INTO table_name (列1, 列2,...) VALUES (值1, 值2,....)
	sqlStr := "insert into test_nums(exid,num) values (5,6)"
	res, err := db.Exec(sqlStr)
	if err != nil {
		fmt.Println("insert into failed err", err)
		return
	}
	// 执行完毕后返回影响的条数
	line, _ := res.RowsAffected()
	// 插入数据后获取插入后的自增id
	id, err := res.LastInsertId()
	if err != nil {
		fmt.Println("get id faild err:", err)
		return
	}
	fmt.Printf("id:%d,line:%d\n", id, line)
}

/**
 * @name: 更新数据
 * @test: test font
 * @msg:
 * @param {int} newExid
 * @param {int} id
 * @return {*}
 */
func updateRow(newExid int, id int) {
	sqlStr := "update test_nums set exid=? where id=?"
	res, err := db.Exec(sqlStr, newExid, id)
	if err != nil {
		fmt.Println("update failed! err:", err)
		return
	}
	// 执行完毕后返回影响的条数
	line, err := res.RowsAffected()
	if err != nil {
		fmt.Println("updata sucess error err", err)
		return
	}
	fmt.Printf("更新了:%d 行\n", line)
}

/**
 * @name: 删除数据
 * @test: test font
 * @msg:
 * @param {int} id
 * @return {*}
 */
func deleteData(id int) {
	sqlStr := "delete from test_nums where id=?"
	res, err := db.Exec(sqlStr, id)
	if err != nil {
		fmt.Println("delete faild err:", err)
		return
	}
	// 执行完毕后返回影响的条数
	line, err := res.RowsAffected()
	if err != nil {
		fmt.Println("updata sucess error err", err)
		return
	}
	fmt.Printf("删除了:%d 行\n", line)
}

// 预处理
func prepareData() {
	sqlStr := "insert into test_nums(exid,num) value(?,?)"
	stmt, err := db.Prepare(sqlStr) // 预处理先把语句编译好
	if err != nil {
		fmt.Println("prepare faild err:", err)
		return
	}
	defer stmt.Close() // 记得关闭
	var m = map[int]int{
		1: 2,
		3: 4,
		5: 6,
		7: 8,
	}
	for k, v := range m {
		stmt.Exec(k, v)
	}

}

func main() {
	err := initDB()
	if err != nil {
		fmt.Println("数据库初始化操作失败！err:", err)
	}
	// queryOne(3)
	// queryMore(1)
	// insertData()
	// updateRow(2, 5)
	// deleteData(7)
	prepareData()
}
