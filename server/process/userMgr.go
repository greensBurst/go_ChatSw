package process2
import (
	"fmt"
)

//UserMgr在服务器端有且只有一个
//在很多地方都会使用到，因此做成全局变量
var (
	userMgr *UserMgr
)

type UserMgr struct {
	onlineUsers map[string]*UserProcess
}

//完成对userMgr的初始化工作
func init() {
	userMgr = &UserMgr {
		onlineUsers : make(map[string]*UserProcess,1024),
	}
}

//完成对onlineUsers添加
func (this *UserMgr) AddOnlineUser(up *UserProcess) {
	this.onlineUsers[up.UserId] = up
}

//删除
func (this *UserMgr) DelOnlineUser(userId string) {
	delete(this.onlineUsers,userId)
}

//返回所有在线用户
func (this *UserMgr) GetOnlineUser() map[string]*UserProcess {
	return this.onlineUsers
}

//根据id返回对应的值
func (this *UserMgr) GetOnlineUserById(userId string) (up *UserProcess,err error) {
	up,ok := this.onlineUsers[userId]
	if !ok { //说明查找的这个用户不在线
		err = fmt.Errorf("用户%s不存在",userId)
		return
	}
	return
}