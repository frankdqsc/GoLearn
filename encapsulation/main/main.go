package main
import (
	"fmt"
	"go_code/project01/encapsulation/model"
)

func main (){
	//创建一个account变量
	account := model.NewAccount("fgf23232","121346",50)
	if account != nil{
		fmt.Println("创建成功=",account)
	}else{
		fmt.Println("创建失败")
	}
}