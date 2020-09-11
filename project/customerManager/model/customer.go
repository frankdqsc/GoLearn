// customerView.go 【界面】 V 【含customerService】
// 1.可以显示界面
// 2.可以接收用户的输入
// 3.根据用户的输入，调用customerService的方法万成客户的管理【修改，删除，显示等等】它是调用

// customerService【处理 业务逻辑】
// 1.完成对用户的各种操作
// 2.对客户的增删改查
// 3.会声明一个customer的切片

// customer【表示数据】model层
// 1.表示一个客户
// 2.客户的各种字段

package model
import "fmt"

//声明一个Customer结构体 表示一个客户信息→  字段取决于需求分析
type Customer struct{
	Id int
	Name string
	Gender string
	Age int
	Phone string
	Email string
}


//使用工厂模式，返回一个Customer的实例
func NewCustomer (id int,name string,gender string,age int,phone string,email string) Customer {
	return Customer{
		Id : id,
		Name : name,
		Gender : gender,
		Age : age,
		Phone : phone,
		Email : email,
	}
}

//第二种创建customer实例的方法 不带 id
func NewCustomer2 (name string,gender string,age int,phone string,email string) Customer {
	return Customer{
		Name : name,
		Gender : gender,
		Age : age,
		Phone : phone,
		Email : email,
	}
}

//返回用户的信息,格式化的字符串
func (this Customer) GetInfo() string {
	info := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\t",this.Id,this.Name,this.Gender,this.Age,this.Phone,this.Email)
	return info
}