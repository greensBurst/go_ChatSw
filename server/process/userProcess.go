package process2
import (
	"fmt"
	"net"
	"go_ChatSw/public"
	"encoding/json"
)

type UserProcess struct {
	Conn net.Conn
}


//编写一个函数serverProcessLogin,专门处理登录
func (this *UserProcess) ServerProcessLogin(mes *public.Message) (err error) {
	//核心代码
	//1.先从mes中取出mes.Data，并直接反序列化成LoginMes
	var loginMes public.LoginMes
	err = json.Unmarshal([]byte(mes.Data),&loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal([]byte(mes.Data),&loginMes) error:",err)
		return
	}

	//1先声明一个resMes
	var resMes public.Message
	resMes.Type = public.LoginResMesType

	//2在声明一个LoginResMes，并完成赋值
	var loginResMes public.LoginResMes

	//如果用户的id为100，密码为123456，认为是正确的
	if loginMes.UserId == "100" && loginMes.UserPwd == "123456" {
		//合法
		loginResMes.Code = 200
	} else {
		//不合法
		loginResMes.Code = 500 //500状态码表示用户不存在
		loginResMes.Error = "该用户不存在"
	}

	//3将loginResMes序列化
	data,err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal(loginResMes) error:",err)
		return
	}

	//4将data赋值给resMes
	resMes.Data = string(data)

	//5对resMes序列化，准备发送
	data,err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal(resMes) error:",err)
		return
	}

	//6发送data 我们将其封装到writePkg函数中
	//因为使用分层模式(mvc)，我们先创建一个Tranfer实例，然后读取
	tf := &public.Transfer {
		Conn:this.Conn,
	}
	err = tf.WritePkg(data)
	return

}