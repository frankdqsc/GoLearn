package process

import (
	"encoding/json"
	"fmt"
	"go_code/chatroom/client/utils"
	"go_code/chatroom/common/message"
	"net"
	"os"
)

//显示登录成功的界面。。
func ShowMenu() {
	fmt.Println("-----------登录成功-----------")
	fmt.Println("-----------1.显示在线用户列表------------")
	fmt.Println("-----------2.发送消息------------")
	fmt.Println("-----------3.信息列表------------")
	fmt.Println("-----------4.退出系统------------")
	fmt.Println("请选择(1-4)：")
	var key int
	var content string

	//因为多出使用 SmsProcess 实例，因此将其定义在 switch 外部
	smsProcess := &SmsProcess{}

	fmt.Scanf("%d\n", &key)
	switch key {
	case 1:
		//fmt.Println("显示在线用户列表")
		outputOnlineUser()
	case 2:
		fmt.Println("你想对大家说什么")
		fmt.Scanf("%s\n", &content)
		smsProcess.sendGroupMes(content)
	case 3:
		fmt.Println("信息列表")
	case 4:
		fmt.Println("你选则了退出系统。。")
		os.Exit(0)
	default:
		fmt.Println("你输入的选项有误")
	}
}

//和服务器端保持通讯
func serverProcessMes(conn net.Conn) {

	//创建一个Transfer 实例 不停的读取服务器发送的消息
	tf := &utils.Transfer{
		Conn: conn,
	}
	for {
		fmt.Println("客户端正在等待读取服务器发送得消息")
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("tf.ReadPkg err =", err)
			return
		}

		//如果读取到消息 进行下一步处理
		//fmt.Printf("mes=%v\n", mes)
		switch mes.Type {

		case message.NotifyUserStatusMesType: //有人上线了
			//处理
			//1.取出 NotifyUserStatueMes
			var notifyUserStatueMes message.NotifyUserStatusMes
			json.Unmarshal([]byte(mes.Data), &notifyUserStatueMes)
			//2.把这个人加入到 客户端 map[int]User 中
			updateUserStatus(&notifyUserStatueMes)
		case message.SmsMesType: //有人群发消息
			outputGroupMes(&mes)
		default:
			fmt.Println("服务器返回未知消息类型")
		}
	}
}
