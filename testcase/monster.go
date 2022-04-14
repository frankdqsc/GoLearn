package monster

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)
type Monster struct{
	Name string
	Age int
	Skill string
}

//给Monster绑定Store,可以将一个Monster变量（对象），序列化后保存到文件中

func (this *Monster)Store() bool{
	//先序列化
	data,err := json.Marshal(this)
	if err != nil{
		fmt.Println("marshal err=",err)
		return false
	}
	//保存到文件
	filePath := "d:/monster.ser"
	
	err = ioutil.WriteFile(filePath,data,0666)
	if err != nil{
		fmt.Println("保存文件出错 =",err)
		return false
	}
	return true
}

//将Monster绑定方法Restore，可以将一个序列化的Monster，从文件中读取
//并反序列化为Monster对象，检查反序列化，名字得正确
func (this *Monster)Restore() bool{
	//1.先从文件中，读取序列化的字符串
	filePath := "d:/monster.ser"
	data,err := ioutil.ReadFile(filePath)
	if err != nil{
		fmt.Println("读文件错误=",err)
		return false
	}

	//2.使用读取到的data[]byte ,对反序列化
	err = json.Unmarshal(data,this)
	if err != nil{
		fmt.Println("反序列化错误=",err)
		return false
	}
	return true
}