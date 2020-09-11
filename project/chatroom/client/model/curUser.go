package model

import (
	"go_code/chatroom/common/message"
	"net"
)

//客户端很多地方会用到 curUser
type CurUser struct {
	Conn net.Conn
	message.User
}
