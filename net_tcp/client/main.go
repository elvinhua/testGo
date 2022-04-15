/*
 * @title:
 * @Descripttion:
 * @version:
 * @Author: Guoyh
 * @Date: 2022-03-28 18:15:14
 * @LastEditors: Guoyh
 * @LastEditTime: 2022-03-29 13:59:55
 */
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// 与服务端建立tcp网络链接
	conn, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println("client err err:", err)
	}
	// 获取窗口输入的内容 os.Stdin是标准输入指针
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("请输入：")
		msg, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("reader string error err:", err)
			break
		}
		// 去除两边空格
		msg = strings.TrimSpace(msg)
		// 将数据发送给服务端
		conn.Write([]byte(msg))
	}
	conn.Close()
}
