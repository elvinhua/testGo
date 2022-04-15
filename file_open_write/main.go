/*
 * @title:写文件
 * @Descripttion:
 * @version:
 * @Author: Guoyh
 * @Date: 2022-03-09 10:50:07
 * @LastEditors: Guoyh
 * @LastEditTime: 2022-03-09 16:30:51
 */
package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

/**
 * @name: Write 直接写操作
 * @test: test font
 * @msg:
 * @param {*os.File} file
 * @return {*}
 */
func fileForWrite(file *os.File) {
	// file.Write([]byte(""))
	file.WriteString("你好,再次叫你大帅哥！你敢答应吗？")
}

/**
 * @name: Bufio
 * @test: bufio.NewWrite
 * @msg:Bufio 要将打开的文件开启（newWrite）一个新的缓存对象
 * @param {*os.File} file
 * @return {*}
 */
func fileForBuifioWrite(file *os.File) {
	// 创建写的对象
	w := bufio.NewWriter(file)
	// 将字符串内容录入缓存
	numByte, err := w.WriteString("不敢答应\n")
	if err != nil {
		fmt.Println("写入缓存失败，错误为：", err)
		return
	}
	//将录入好的内容导入文件中
	w.Flush()
	fmt.Printf("写入了%d个字符", numByte)
}

/**
 * @name:Ioutil 直接写入
 * @test: test font
 * @msg:
 * @param {*}
 * @return {*}
 */
func fileForIoutil() {
	str := "这里用Ioutil写入\n"
	err := ioutil.WriteFile("./xx.txt", []byte(str), 0666)
	if err != nil {
		fmt.Println("写入内容失败，错误为：", err)
		return
	}
}

func main() {
	file, err := os.OpenFile("./xx.txt", os.O_TRUNC|os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("打开文件失败,错误为:", err)
		return
	}
	defer file.Close()

	// fileForWrite(file)
	// fileForBuifioWrite(file)
	// fileForIoutil()
}
