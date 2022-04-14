// main.go 显示第一级菜单 根据用户的输入去调用相应的处理器

//go build -o server.exe go_code/Learn_Go/project/chatroom/server/main
//go build -o client.exe go_code/Learn_Go/project/chatroom/client/main
package main

import (
	"fmt"
	//"go_code/chatroom/client/process"
	"go_code/Learn-Go/project/chatroom/client/process"
	"os"
)

var userId int
var userPwd string
var userName string

//定义两个全局变量 用户id和密码

func main() {

	//接收用户的选择
	var key int
	//判断是否还继续显示菜单
	//var loop = true
	for true {
		fmt.Println("--------------------欢迎登陆聊天系统--------------------")
		fmt.Println("\t\t\t 1.登录聊天室")
		fmt.Println("\t\t\t 2.注册用户")
		fmt.Println("\t\t\t 3.退出系统")
		fmt.Println("\t\t\t 请选择(1-3):")

		fmt.Scanf("%d\n", &key) // scanf得加\n
		switch key {
		case 1:
			fmt.Println("登录聊天室")
			fmt.Println("请输入用户的Id")
			//fmt.Println(&userId)
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户的密码")
			//fmt.Println(&userPwd)
			fmt.Scanf("%s\n", &userPwd)

			//完成登录
			//1.创建一个 UserProcess 的实例
			up := &process.UserProcess{}
			up.Login(userId, userPwd)
		case 2:
			fmt.Println("注册用户")

			fmt.Println("请输入用户的Id")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户的密码")
			fmt.Scanf("%s\n", &userPwd)
			fmt.Println("请输入用户的名字")
			fmt.Scanf("%s\n", &userName)

			//2.调用UserProcess 完成注册的请求
			up := &process.UserProcess{}
			up.Register(userId, userPwd, userName)
		case 3:
			fmt.Println("退出系统")
			os.Exit(0)
		default:
			fmt.Println("你的输入有误，请重新输入")
		}
	}

	//根据用户的输入，显示新的提示信息
	// if key == 1 {
	// 	//用户登录
	// 	fmt.Println("请输入用户的Id")
	// 	//fmt.Println(&userId)
	// 	fmt.Scanf("%d\n", &userId)

	// 	fmt.Println("请输入用户的密码")
	// 	//fmt.Println(&userPwd)
	// 	fmt.Scanf("%s\n", &userPwd)

	//

	//先把登录的函数写到另一个文件 比如 login.go
	//重新调用
	//login(userId, userPwd)
	// if err != nil {
	// 	fmt.Println("登陆失败")
	// } else {
	// 	fmt.Println("登录成功")
	// }
	// } else if key == 2 {
	// 	fmt.Println("进行用户注册的逻辑。")

	// }
}
