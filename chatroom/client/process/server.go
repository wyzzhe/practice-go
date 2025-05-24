package process

import (
	"fmt"
	"os"
)

// 显示登录成功的界面
func ShowMenu(userId int) {
	for {
		fmt.Printf("----------恭喜%d登陆成功----------\n", userId)
		fmt.Println("\t\t\t 1. 显示在线用户列表")
		fmt.Println("\t\t\t 2. 发送消息")
		fmt.Println("\t\t\t 3. 信息列表")
		fmt.Println("\t\t\t 4. 退出系统")
		fmt.Println("\t\t\t 请选择(1-4)...")

		var key int
		fmt.Scanf("%d", &key)
		switch key {
		case 1:
			fmt.Println("显示在线用户列表")
		case 2:
			fmt.Println("发送消息")
		case 3:
			fmt.Println("信息列表")
		case 4:
			fmt.Println("你退出了系统")
			os.Exit(0)
		default:
			fmt.Println("输入错误，请重新输入")
		}
	}
}
