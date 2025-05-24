package main

import (
	"fmt"
	"io"
	"net"

	"github.com/wyzzhe/practice-go/chatroom/server/utils"
)

// 处理客户端连接
func processConn(conn net.Conn) {
	defer conn.Close()

	// 初始化Process结构体
	p := &Process{
		Conn: conn,
	}

	// 初始化Transfer结构体
	t := &utils.Transfer{
		Conn: conn,
	}

	// 循环读取当前客户端消息
	for {
		mes, err := t.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("当前客户端已退出，对应的服务端连接协程退出")
				return
			}
			fmt.Println("readPkg err=", err)
		}
		fmt.Println("读取到的消息体为 mes=", mes)
		err = p.ProcessMessage(&mes)
		if err != nil {
			return
		}
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
		go processConn(conn)
	}
}
