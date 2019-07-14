package main
import (
	"fmt"
	"net"
	"go_ChatSw/public"
	"go_ChatSw/server/process"
	"io"
)

//先创建一个Processor结构体
type Processor struct {
	Conn net.Conn
}

//编写一个ServerProcessMes函数
//根据客户端发送消息种类不同，决定调用哪个函数来处理
func (this *Processor) serverProcessMes(mes *public.Message) (err error) {
	switch mes.Type {
	case public.LoginMesType:
		//处理登录
		//创建一个UserProcess实例
		up := &process2.UserProcess{
			Conn:this.Conn,
		}
		err = up.ServerProcessLogin(mes)
	case public.RegisterMesType:
		//处理注册
		up := &process2.UserProcess{
			Conn:this.Conn,
		}
		err = up.ServerProcessRegister(mes)
	default:
		fmt.Println("消息类型不存在，无法处理...")
	}
	return
}

func (this *Processor) process2() (err error) {
	//循环读取客户端发送的信息
	for {
		//这里将读取数据包直接封装成一个函数readPkg(),返回Message,Err
		//创建一个Transfer实例，完成读报任务
		tf := &public.Transfer{
			Conn:this.Conn,
		}
		mes,err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出，服务端也退出...")
				return err
			} else {
				fmt.Println("readPkg() error:",err)
				return err
			}
		}
		// fmt.Println("mes:",mes)

		err = this.serverProcessMes(&mes)
		if err != nil {
			return err
		}
	}
}