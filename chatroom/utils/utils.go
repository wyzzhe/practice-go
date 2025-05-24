package utils

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"

	"github.com/wyzzhe/practice-go/chatroom/common/message"
)

// 从 客户端 / 服务器 读取发来的消息
func ReadPkg(conn net.Conn) (mes message.Message, err error) {
	// 客户端消息读取到缓冲区
	buf := make([]byte, 8096)
	// 消息头长度为4，读取的消息存入buf[:4]
	_, err = conn.Read(buf[:4])
	if err != nil {
		fmt.Println("conn.Read(buf[:4]) err =", err)
		return
	}

	fmt.Println("读取到的消息头长度为buf =", buf[:4])

	// 根据 buf[:4] 转成一个 uint32类型
	pkgLen := binary.BigEndian.Uint32(buf[:4])

	// 根据 pkgLen 读取消息体内容，读取的消息存入buf[:pkgLen]
	n, err := conn.Read(buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		fmt.Printf("conn.Read(buf[:pkgLen]) failed err=%s\n", err)
		return
	}

	// 把 pkgLen 的消息反序列化为message
	err = json.Unmarshal(buf[:pkgLen], &mes)
	if err != nil {
		fmt.Printf("json.Unmarshal(buf[:pkgLen]) failed err=%s\n", err)
		return
	}

	return
}

// 向 客户端 / 服务器 发送消息
func WritePkg(conn net.Conn, data []byte) (err error) {
	// 消息体长度uint32转换为[]byte
	pkgLen := uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[:4], pkgLen)
	// 先发送消息头
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Printf("conn.Write(messsLen) failed, err= %s\n", err)
		return
	}

	// 再发送消息体
	n, err = conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Printf("conn.Write(data) failed, err= %s\n", err)
		return
	}

	return
}
