package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
)

/*
解决办法：客户端进行封包
出现"粘包"的关键在于接收方不确定将要传输的数据包的大小，因此我们可以对数据包进行封包和拆包的操作。
封包：封包就是给一段数据加上包头，这样一来数据包就分为包头和包体两部分内容了(过滤非法包时封包会加入"包尾"内容)；
包头部分的长度是固定的，并且它存储了包体的长度，根据包头长度固定以及包头中含有包体长度的变量就能正确的拆分出一个完整的数据包。
我们可以自己定义一个协议，比如数据包的前4个字节为包头，里面存储的是发送的数据的长度。
*/
func resolveTcpClientStickyPackets(address string) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println("解决粘包客户端 dial failed, err", err)
		return
	}
	defer conn.Close()
	for i := 0; i < 20; i++ {
		msg := `Hello, Hello. How are you?`
		data, err := Encode(msg)
		if err != nil {
			fmt.Println("解决粘包客户端 encode msg failed, err:", err)
			return
		}
		conn.Write(data)
	}
}

// Encode 将消息编码
func Encode(message string) ([]byte, error) {
	// 读取消息的长度，转换成int32类型（占4个字节）
	var length = int32(len(message))
	var pkg = new(bytes.Buffer)
	// 写入消息头
	err := binary.Write(pkg, binary.LittleEndian, length)
	if err != nil {
		return nil, err
	}
	// 写入消息实体
	err = binary.Write(pkg, binary.LittleEndian, []byte(message))
	if err != nil {
		return nil, err
	}
	//最后pkg = 消息头 + 消息实体，所以读取的时候要处理掉消息头，再读消息实体，防止数据解码乱码或解码失败
	return pkg.Bytes(), nil
}
