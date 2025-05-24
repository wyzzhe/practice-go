package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"

	"github.com/wyzzhe/practice-go/chatroom/common/message"
)

func login(userId int, userPwd string) (err error) {
	// 客户端向服务器请求tcp连接
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err =", err)
		return
	}
	defer conn.Close()

	// 消息类型为(登陆消息)
	// 初始化消息
	var mes message.Message
	mes.Type = message.LoginMesType
	// 初始化登陆消息
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd

	// 序列化登录消息
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	mes.Data = string(data)

	// 序列化消息
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	// 消息长度转为[]byte
	pkgLen := uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[:], pkgLen)

	// 客户端向服务器发送消息
	n, err := conn.Write(buf[:])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes) err=", err)
		return
	}

	fmt.Printf("客户端成功发送消息长度%d, 发送内容%s", len(data), string(data))
	return
}
