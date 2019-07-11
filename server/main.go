package main
import (
	"fmt"
	"net"
	"../public"
	"encoding/binary"
	"encoding/json"
	"io"
)

func readPkg(conn net.Conn) (mes public.Message,err error){
	buf := make([]byte,8096)
	fmt.Println("等待读取客户端发送的数据")
	_,err = conn.Read(buf[:4])
	if err != nil {
		fmt.Println("conn.Read(1) error:",err)
		return
	}
	
	//根据buf[:4]转成一个uint32
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[:4])

	//根据pkgLen读取消息内容
	n,err := conn.Read(buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Read(2) error:",err)
	}

	//把buf反序列化 -> Message
	err = json.Unmarshal(buf[:pkgLen],&mes)
	if err != nil {
		fmt.Println("json.Unmarshal() error:",err)
		return
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
		mes,err := readPkg(conn)
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出，服务端也退出...")
				return
			} else {
				fmt.Println("readPkg() error:",err)
			}
		}
		fmt.Println("mes:",mes)
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