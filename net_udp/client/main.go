/*
 * @title:
 * @Descripttion:
 * @version:
 * @Author: Guoyh
 * @Date: 2022-03-29 13:14:37
 * @LastEditors: Guoyh
 * @LastEditTime: 2022-03-29 13:48:32
 */
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// 与服务端建立链接 DialUDP
	clientStoket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 30000,
	})
	if err != nil {
		fmt.Println("DialUDP error err:", err)
		return
	}
	defer clientStoket.Close()
	// 定义接收的回复信息
	var reply [1024]byte

	// 读取命令窗口的输入信息
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("请输入：")
		// 将数据以\n为终点截取
		msg, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("ReadString error err:", err)
			break
		}
		// 将信息整理好写入网络发送服务端
		clientStoket.Write([]byte(msg))
		// 读取回复的信息
		rn, _, err := clientStoket.ReadFromUDP(reply[:])
		if err != nil {
			fmt.Println("ReadFromUDP error err:", err)
			break
		}
		fmt.Println("收到回复", string(reply[:rn]))
	}
}
