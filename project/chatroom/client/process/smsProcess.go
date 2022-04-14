package process

import (
	"encoding/json"
	"fmt"
	"go_code/Learn-Go/project/chatroom/client/utils"
	"go_code/Learn-Go/project/chatroom/common/message"
)

type SmsProcess struct {
}

//发送群发消息
func (this *SmsProcess) sendGroupMes(content string) (err error) {

	//1.创建一个 Mes
	var mes message.Message
	mes.Type = message.SmsMesType

	//创建一个 SmsMes 实例
	var smsMes message.SmsMes
	smsMes.UserId = CurUser.UserId
	smsMes.UserStatus = CurUser.UserStatus

	//3，序列化 smsMes
	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("sendGroupMes json.Marshal err= ", err.Error())
		return
	}
	mes.Data = string(data)

	//4.对 mes 再次序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("sendGroupMes json.Marshal err= ", err.Error())
		return
	}

	//5.将 mes 发送给服务器
	tf := &utils.Transfer{
		Conn: CurUser.Conn,
	}

	//6.发送
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("sendGroupMes err= ", err.Error())
	}
	return
}
