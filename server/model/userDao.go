package model
import (
	"fmt"
	"redigo/redis"
	"encoding/json"
)

//我们在服务器启动后，就初始化一个userDao实例
//把他做成全局的变量，在需要和redis操作时，就直接使用即可
var (
	MyUserDao *UserDao
)

//定义一个UserDao的结构体
//完成对User结构体的各种操作

type UserDao struct {
	pool *redis.Pool
}

//使用工厂模式，创建一个UserDao实例
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {

	userDao = &UserDao {
		pool:pool,
	}
	return
}



//1.根据一个用户id返回一个User实例 + error
func (this *UserDao) getUserById(conn redis.Conn,id string) (user *User,err error) {

	//通过给定的id取redis查询这个用户
	res,err := redis.String(conn.Do("HGet","users",id))
	if err != nil {
		if err == redis.ErrNil { //表示在user Hash中没有找到对应的id
			err = ERROR_USER_NOTEXISTS
		}
		return
	}

	user = &User{}

	//这里需要把res反序列化成一个user实例
	err = json.Unmarshal([]byte(res),user)
	if err != nil {
		fmt.Println("json.Unmarshal([]byte(res),user) error:",err)
		return
	}

	return
}

//完成登录的校验
//1.Login 完成对用户的验证
//2.如果用户的id和pwd都正确，则返回一个user实例
//3.如果id或者pwd有错，返回对应的错误信息
func (this *UserDao) Login(userId string,userPwd string) (user *User,err error) {

	//先从userDao的连接池中取出一根连接
	conn := this.pool.Get()
	defer conn.Close()
	user,err = this.getUserById(conn,userId)
	if err != nil {
		return
	}
	//这是证明用户获取到了
	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
		return
	}
	return
}
