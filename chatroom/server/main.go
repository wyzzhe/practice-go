package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"net"

	"github.com/wyzzhe/practice-go/chatroom/common/message"
)

// 读取客户端发来的消息
func readPkg(conn net.Conn) (mes message.Message, err error) {
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
		fmt.Printf("conn.Read(buf[:pkgLen]) failed err=%s", err)
		return
	}

	// 把 pkgLen 的消息反序列化为message
	err = json.Unmarshal(buf[:pkgLen], &mes)
	if err != nil {
		fmt.Printf("json.Unmarshal(buf[:pkgLen]) failed err=%s", err)
		return
	}

	return
}

// 处理客户端连接
func process(conn net.Conn) {
	defer conn.Close()

	// 循环读取当前客户端消息
	for {
		mes, err := readPkg(conn)
		if err != nil {
			if err == io.EOF {
				fmt.Println("当前客户端已退出，对应的服务端连接协程退出")
				return
			}
			fmt.Println("readPkg err=", err)
		}
		fmt.Println("读取到的消息体为 mes=", mes)
	}
}

func main() {
	fmt.Println("服务器在8889监听...")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	if err != nil {
		fmt.Println("net.Listen err =", err)
		return
	}
	defer listen.Close()
	for {
		fmt.Println("服务器等待客户端连接")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err =", err)
			return
		}
		// 异步处理客户端连接
		go process(conn)
	}
}
