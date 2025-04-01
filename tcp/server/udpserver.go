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
func udpServer(a, b, c, d byte, port int) {
	//服务端
	laddr := &net.UDPAddr{
		IP:   net.IPv4(a, b, c, d),
		Port: port,
	}
	listen, err := net.ListenUDP("udp", laddr)
	if err != nil {
		fmt.Println("UDP服务端 listen failed, err:", err)
		return
	}
	defer listen.Close()
	for {
		var data [1024]byte
		n, addr, err := listen.ReadFromUDP(data[:]) // 接收数据
		if err != nil {
			fmt.Println("UDP服务端 read udp failed, err:", err)
			continue
		}
		fmt.Printf("2.UDP服务端%v接收到来自UDP客户端%v的数据长度:%v，数据:%v\n", laddr, addr, n, string(data[:n]))
		snedMessage := "我是来自UDP服务端的数据"
		fmt.Printf("3.UDP服务端%v往UDP客户端%v发送的数据为:%v\n", laddr, addr, snedMessage)
		_, err = listen.WriteToUDP([]byte(snedMessage), addr) // 发送数据
		if err != nil {
			fmt.Println("UDP服务端 write to udp failed, err:", err)
			continue
		}
	}
}
