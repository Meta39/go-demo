package main //千万要注意main函数所在package包一定是main包

// package client 	//main函数使用项目包client是错误的！！！

func main() {
	//注意：客户端的通讯协议和地址要和服务端一致。
	go baseTcpClient("127.0.0.1:20000")                 //基础版 TCP 客户端（可能会出现粘包情况）
	go tcpClientStickyPackets("127.0.0.1:20001")        //TCP 客户端粘包（暴露粘包问题）
	go resolveTcpClientStickyPackets("127.0.0.1:20002") //解决 TCP 客户端粘包（解决粘包问题）

	go udpClient(0, 0, 0, 0, 20003, 20004) //UDP客户端

	// 主程序阻塞（按Ctrl+C退出）
	select {}
}
