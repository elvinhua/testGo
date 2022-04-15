/*
 * @title:打开文件读取文件
 * @Descripttion:三种读取方式根据需求选择Read、Bufio、Ioutil
 * @version:
 * @Author: Guoyh
 * @Date: 2022-03-08 16:38:03
 * @LastEditors: Guoyh
 * @LastEditTime: 2022-03-09 17:16:04
 */
package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

/**
 * @name:Read 方式读取文件内容
 * @test: 按照设置的字节数读取
 * @msg:
 * @param {*os.File} file
 * @return {*}
 */
func fileForRead(file *os.File) {
	// tmp := make([]byte, 128)
	var tmp [128]byte
	for {
		n, err := file.Read(tmp[:])
		// 当文件循环结束,会有个到底的io,报出EOF(底,文件结束符)
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Print("读取文件失败，错误数据为：", err)
			// fmt.Printf("读取文件失败，错误数据为：%v ,类型为%T", err, err)
			return
		}
		fmt.Printf("读了%d个字节\n", n)
		// 打印读取的数据并将切片转为字符串输出
		fmt.Println(string(tmp[:n]))
	}
}

/**
 * @name:Buifio 方式读取文件
 * @test: 一行一行的读取
 * @msg:bufio.NewReader返回值设置结束行参数 ReadString()
 * @param {*}
 * @return {*}
 */
func fileForBuifio(file *os.File) {
	// 创建一个读取对象
	reader := bufio.NewReader(file)
	for {
		// 以某个字符结束为一行，for循环读取，这里用换行符'\n'
		s, err := reader.ReadString('\n')
		if err == io.EOF {
			return
		} else if err != nil {
			fmt.Print("读取文件失败,错误信息为:", err)
			return
		}
		fmt.Print(s)
	}
}

/**
 * @name: Ioutil方式读取文件
 * @test: test font
 * @msg:ioutil.ReadFile
 * @param {*}
 * @return {*}
 */
func fuleForIoutil() {
	ret, err := ioutil.ReadFile("./main.go")
	if err != nil {
		fmt.Print("读取文件失败,错误信息为:", err)
	}
	fmt.Print(string(ret))
}

func main() {
	file, err := os.Open("./main.go")

	if err != nil {
		fmt.Print("打开文件失败,错误信息为:", err)
		return
	}
	defer file.Close()
	// Read
	// fileForRead(file)
	// Buifio
	fileForBuifio(file)
	// Ioutil
	// fuleForIoutil()
}
