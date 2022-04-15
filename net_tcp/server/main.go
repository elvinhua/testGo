/*
 * @title:tcp 链接 服务端
 * @Descripttion:
 * @version:
 * @Author: Guoyh
 * @Date: 2022-03-28 17:52:12
 * @LastEditors: Guoyh
 * @LastEditTime: 2022-03-29 13:52:32
 */
package main

import (
	"fmt"
	"io"
	"net"
)

func processConn(conn net.Conn) {
	defer conn.Close()
	// 定义一个接收数据的数组类型
	var tmp [128]byte
	// 循环接收各个客户端传来的消息
	for {
		// 读取消息
		n, err := conn.Read(tmp[:])
		// 如果读完了停止循环接收
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("read error err:", err)
			break
		}
		// 打印转为字符串的消息内容
		fmt.Println(string(tmp[:n]))
	}
}

func main() {
	// 建立tcp 监听
	listen, err := net.Listen("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Printf("the net conn error %v\n", err)
	}
	// 循环等待外部链接
	for {
		// 等待链接进来
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept conn error err:", err)
			break
		}
		// 创建消息函数 将链接传输内容放入
		go processConn(conn)
	}

}
