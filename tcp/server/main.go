package main //千万要注意main函数所在package包一定是main包

// package server	//main函数使用项目包是错误的！！！

func main() {
	//注意：客户端的通讯协议和地址要和服务端一致。
	go baseTcpServer("127.0.0.1:20000")                 //基础版 TCP服务端（可能会出现粘包情况）
	go tcpServerStickyPackets("127.0.0.1:20001")        //TCP 服务端粘包（暴露粘包问题）
	go resolveTcpServerStickyPackets("127.0.0.1:20002") //解决 TCP 服务端粘包（解决粘包问题）

	go udpServer(0, 0, 0, 0, 20003) //UDP服务端

	// 主程序阻塞（按Ctrl+C退出）
	select {}
}
