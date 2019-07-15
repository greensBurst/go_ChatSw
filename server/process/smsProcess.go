package process2
import (
	"fmt"
	"go_ChatSw/public"
	"encoding/json"
	"net"
)

type SmsProcess struct {

}

//写方法转发消息
func (this *SmsProcess) SendGroupMes(mes *public.Message) {

	//遍历服务器端的onlineUsers map[string]*UserProcess
	//将消息转发出去

	//取出mes的内容 SmsMes
	var smsMes public.SmsMes
	err := json.Unmarshal([]byte(mes.Data),&smsMes)
	if err != nil {
		fmt.Println("json.Unmarshal error:",err)
		return
	}

	data,err := json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal(mes) error:",err)
		return
	}

	for id,up := range userMgr.onlineUsers {

		//这里还需要过滤掉自己
		if id == smsMes.UserId {
			continue
		}
		this.SendMesToEachOnlineUser(data,up.Conn)
	}
}

func (this *SmsProcess) SendMesToEachOnlineUser(data []byte,conn net.Conn) {

	//创建一个Transfer发送data
	tf := &public.Transfer {
		Conn:conn,
	}
	err := tf.WritePkg(data)
	if err != nil {
		fmt.Println("转发消息失败:",err)
	}
}