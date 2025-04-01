package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"net"
)

/*
解决办法
出现"粘包"的关键在于接收方不确定将要传输的数据包的大小，因此我们可以对数据包进行封包和拆包的操作。
封包：封包就是给一段数据加上包头，这样一来数据包就分为包头和包体两部分内容了(过滤非法包时封包会加入"包尾"内容)；
包头部分的长度是固定的，并且它存储了包体的长度，根据包头长度固定以及包头中含有包体长度的变量就能正确的拆分出一个完整的数据包。
我们可以自己定义一个协议，比如数据包的前4个字节为包头，里面存储的是发送的数据的长度。
*/
func resolveTcpServerStickyPackets(address string) {
	listen, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println("解决粘包服务端 listen failed, err:", err)
		return
	}
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("解决粘包服务端 accept failed, err:", err)
			continue
		}
		go resolveTcpServerStickyPacketsProcess(conn)
	}
}

// 处理函数
func resolveTcpServerStickyPacketsProcess(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		msg, err := Decode(reader)
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println("解决粘包服务端 decode msg failed, err:", err)
			return
		}
		fmt.Println("收到解决粘包客户端发来的数据（每次只有一条）：", msg)
	}
}

// Decode 解码消息
func Decode(reader *bufio.Reader) (string, error) {
	header := 4                  //因为客户端 Encode 封包的头部是int32，占4个字节，所以这里是4。如果用其它类型，可能是其它数字，所以不必纠结这里为什么是4.
	headerInt32 := int32(header) //int -> int32
	// 读取消息的长度
	lengthByte, _ := reader.Peek(header) // 读取前4个字节的数据（即：读取消息头）
	lengthBuff := bytes.NewBuffer(lengthByte)
	var length int32
	err := binary.Read(lengthBuff, binary.LittleEndian, &length)
	if err != nil {
		return "", err
	}
	// Buffered返回缓冲中现有的可读取的字节数。
	if int32(reader.Buffered()) < length+headerInt32 {
		return "", err
	}

	// 读取真正的消息数据
	pack := make([]byte, int(headerInt32+length)) //读取的数据 = 消息头 + 消息实体
	_, err = reader.Read(pack)
	if err != nil {
		return "", err
	}
	return string(pack[header:]), nil //用切片截取掉消息头，只保留消息实体。
}
