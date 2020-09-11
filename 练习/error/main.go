package main
import "fmt"

func test() {
	//使用defer + recover 来辅助和处理异常
	defer func() {
		err := recover() //recover()内置函数，可以捕获到异常
		if err != nil {  //说明捕获到异常
			//这里就可以将错误信息发送给管理员
			fmt.Println("有错误，发送邮件给admin@soho.com")
		}
	}()
	num1 :=10
	num2 :=0
	res  := num1 / num2
	fmt.Println("res=",res)
}

func main() {
	test()
	fmt.Println("代码先通过")
}


