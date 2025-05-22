package message

const (
	LoginMesType    = "LoginMes"
	LoginResMesType = "LoginResMes"
)

// 客户端与服务器之间传输的消息体
type Message struct {
	Type string `json:"type"` // 消息类型
	Data string `json:"data"` // 消息数据(具体消息数据类型多种， 共同用json序列化后的string表示)
}

// 登陆消息
type LoginMes struct {
	UserId   int    `json:"userId"`
	UserPwd  string `json:"userPwd"`
	UserName string `json:"userName"` // 用户名
}

// 登录返回消息
type LoginResMes struct {
	Code  int    `json:"code"`  // 500=未注册 200=成功
	Error string `json:"error"` // 返回错误信息
}
