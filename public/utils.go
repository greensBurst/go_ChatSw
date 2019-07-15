package public
import (
	"fmt"
	"net"
	"encoding/json"
	"encoding/binary"
)

//这里将这些方法关联到结构体中
type Transfer struct {
	//分析他应该有哪些字段
	Conn net.Conn
	Buf [8096]byte //这是传输时，使用缓冲
}


func (this *Transfer) ReadPkg() (mes Message,err error){
	// fmt.Println("等待读取对方发送的数据")
	_,err = this.Conn.Read(this.Buf[:4])
	if err != nil {
		// fmt.Println("conn.Read(1) error:",err)
		return
	}
	
	//根据buf[:4]转成一个uint32
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(this.Buf[:4])

	//根据pkgLen读取消息内容
	n,err := this.Conn.Read(this.Buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Read(2) error:",err)
	}

	//把buf反序列化 -> Message
	err = json.Unmarshal(this.Buf[:pkgLen],&mes)
	if err != nil {
		fmt.Println("json.Unmarshal() error:",err)
		return
	}
	return 
}

func (this *Transfer) WritePkg(data []byte) (err error) {
	
	//先发送一个长度给对方
	var pkgLen uint32
	pkgLen = uint32(len(data))
	binary.BigEndian.PutUint32(this.Buf[0:4],pkgLen)
	//发送长度
	n,err := this.Conn.Write(this.Buf[:4]) //n是发送了多少字节数据
	if n != 4 || err != nil {
		fmt.Println("conn.Write(buf) error:",err)
		return
	}

	//发送data本身
	n,err = this.Conn.Write(data) //n是发送了多少字节数据
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Write(data) error:",err)
		return
	}
	return
}