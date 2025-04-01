package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

/*
粘包 TCP 服务端（暴露粘包问题）
为什么会出现粘包？
主要原因就是tcp数据传递模式是流模式，在保持长连接的时候可以进行多次的收和发。
“粘包"可发生在发送端也可发生在接收端：
1.由Nagle算法造成的发送端的粘包：Nagle算法是一种改善网络传输效率的算法；
简单来说就是当我们提交一段数据给TCP发送时，TCP并不立刻发送此段数据，而是等待一小段时间看看在等待期间是否还有要发送的数据，若有则会一次把这两段数据发送出去。
2.接收端接收不及时造成的接收端粘包：TCP会把接收到的数据存在自己的缓冲区中，然后通知应用层取数据；
当应用层由于某些原因不能及时的把TCP的数据取出来，就会造成TCP缓冲区中存放了几段数据。
*/
func tcpServerStickyPackets(address string) {
	listen, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println("服务端暴露粘包 listen failed, err:", err)
		return
	}
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("服务端暴露粘包 accept failed, err:", err)
			continue
		}
		go tcpServerStickyPacketsProcess(conn)
	}
}

func tcpServerStickyPacketsProcess(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	var buf [1024]byte
	for {
		n, err := reader.Read(buf[:])
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("read from client failed, err:", err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("收到客户端端暴露粘包发来的数据（可能会有多条）：", recvStr)
	}
}
