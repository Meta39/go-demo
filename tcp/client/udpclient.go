package main

import (
	"fmt"
	"net"
)

/*
UDP协议
UDP协议（User Datagram Protocol）中文名称是用户数据报协议，是OSI（Open System Interconnection，开放式系统互联）参考模型中一种无连接的传输层协议；
不需要建立连接就能直接进行数据发送和接收，属于不可靠的、没有时序的通信，但是UDP协议的实时性比较好，通常用于视频直播相关领域。
*/
func udpClient(a, b, c, d byte, raddrPort, laddrPort int) {
	//连接服务端地址
	raddr := &net.UDPAddr{
		IP:   net.IPv4(a, b, c, d),
		Port: raddrPort,
	}
	//客户端绑定固定端口
	laddr := &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: laddrPort,
	}
	socket, err := net.DialUDP("udp", laddr, raddr)
	if err != nil {
		fmt.Println("连接UDP服务端失败，err:", err)
		return
	}
	defer socket.Close()
	dataString := "我是来自UDP客户端的数据"
	fmt.Printf("1.UDP客户端%s往UDP服务端%s发送的数据为%s\n", laddr, raddr, dataString)
	sendData := []byte(dataString)
	_, err = socket.Write(sendData) // 发送数据
	if err != nil {
		fmt.Println("UDP客户端发送数据失败，err:", err)
		return
	}
	data := make([]byte, 4096)
	n, remoteAddr, err := socket.ReadFromUDP(data) // 接收数据
	if err != nil {
		fmt.Println("UDP客户端接收数据失败，err:", err)
		return
	}
	fmt.Printf("4.UDP客户端%v接收到UDP服务端%v返回的数据长度:%v，数据:%v\n", laddr, remoteAddr, n, string(data[:n]))
}
