package process

import (
	"fmt"
	"go_ChatSw/public"
	"encoding/json"
)

type SmsProcess struct {

}

//发送群聊的消息
func (this *SmsProcess) SendGroupMes(content string) (err error) {


	//1.创建一个Mes
	var mes public.Message
	mes.Type = public.SmsMesType

	//2.创建一个SmsMes
	var smsMes public.SmsMes
	smsMes.Content = content
	smsMes.UserId = CurUser.UserId
	smsMes.UserStatus = CurUser.UserStatus

	//3.序列化 smsMes
	data,err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("json.Marshal(smsMes) error:",err)
		return
	}
	mes.Data = string(data)

	//4.对mes再次序列化
	data,err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal(mes) error:",err)
		return
	}

	//5.将mes发送给服务器
	tf := &public.Transfer {
		Conn:CurUser.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("tf.WritePkg(data) error:",err)
		return
	}

	return
}