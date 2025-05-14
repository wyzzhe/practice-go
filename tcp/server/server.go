package main

import (
	"bufio"
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()

	for {
		reader := bufio.NewReader(conn)
		// 缓冲区
		var buf [4099]byte
		bufSlice := buf[:]
		n, err := reader.Read(bufSlice)
		if err != nil {
			fmt.Println("read failed, err: ", err)
			break
		}
		// buf中的128个byte都初始化成了0，buf[:n]只截取到实际读到的数据，不显示之后多余的0
		recvStr := string(buf[:n])
		fmt.Println("收到客户端发送的数据: ", recvStr)
		conn.Write([]byte(recvStr))
	}
}

func main() {
	// 1. 监听端口
	listenTCP, err := net.Listen("tcp", "127.0.0.1:20001")
	if err != nil {
		fmt.Println("listen failed, err: ", err)
		return
	}
	// 2. 等待客户端链接
	for {
		conn, err := listenTCP.Accept()
		if err != nil {
			fmt.Println("conn failed, err: ", err)
			continue
		}
		// 3. 处理客户端链接
		go process(conn)
	}
}
