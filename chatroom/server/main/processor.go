package main

import (
	"fmt"
	"net"

	"github.com/wyzzhe/practice-go/chatroom/common/message"
	"github.com/wyzzhe/practice-go/chatroom/server/process"
)

// 根据消息类型调用不同的处理函数
func processMessage(conn net.Conn, mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMesType:
		err = process.ServerProcessLogin(conn, mes)
	case message.RegisterMesType:
		process.ServerProcessRegister(conn)
	default:
		fmt.Println("消息类型不存在，无法处理...")
	}
	return
}
