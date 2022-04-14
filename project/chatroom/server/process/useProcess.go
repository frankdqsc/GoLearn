// 处理和用户相关的请求 登录 注册 注销 用户列表管理
package process2

import (
	"encoding/json"
	"fmt"
	"go_code/Learn-Go/project/chatroom/common/message"
	"go_code/Learn-Go/project/chatroom/server/model"
	"go_code/Learn-Go/project/chatroom/server/utils"
	"net"
)

type UserProcess struct {
	//分析有哪些字段  考虑将来绑定和关联的方法中有哪些东西
	Conn net.Conn
	//增加一个字段 表示该Conn 是哪个用户
	UserId int
}

//编写通知所有在线用户的方法
//userId 要通知其他在线用户 我上线
func (this *UserProcess) NotifyOtherOnlieUser(userId int) {

	//遍历 onlineUsers,然后一个一个发送 NotifyUserStatusMes
	for id, up := range userMgr.onlineUsers {
		//过滤自己
		if id == userId {
			continue
		}
		//开始通知 【单独写一个方法】
		up.NotifyMeOnline(userId)
	}
}

func (this *UserProcess) NotifyMeOnline(userId int) {

	//组装 NotifyUserStatusMes
	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType

	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserId = userId
	notifyUserStatusMes.Status = message.UserOnline

	//将 notifyUserStatusMes 序列化
	data, err := json.Marshal(notifyUserStatusMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	//将序列化后的 notifyUserStatusMes 赋值给 mes.Data
	mes.Data = string(data)

	//对 mes 再次序列化 准备发送
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	//发送，创建 Transfer 实例 发送
	tf := &utils.Transfer{
		Conn: this.Conn,
	}

	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("NotifyMesOnline err =", err)
		return
	}
}

func (this *UserProcess) ServerProcessRegister(mes *message.Message) (err error) {

	//1.先从mes中取出 mes.Data ,并直接反序列化成RegisterMes
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Println("json.Unmarshal fail err =", err)
		return
	}

	//1.先声明一个 resMes
	var resMes message.Message
	resMes.Type = message.RegisterResMesType
	var registerResMes message.RegisterResMes

	//需要到redis数据库进行验证
	//使用 model.MyUserDao 到 redis验证
	err = model.MyUserDao.Register(&registerMes.User)

	if err != nil {
		if err == model.ERROR_USER_EXISTS {
			registerResMes.Code = 505
			registerResMes.Error = model.ERROR_USER_EXISTS.Error()
		} else {
			registerResMes.Code = 506
			registerResMes.Error = "注册未知错误"
		}
	} else {
		registerResMes.Code = 200
	}

	data, err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Println("json.Marshal fail", err)
		return
	}

	//4.将 data 赋值给 resMes
	resMes.Data = string(data)

	//5.将 resMes 进行序列化 ，准备发送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal fail", err)
		return
	}

	//6.发送 data 将其封装到 writePkg 函数中
	//因为使用了分层（MVC）我们先创建一个 Transfer 实例，然后读取
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	return

}

//编写一个函数 serverProcessLogin 专门处理登录的请求
func (this *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
	//核心代码
	//先从mes中取出 mes.Data ,并直接反序列化成LoginMes
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal fail err =", err)
		return
	}

	//1.先声明一个 resMes
	var resMes message.Message
	resMes.Type = message.LoginResMesType

	//2.再声明一个 LoginResMes  并完成赋值
	var loginResMes message.LoginResMes

	//需要到redis数据库进行验证
	//使用 model.MyUserDao 到 redis验证
	user, err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)

	if err != nil {
		if err == model.ERROR_USER_NOTEXISTS {
			loginResMes.Code = 500
			loginResMes.Error = err.Error()
		} else if err == model.ERROR_USER_PWD {
			loginResMes.Code = 403
			loginResMes.Error = err.Error()
		} else {
			loginResMes.Code = 505
			loginResMes.Error = "服务器内部错误..."
		}
	} else {
		loginResMes.Code = 200
		//这里 因为用户登录成功 把登录成功放入到userMgr中
		//将登录成功的用户userId 赋给this
		this.UserId = loginMes.UserId
		userMgr.AddOnlineUser(this)

		//通知其他的在线用户 我上线
		this.NotifyOtherOnlieUser(loginMes.UserId)
		//将当前在线用户的 id 放入到 loginResMes.UsersId
		//遍历 userMgr.onlineUsers
		for id, _ := range userMgr.onlineUsers {
			loginResMes.UsersId = append(loginResMes.UsersId, id)
		}
		fmt.Println(user, "登录成功")
	}
	//如果用户id= 100 ，密码= 123456认为合法 否则不合法
	// if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
	// 	//合法
	// 	loginResMes.Code = 200

	// } else {
	// 	//不合法
	// 	loginResMes.Code = 500 //表示该用户不存在
	// 	loginResMes.Error = "用户不存在 请注册再使用"
	// }

	//3.将 loginResMes 序列化
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal fail", err)
		return
	}

	//4.将 data 赋值给 resMes
	resMes.Data = string(data)

	//5.将 resMes 进行序列化 ，准备发送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal fail", err)
		return
	}

	//6.发送 data 将其封装到 writePkg 函数中
	//因为使用了分层（MVC）我们先创建一个 Transfer 实例，然后读取
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	return
}
