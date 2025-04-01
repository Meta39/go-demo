package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

/*
基础版 TCP 客户端（可能会出现粘包情况）
一个TCP客户端进行TCP通信的流程如下：
1.建立与服务端的链接
2.进行数据收发
3.关闭链接
*/
func baseTcpClient(address string) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println("err :", err)
		return
	}
	defer conn.Close() // 关闭连接
	inputReader := bufio.NewReader(os.Stdin)
	for {
		input, _ := inputReader.ReadString('\n') // 读取命令行用户输入的内容
		inputInfo := strings.Trim(input, "\r\n")
		if strings.ToUpper(inputInfo) == "Q" { // 如果输入q/Q就退出
			fmt.Println("用户输入q/Q，已停止基础客户端命令行输入发送消息......")
			return
		}
		_, err = conn.Write([]byte(inputInfo)) // 发送数据
		if err != nil {
			return
		}
		buf := [512]byte{}
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Println("基础客户端 recv failed, err:", err)
			return
		}
		fmt.Println(string(buf[:n]))
	}
}
