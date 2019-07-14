package public

const (
	LoginMesType = "LoginMes"
	LoginResMesType = "LoginResMes"
	RegisterMesType = "RegisterMes"
	RegisterResMesType = "RegisterResMes"
)

//系统中统一的消息传递格式
type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

//登录的消息，发送给服务端的
type LoginMes struct {
	UserId string `json:"userId"`
	UserPwd string `json:"userPwd"`
	UserName string `json:"userName"`
}

//登陆过后服务端返回的消息
type LoginResMes struct {
	Code int  `json:"code"` //状态码，500:用户未注册;200:登录成功
	Error string `json:"error"` //返回的错误信息，没有就是nil
}

type RegisterMes struct {
	User User `json:"user"` //类型就是User结构体
}

type RegisterResMes struct {
	Code int  `json:"code"`  //400用户已占用 200注册成功
	Error string `json:"error"`
}