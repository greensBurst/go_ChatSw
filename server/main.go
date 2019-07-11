package main
import (
	"fmt"
	"net"
)

//处理和客户端的通讯
func process(conn net.Conn) {
	//这里需要延时关闭conn
	defer conn.Close()
	//循环读取客户端发送的信息
	for {
		buf := make([]byte,8096)
		fmt.Println("等待读取客户端发送的数据")
		n,err := conn.Read(buf[:4])
		if n != 4 || err != nil {
			fmt.Println("conn.Read() error:",err)
			return
		}
		fmt.Println("读到的长度buf:",buf[:4])
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