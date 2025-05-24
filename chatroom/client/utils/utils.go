package utils

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"

	"github.com/wyzzhe/practice-go/chatroom/common/message"
)

// 注意：工具类函数不应该打印错误，防止重复打印，应该交由中间层（打印业务错误）或顶层函数（打印所有未捕获的错误）处理

// 定义传输者结构体
type Transfer struct {
	Conn net.Conn   // 客户端与服务器的连接
	Buf  [8096]byte // 服务器接收消息缓冲区，数组会自动初始化为数组值类型的零值
}

// 从 客户端 / 服务器 读取发来的消息
func (t *Transfer) ReadPkg() (mes message.Message, err error) {
	// 客户端消息读取到缓冲区
	// 消息头长度为4，读取的消息存入buf[:4]
	_, err = t.Conn.Read(t.Buf[:4])
	if err != nil {
		fmt.Println("conn.Read(buf[:4]) err =", err)
		return
	}

	fmt.Println("读取到的消息头长度为buf =", t.Buf[:4])

	// 根据 buf[:4] 转成一个 uint32类型
	pkgLen := binary.BigEndian.Uint32(t.Buf[:4])

	// 根据 pkgLen 读取消息体内容，读取的消息存入buf[:pkgLen]
	n, err := t.Conn.Read(t.Buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		fmt.Printf("conn.Read(buf[:pkgLen]) failed err=%s\n", err)
		return
	}

	// 把 pkgLen 的消息反序列化为message
	err = json.Unmarshal(t.Buf[:pkgLen], &mes)
	if err != nil {
		fmt.Printf("json.Unmarshal(buf[:pkgLen]) failed err=%s\n", err)
		return
	}

	return
}

// 向 客户端 / 服务器 发送消息
func (t *Transfer) WritePkg(data []byte) (err error) {
	// 消息体长度uint32转换为[]byte
	pkgLen := uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[:4], pkgLen)
	// 先发送消息头
	n, err := t.Conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Printf("conn.Write(messsLen) failed, err= %s\n", err)
		return
	}

	// 再发送消息体
	n, err = t.Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Printf("conn.Write(data) failed, err= %s\n", err)
		return
	}

	return
}
