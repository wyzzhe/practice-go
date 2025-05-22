package main

import (
	"fmt"
	"net"
)

// 处理客户端连接
func process(conn net.Conn) {
	defer conn.Close()

	// 循环读取客户端连接
	for {
		buf := make([]byte, 8096)
		n, err := conn.Read(buf[:4])
		if n != 4 || err != nil {
			fmt.Println("conn.Read err =", err)
			return
		}

		fmt.Println("读取到的buf =", buf[:4])
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

		go process(conn)
	}

}
