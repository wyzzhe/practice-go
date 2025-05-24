package main

import (
	"fmt"
	"os"

	"github.com/wyzzhe/practice-go/chatroom/client/process"
)

var userId int
var userPwd string

func main() {
	var key int

	for {
		fmt.Println("----------欢迎登陆多人聊天系统----------")
		fmt.Println("\t\t\t 1 登录聊天室")
		fmt.Println("\t\t\t 2 注册用户")
		fmt.Println("\t\t\t 3 退出系统")
		fmt.Println("\t\t\t 请选择(1-3)")

		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("登录聊天室")
			fmt.Println("请输入用户的id")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户的密码")
			fmt.Scanf("%s\n", &userPwd)

			// 初始化UserProcess结构体
			up := &process.UserProcess{}
			// 用户登录
			if err := up.Login(userId, userPwd); err != nil {
				return
			}
		case 2:
			fmt.Println("注册用户")
		case 3:
			fmt.Println("退出系统")
			os.Exit(0)
		default:
			fmt.Println("你的输入有误，请重新输入")
		}
	}
}
