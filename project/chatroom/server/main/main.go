package main

import (
	"fmt"
	"go_code/chatroom/server/model"
	"net"
	"time"
)

// func readPkg(conn net.Conn) (mes message.Message, err error) {

// 	buf := make([]byte, 8096)
// 	fmt.Println("读取客户端发送的数据")
// 	_, err = conn.Read(buf[:4])
// 	if err != nil {
// 		//fmt.Println("conn.Read err=", err)
// 		//err = errors.New("read pkg header error") //自定义error
// 		return
// 	}
// 	//根据读到的buf[:4]长度 转成一个uint32类型
// 	var pkgLen uint32
// 	pkgLen = binary.BigEndian.Uint32(buf[0:4])

// 	//根据 pkgLen 读取内容
// 	n, err := conn.Read(buf[:pkgLen]) //把从conn独到的[:pkgLen]这个长度放到buf中，并检验读取的长度
// 	if n != int(pkgLen) || err != nil {
// 		//fmt.Println("conn.Read fail err=", err)
// 		//err = errors.New("read pkg body error")
// 		return
// 	}

// 	//把pkgLen反序列化成 → message.Message
// 	err = json.Unmarshal(buf[:pkgLen], &mes) //&mes
// 	if err != nil {
// 		fmt.Println("json.Unmarshal err=", err)
// 		return
// 	}

// 	return
// }

// func writePkg(conn net.Conn ,data []byte) (err error){

// 	//先发送一个长度给对方
// 	var pkgLen uint32
// 	pkgLen = uint32(len(data))
// 	var buf [4]byte
// 	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
// 	//发送长度
// 	n, err := conn.Write(buf[:4])
// 	if n != 4 || err != nil {
// 		fmt.Println("conn.Write(buf) fail", err)
// 		return
// 	}

// 	//发送 data 本身
// 	n,err = conn.Write(data)
// 	if n!= int(pkgLen) || err != nil{
// 		fmt.Println("conn.Write(bytes) fail", err)
// 		return
// 	}
// 	return

// }

// //编写一个函数 serverProcessLogin 专门处理登录的请求
// func serverProcessLogin(onn net.Conn,mes *message.Message)(err error){
// 	//核心代码
// 	//先从mes中取出 mes.Data ,并直接反序列化成LoginMes
// 	var loginMes message.LoginMer
// 	err = json>unmarshal([]byte(mes.Data),&loginMes)
// 	if err != nil{
// 		fmt.Println("json.Unmarshal fail err =",err)
// 		return
// 	}

// 	//1.先声明一个 resMes
// 	var resMes message.Message
// 	resMes.Type = message.LoginResMesType

// 	//2.再声明一个 LoginResMes  并完成赋值
// 	var loginResMes message.LoginResMes

// 	//如果用户id= 100 ，密码= 123456认为合法 否则不合法
// 	if loginMes.UserId == 100 && loginMes.UserPwd =="123456"{
// 		//合法
// 		loginResMes.Code = 200

// 	} else{
// 		//不合法
// 		loginResMes.Code = 500  //表示该用户不存在
// 		loginResMes.Error = "用户不存在 请注册再使用"
// 	}

// 	//3.将 loginResMes 序列化
// 	data,err := json.Marshal(loginResMes)
// 	if err != nil {
// 		fmt.Println("json.Marshal fail",err)
// 		return
// 	}

// 	//4.将 data 赋值给 resMes
// 	resMes.Data = string(data)

// 	//5.将 resMes 进行序列化 ，准备发送
// 	data,err = json.Marshal(resMes)
// 	if err != nil {
// 		fmt.Println("json.Marshal fail",err)
// 		return
// 	}

// 	//6.发送 data 将其封装到 writePkg 函数中
// 	err = writePkg(conn,data)
// 	return
// }

// //编写一个ServerProcessMes 函数
// //功能：根据客户端发送消息种类不同，决定调用哪个函数来处理
// func serverProcessMes(conn net.Conn,mes *message.Message)(err error){

// 	switch mes.Type{
// 		case message.LoginMesType :
// 			//处理登录
// 			err = serverProcessLogin(conn ,mes)
// 		case message.RegisterMesType :
// 			//处理注册
// 		default :
// 			fmt.Println("消息类型不存在，无法处理")
// 	}
// 	return
// }

//处理和客户端的通讯
func process(conn net.Conn) {
	//读客户端发送的信息
	defer conn.Close()
	//这里调用总控，创建一个总控
	processor := &Processor{
		Conn: conn,
	}
	err := processor.process2()
	if err != nil {
		fmt.Println("客户端和服务器通讯协程错误=err", err)
		return
	}
}

//这里编写一个函数 完成对 UserDao 的初始化任务M
func initUserDao() {
	//这里的pool 本身就是一个全局的变量
	//注意初始化顺序 initPool在 initUserDao 之前
	model.MyUserDao = model.NewUserDao(pool)
}

func main() {

	//当服务器启动时 初始化redis连接池
	initPool("localhost:6379", 16, 0, 300*time.Second)
	initUserDao()
	//提示信息
	fmt.Println("新 服务器在8889端口监听")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	defer listen.Close()
	if err != nil {
		fmt.Println("net.Listen err=", err)
		return
	}
	//一旦监听成功，就等待客户端来连接服务器
	for {
		fmt.Println("等待客户端来连接服务器")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err=", err)
			return
		}
		//一旦连接成功 则启动一个协程和客户端保持通讯
		go process(conn)
	}
}
