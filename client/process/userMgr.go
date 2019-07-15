package process
import (
	"fmt"
	"go_ChatSw/public"
	"go_ChatSw/client/model"
)

//客户端要维护的map
var onlineUsers map[string]*public.User = make(map[string]*public.User,10)
var CurUser model.CurUser  //在用户登录成功后完成初始化

//在客户端显示当前在线的用户
func outputOnlineUser() {
	
	fmt.Println("当前在线用户列表:")
	for id,_ := range onlineUsers {
		fmt.Println("用户id:\t",id)
	}
}


func updateUserStatus(notifyUserStatusMes *public.NotifyUserStatusMes) {

	user,ok := onlineUsers[notifyUserStatusMes.UserId]
	if !ok {
		user = &public.User {
			UserId:notifyUserStatusMes.UserId,
		}
	}

	user.UserStatus = notifyUserStatusMes.Status
	onlineUsers[notifyUserStatusMes.UserId] = user
	
	outputOnlineUser()
}
