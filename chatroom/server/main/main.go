package main

import (
	"fmt"
	"io"
	"net"

	"github.com/wyzzhe/practice-go/chatroom/server/utils"
)

// // 处理登录请求
// func serverProcessLogin(conn net.Conn, mes *message.Message) (err error) {
// 	// 解析登录请求
// 	var loginMes message.LoginMes
// 	err = json.Unmarshal([]byte(mes.Data), &loginMes)
// 	if err != nil {
// 		fmt.Printf("json.Unmarshal([]byte(mes.Data), &loginMes) failed, err= %s\n", err)
// 	}

// 	// 构造消息类型和数据
// 	var resMes message.Message
// 	resMes.Type = message.LoginResMesType
// 	// 构造回复消息数据
// 	var loginResMes message.LoginResMes

// 	// 判断登录请求是否合理
// 	if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
// 		// 合理
// 		loginResMes.Code = 200
// 	} else {
// 		// 不合理
// 		loginResMes.Code = 500
// 		loginResMes.Error = "用户未注册"
// 	}

// 	// 将 loginResMes 序列化
// 	data, err := json.Marshal(loginResMes)
// 	if err != nil {
// 		fmt.Printf("json.Marshal(loginResMes) failed, err= %s\n", err)
// 		return
// 	}

// 	// 将 loginResMes 赋值给message
// 	resMes.Data = string(data)

// 	// 将 resMes 序列化
// 	data, err = json.Marshal(resMes)
// 	if err != nil {
// 		fmt.Printf("json.Marshal(loginResMes) failed, err= %s\n", err)
// 		return
// 	}

// 	// 发送消息体
// 	err = utils.WritePkg(conn, data)
// 	return
// }

// // 处理注册请求
// func serverProcessRegister(conn net.Conn) {

// }

// // 根据消息类型调用不同的处理函数
// func processMessage(conn net.Conn, mes *message.Message) (err error) {
// 	switch mes.Type {
// 	case message.LoginMesType:
// 		err = serverProcessLogin(conn, mes)
// 	case message.RegisterMesType:
// 		serverProcessRegister(conn)
// 	default:
// 		fmt.Println("消息类型不存在，无法处理...")
// 	}
// 	return
// }

// 处理客户端连接
func processConn(conn net.Conn) {
	defer conn.Close()

	// 循环读取当前客户端消息
	for {
		mes, err := utils.ReadPkg(conn)
		if err != nil {
			if err == io.EOF {
				fmt.Println("当前客户端已退出，对应的服务端连接协程退出")
				return
			}
			fmt.Println("readPkg err=", err)
		}
		fmt.Println("读取到的消息体为 mes=", mes)
		err = processMessage(conn, &mes)
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
