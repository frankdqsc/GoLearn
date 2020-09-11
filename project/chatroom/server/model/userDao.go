package model

import (
	"encoding/json"
	"fmt"
	"go_code/chatroom/common/message"

	"github.com/garyburd/redigo/redis"
)

//在服务器启动后 初始化一个userDao实例
//把它做成全局变量 在需要和redis操作时 直接使用
var (
	MyUserDao *UserDao
)

//定义一个 UserDao 结构体
//完成对 User 结构体的各种操作
type UserDao struct {
	pool *redis.Pool
}

//使用工厂模式 创建一个UserDao 实例
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {

	userDao = &UserDao{
		pool: pool,
	}
	return
}

//思考在UserDao中有哪些方法 验证id和 返回实例和error
func (this *UserDao) getUserById(conn redis.Conn, id int) (user *User, err error) {

	//通过给定的id redis查询这个用户
	res, err := redis.String(conn.Do("HGet", "users", id))
	if err != nil {
		//错误
		if err == redis.ErrNil { //表示在 users 哈希中 没有找到对应的id
			err = ERROR_USER_NOTEXISTS
		}
		return
	}

	user = &User{}

	//这里需要把 res 反序列化成User实例
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("json.Unmarshal err =", err)
		return
	}
	return
}

//完成登录的校验
//1.Login 完成对用户的校验
//2.如果用户的id和pwd都正确，则返回一个user实例
//3.如果用户的id和pwd有错误，则返回对应的错误信息
func (this *UserDao) Login(userId int, userPwd string) (user *User, err error) {

	//先从UserDao 的连接中取出一个连接
	conn := this.pool.Get()
	defer conn.Close()
	_, err = this.getUserById(conn, userId)
	if err != nil {
		return
	}
	//证明用户是获取到的
	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
		return
	}
	return
}

func (this *UserDao) Register(user *message.User) (err error) {

	//先从UserDao 的连接中取出一个连接
	conn := this.pool.Get()
	defer conn.Close()
	_, err = this.getUserById(conn, user.UserId)
	if err == nil {
		err = ERROR_USER_EXISTS
		return
	}
	//说明该ID还没注册过
	data, err := json.Marshal(user) //序列化   结构以 → 字符串
	if err != nil {
		return
	}

	//入库
	_, err = conn.Do("HSet", "users", user.UserId, string(data))
	if err != nil {
		fmt.Println("保存注册用户错误 err =", err)
		return
	}
	return
}
