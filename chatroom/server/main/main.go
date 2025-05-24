package main

import (
	"fmt"
	"net"
)

// 处理客户端连接
func processConn(conn net.Conn) {
	defer conn.Close()

	// 初始化Process结构体
	p := &Process{
		Conn: conn,
	}

	err := p.ProcessConn()
	if err != nil {
		fmt.Println("客户端和服务器通讯协程错误 failed, err =", err)
		return
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
