package main
import (
	"fmt"
	"net"
	"../public"
	"encoding/json"
	"io"
)


//编写一个函数serverProcessLogin,专门处理登录
func  serverProcessLogin(conn net.Conn,mes *public.Message) (err error) {
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
	err = public.WritePkg(conn,data)
	return

}


//编写一个ServerProcessMes函数
//根据客户端发送消息种类不同，决定调用哪个函数来处理
func serverProcessMes(conn net.Conn,mes *public.Message) (err error) {
	switch mes.Type {
	case public.LoginMesType:
		//处理登录
		err = serverProcessLogin(conn,mes)
	case public.RegisterMesType:
		//处理注册
	default:
		fmt.Println("消息类型不存在，无法处理...")
	}
	return
}


//处理和客户端的通讯
func process(conn net.Conn) {
	//这里需要延时关闭conn
	defer conn.Close()
	//循环读取客户端发送的信息
	for {
		//这里将读取数据包直接封装成一个函数readPkg(),返回Message,Err
		mes,err := public.ReadPkg(conn)
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出，服务端也退出...")
				return
			} else {
				fmt.Println("readPkg() error:",err)
			}
		}
		// fmt.Println("mes:",mes)

		err = serverProcessMes(conn,&mes)
		if err != nil {
			return
		}
	}
}

func main() {
	//提示信息
	fmt.Println("服务器在8889端口监听...")
	listen,err := net.Listen("tcp","0.0.0.0:8889")
	defer listen.Close()
	if(err != nil) {
		fmt.Println("net.Listen() error:",err)
		return
	}

	//监听成功就等待客户端来连接服务器
	for {
		fmt.Println("等待客户端连接...")
		conn,err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept() error:",err)
		}

		//一旦连接成功，则启动一个协程与客户端保持通讯
		go process(conn)
	}
}