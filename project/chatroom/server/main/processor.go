//processor.go 总的处理器 一般放在 main 包里
//1.根据客户端的请求，调用相应的处理器，完成相应的任务
package main

import (
	"fmt"
	"go_code/chatroom/common/message"
	process2 "go_code/chatroom/server/process"
	"go_code/chatroom/server/utils"
	"io"
	"net"
)

//先创建一个 Processor 的结构体
type Processor struct {
	Conn net.Conn
}

//编写一个ServerProcessMes 函数
//功能：根据客户端发送消息种类不同，决定调用哪个函数来处理
func (this *Processor) serverProcessMes(mes *message.Message) (err error) {

	//测试 是否能接收到客户端群发得消息
	fmt.Println("mes=", mes)

	switch mes.Type {
	case message.LoginMesType:
		//处理登录
		//创建一个 UserProcess 实例
		up := &process2.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessLogin(mes)
	case message.RegisterMesType:
		//处理注册
		up := &process2.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessRegister(mes)
	case message.SmsMesType:
		//创建一个 SmsProcess 实例完成转发群聊消息
		SmsProcess := &process2.SmsProcess{}
		SmsProcess.SendGroupMes(mes)
	default:
		fmt.Println("消息类型不存在，无法处理")
	}
	return
}

func (this *Processor) process2() (err error) {
	//循环读客户端发送得消息
	for {
		//这里将读取数据包封住成一个 readkg(),返回Message ,Err
		//创建一个 Transfer 实例完成读包任务
		tf := &utils.Transfer{
			Conn: this.Conn,
		}
		mes, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出，服务器端也退出")
				return err
			} else {
				fmt.Println("readPkg err=", err)
				return err
			}
		}
		err = this.serverProcessMes(&mes)
		if err != nil {
			return err
		}
	}
}
