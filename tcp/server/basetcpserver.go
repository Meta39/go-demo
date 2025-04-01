package main

import (
	"bufio"
	"fmt"
	"net"
)

/*
TCP server端
一个TCP服务端可以同时连接很多个客户端，例如世界各地的用户使用自己电脑上的浏览器访问淘宝网。
因为Go语言中创建多个goroutine实现并发非常方便和高效，所以我们可以每建立一次链接就创建一个goroutine去处理。

TCP服务端程序的处理流程：

1.监听端口
2.接收客户端请求建立链接
3.创建goroutine处理链接。
*/

// 基础版 TCP服务端（可能会出现粘包情况）
func baseTcpServer(address string) {
	listen, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println("基础服务端 listen failed, err:", err)
		return
	}
	defer listen.Close()
	for {
		conn, err := listen.Accept() // 建立连接
		if err != nil {
			fmt.Println("基础服务端 accept failed, err:", err)
			continue
		}
		go baseTcpServerProcess(conn) // 启动一个goroutine处理连接
	}
}

// 处理函数
func baseTcpServerProcess(conn net.Conn) {
	defer conn.Close()
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:]) // 读取数据
		if err != nil {
			fmt.Println("基础服务端 read from client failed, err:", err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("收到基础客户端发来的数据（可能会有多条）：", recvStr)
		_, _ = conn.Write([]byte(recvStr)) // 发送数据。返回值和错误不是重点，先忽略返回值和错误
	}
}
