package main
import (
	"fmt"
	"go_ChatSw/client/process"
)

var userId string
var userPwd string

func main() {
	
	//接收用户输入
	var key string

	//主界面
	loop:
	for {
		fmt.Println("----------欢迎登陆多人聊天系统----------")
		fmt.Println("             1 登录系统")
		fmt.Println("             2 注册用户")
		fmt.Println("             3 退出系统")
		fmt.Print("请选择(1 - 3):")

		fmt.Scanln(&key)
		switch key {
		case "1":
			fmt.Println("登录系统")
			fmt.Print("请输入用户id:")
			fmt.Scanln(&userId)
			fmt.Print("请输入用户密码:")
			fmt.Scanln(&userPwd)
			//完成登录
			//1.创建一个UserProcess实例
			up := &process.UserProcess {}
			up.Login(userId,userPwd)
		case "2":
			fmt.Println("注册用户")
		case "3":
			fmt.Println("退出系统")
			break loop
		default:
			fmt.Println("输入有误")
		}
	}
}