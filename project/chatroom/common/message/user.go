package message

//定义一个用户的结构体
type User struct {
	//确定字段信息
	//为了序列化和反序列化
	//用户信息的json字符串的key和结构体的字段对应的 tag 得一致
	UserId     int    `json:"userId"`
	UserPwd    string `json:"userPwd"`
	UserName   string `json:"userName"`
	UserStatus int    `json:"userStatus"` //用户的状态 在线离线。。
	Sex        string `json:"sex"`
}
