package process

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go_code/chatroom/client/utils"
	"go_code/chatroom/common/message"
	"net"
	"os"
)

//关联一个用户登录的结构体
type UserProcess struct {
	//暂时不需要字段
}

//增加一个Register方法 完成注册请求
func (this *UserProcess) Register(userId int, userPwd string, userName string) (err error) {

	//1.连接到服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}
	//延时关闭
	defer conn.Close()

	//2.准备通过conn发送消息给服务
	var mes message.Message
	mes.Type = message.RegisterMesType

	//3.创建一个LoginMes结构体
	var registerMes message.RegisterMes
	registerMes.User.UserId = userId
	registerMes.User.UserPwd = userPwd
	registerMes.User.UserName = userName

	//4.将 registerMes 序列化
	data, err := json.Marshal(registerMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	//5.把data赋给mes.Data字段
	mes.Data = string(data)

	//6.将mes 进行序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	//创建一个 Transfer 实例
	tf := &utils.Transfer{
		Conn: conn,
	}

	//发送 data 给服务器端
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("注册发送信息错误 err = ", err)
	}

	mes, err = tf.ReadPkg() // mes 就是RegisterResMes
	if err != nil {
		fmt.Println("readPkg(conn) err =", err)
		return
	}
	//将 mes 的Data 部分反序列化成 RegisterResMes
	var registerResMes message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data), &registerResMes)
	if registerResMes.Code == 200 {
		fmt.Println("注册成功，请重新登录")
		os.Exit(0)
	} else {
		fmt.Println(registerResMes.Error)
		os.Exit(0)
	}
	return
}

//写一个函数，完成登录
func (this *UserProcess) Login(userId int, userPwd string) (err error) {
	//开始定协议
	// fmt.Printf("userId = %d userPwd = %s\n", userId, userPwd)
	// return nil

	//1.连接到服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}
	//延时关闭
	defer conn.Close()

	//2.准备通过conn发送消息给服务
	var mes message.Message
	mes.Type = message.LoginMesType

	//3.创建一个LoginMes结构体
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd

	//4.将loginMes 序列化
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	//5.把data赋给mes.Data字段
	mes.Data = string(data)

	//6.将mes 进行序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	//7.到这个时候 data就是我们要发送得消息
	//7，1 先把data的长度发送给服务器
	//先获取到 data的长度 转成一个表示长度byte的切片
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	//发送长度
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(buf) fail", err)
		return
	}

	fmt.Printf("客户端，发送信息的长度=%d 内容=%s", len(data), string(data))

	//发送消息本身
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write(data) fail", err)
		return
	}

	//休眠20
	// time.Sleep(20 * time.Second)
	// fmt.Println("休眠了20")
	//这里还需要处理服务器端返回的消息
	//创建一个 Transfer 实例
	tf := &utils.Transfer{
		Conn: conn,
	}
	mes, err = tf.ReadPkg() //mes

	if err != nil {
		fmt.Println("readPkg(conn) err =", err)
		//fmt.Println("mes =", mes)
		return
	}

	//将 mes 的Data 部分反序列化成 LoginResMes
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	//fmt.Println("mes.Data =", mes)
	if loginResMes.Code == 200 {
		//初始化 CurUser
		CurUser.Conn = conn
		CurUser.UserId = userId
		CurUser.UserStatus = message.UserOnline

		//fmt.Println("登录成功")

		//显示当前在线用户的列表 遍历 loginResMes.UsersId
		fmt.Println("当前用户列表如下：")
		for _, v := range loginResMes.UsersId {
			//要求不显示自己在线
			if v == userId {
				continue
			}
			fmt.Println("用户id:\t", v)
			//完成 客户端的 onlineUsers 初始化
			user := &message.User{
				UserId:     v,
				UserStatus: message.UserOnline,
			}
			onlineUsers[v] = user
		}
		fmt.Print("\n\n")

		//这里还需要再客户端启动该一个协程 监控服务器给客户端的发送
		//该协程保持和服务器的通讯，接收服务器的推送消息并显示在客户端
		go serverProcessMes(conn)

		//1.显示登录成功后的菜单  循环显示
		for {
			ShowMenu()
		}
	} else {
		fmt.Println(loginResMes.Error)
	}
	return
}
