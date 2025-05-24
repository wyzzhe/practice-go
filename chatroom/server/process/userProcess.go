package process

import (
	"encoding/json"
	"fmt"
	"net"

	"github.com/wyzzhe/practice-go/chatroom/common/message"
	"github.com/wyzzhe/practice-go/chatroom/server/utils"
)

// 处理登录请求
func ServerProcessLogin(conn net.Conn, mes *message.Message) (err error) {
	// 解析登录请求
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Printf("json.Unmarshal([]byte(mes.Data), &loginMes) failed, err= %s\n", err)
	}

	// 构造消息类型和数据
	var resMes message.Message
	resMes.Type = message.LoginResMesType
	// 构造回复消息数据
	var loginResMes message.LoginResMes

	// 判断登录请求是否合理
	if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
		// 合理
		loginResMes.Code = 200
	} else {
		// 不合理
		loginResMes.Code = 500
		loginResMes.Error = "用户未注册"
	}

	// 将 loginResMes 序列化
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Printf("json.Marshal(loginResMes) failed, err= %s\n", err)
		return
	}

	// 将 loginResMes 赋值给message
	resMes.Data = string(data)

	// 将 resMes 序列化
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Printf("json.Marshal(loginResMes) failed, err= %s\n", err)
		return
	}

	// 发送消息体
	err = utils.WritePkg(conn, data)
	return
}

// 处理注册请求
func ServerProcessRegister(conn net.Conn) {

}
