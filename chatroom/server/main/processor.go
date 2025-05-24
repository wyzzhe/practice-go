package main

import (
	"fmt"
	"io"
	"net"

	"github.com/wyzzhe/practice-go/chatroom/common/message"
	"github.com/wyzzhe/practice-go/chatroom/server/process"
	"github.com/wyzzhe/practice-go/chatroom/server/utils"
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

// 处理客户端连接
func (p *Process) ProcessConn() (err error) {
	// 初始化Transfer结构体
	t := &utils.Transfer{
		Conn: p.Conn,
	}

	// 循环读取当前客户端消息
	for {
		mes, err := t.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("当前客户端已退出，对应的服务端连接协程退出")
				return err
			} else {
				fmt.Println("readPkg err=", err)
				return err
			}
		}
		fmt.Println("读取到的消息体为 mes=", mes)
		err = p.ProcessMessage(&mes)
		if err != nil {
			return err
		}
	}
}
