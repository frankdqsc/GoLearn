package process

import (
	"fmt"
	"go_code/Learn-Go/project/chatroom/client/model"
	"go_code/Learn-Go/project/chatroom/common/message"
)

//客户端要维护的map
var onlineUsers map[int]*message.User = make(map[int]*message.User, 10)
var CurUser model.CurUser //在用户登录成功后完成对 CurUser 初始化

//在客户端显示当前在线的用户
func outputOnlineUser() {
	//遍历 onlineUsers
	fmt.Println("当前在线用户列表：")
	for id, _ := range onlineUsers {
		//如果不显示自己
		fmt.Println("用户id:\t", id)
	}
}

//编写一个方法 处理返回的 NotifyUserStatusMes
func updateUserStatus(notifyUserStatusMes *message.NotifyUserStatusMes) {

	//适当优化
	user, ok := onlineUsers[notifyUserStatusMes.UserId]
	if !ok { //原来没有
		user = &message.User{
			UserId: notifyUserStatusMes.UserId,
		}
	}
	user.UserStatus = notifyUserStatusMes.Status
	onlineUsers[notifyUserStatusMes.UserId] = user

	outputOnlineUser()
}
