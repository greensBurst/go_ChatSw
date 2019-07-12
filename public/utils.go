package public
import (
	"fmt"
	"net"
	"encoding/json"
	"encoding/binary"
)


func ReadPkg(conn net.Conn) (mes Message,err error){
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

func WritePkg(conn net.Conn,data []byte) (err error) {
	
	//先发送一个长度给对方
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf []byte = make([]byte,4,4)
	binary.BigEndian.PutUint32(buf,pkgLen)
	//发送长度
	n,err := conn.Write(buf) //n是发送了多少字节数据
	if n != 4 || err != nil {
		fmt.Println("conn.Write(buf) error:",err)
		return
	}

	//发送data本身
	n,err = conn.Write(data) //n是发送了多少字节数据
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Write(data) error:",err)
		return
	}
	return
}