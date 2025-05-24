package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"

	"github.com/wyzzhe/practice-go/chatroom/client/utils"
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
	binary.BigEndian.PutUint32(buf[:4], pkgLen)

	// 客户端向服务器发送消息头
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes) err=", err)
		return
	}

	fmt.Printf("客户端成功发送消息长度%d, 发送内容%s\n", len(data), string(data))

	// 客户端向服务器发送消息体
	_, err = conn.Write(data)
	if err != nil {
		fmt.Printf("conn.Write(data) failed err=%s\n", err)
		return
	}

	// 客户端接收服务器返回消息
	resMes, err := utils.ReadPkg(conn)
	if err != nil {
		fmt.Printf("utils.ReadPkg(conn) failed err=%s\n", err)
		return
	}

	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(resMes.Data), &loginResMes)
	if err != nil {
		fmt.Printf("json.Unmarshal(bufRead[:4], pkgLen) failed err=%s\n", err)
		return
	}
	// 登陆成功
	if loginResMes.Code == 200 {
		fmt.Println("登陆成功")
	} else if loginResMes.Code == 500 {
		// 登陆失败
		fmt.Println("登陆失败，", loginResMes.Error)
	}
	return
}
