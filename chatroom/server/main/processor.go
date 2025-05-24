package main

import (
	"fmt"
	"net"

	"github.com/wyzzhe/practice-go/chatroom/common/message"
	"github.com/wyzzhe/practice-go/chatroom/server/process"
)

// 定义Process结构体
type Process struct {
	Conn net.Conn // 客户端与服务器的连接
}

// 根据消息类型调用不同的处理函数
func (p *Process) ProcessMessage(mes *message.Message) (err error) {

	switch mes.Type {
	case message.LoginMesType:
		// 初始化userProcess
		up := &process.UserProcess{
			Conn: p.Conn,
		}
		err = up.ServerProcessLogin(mes)
	case message.RegisterMesType:
	default:
		fmt.Println("消息类型不存在，无法处理...")
	}
	return
}
