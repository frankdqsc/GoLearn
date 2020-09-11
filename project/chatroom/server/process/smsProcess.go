package process2

import (
	"encoding/json"
	"fmt"
	"go_code/chatroom/common/message"
	"go_code/chatroom/server/utils"
	"net"
)

// 处理和短消息相关的请求 群聊 点对点聊天

type SmsProcess struct {
	//..暂时不需要字段
}

//写转发消息
func (this *SmsProcess) SendGroupMes(mes *message.Message) {

	//遍历服务器端的 onlineUsers map[int]*UserProcess
	//将消息转发出去

	//取出 mes得内容 SmsMes
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}

	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}

	for id, up := range userMgr.onlineUsers {
		//这里需要过滤掉自己 即不要转发给自己
		if id == smsMes.UserId {
			continue
		}
		this.SendMesToEachOnlineUser(data, up.Conn)
	}
}

func (this *SmsProcess) SendMesToEachOnlineUser(data []byte, conn net.Conn) {

	//创建一个Transfer 实例 发送 data
	tf := &utils.Transfer{
		Conn: conn,
	}
	err := tf.WritePkg(data)
	if err != nil {
		fmt.Println("转发消息失败 err= ", err)
	}
}
