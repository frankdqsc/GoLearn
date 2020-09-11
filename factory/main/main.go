package main 
import (
	"fmt"
	"go_code/project01/factory/model"
)

func main(){
	var stu = model.NewStudent("tom",80)

	fmt.Println(*stu)
	fmt.Println("name=",stu.Name ,"score=", stu.Score)
}