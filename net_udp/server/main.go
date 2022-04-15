/*
 * @title:UDP协议 服务端
 * @Descripttion:广泛用于直播行业
 * @version:
 * @Author: Guoyh
 * @Date: 2022-03-29 11:08:20
 * @LastEditors: Guoyh
 * @LastEditTime: 2022-03-29 13:40:48
 */
package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	// 创建UDP监听
	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 30000,
	})
	if err != nil {
		fmt.Println("listerUDP error err:", err)
	}
	defer listen.Close()
	// 定义一个接收值
	var msg [1024]byte
	// 循环接收信息
	for {
		// 读取数据获取对方发送的数据和地址以及错误信息
		n, addr, err := listen.ReadFromUDP(msg[:])
		if err != nil {
			fmt.Println("ReadFromUDP error err:", err)
		}
		// 将获取到的数据转切片后转为字符串，并变为大写
		reply := strings.ToUpper(string(msg[:n]))
		// 将转好的数据回复给对方
		rn, err := listen.WriteToUDP([]byte(reply), addr)
		if err != nil {
			fmt.Println("WriteMsgUDP error err:", err)
		}
		fmt.Printf("发送成功,发送了%d个字节,全文：%#v\n", rn, string(msg[:rn]))
	}

}
