package service

import (
	"fmt"
	"go_code/Learn-Go/project/customerManager/model"
)

//
type CustomerService struct {
	customers []model.Customer
	//声明一个字段，表示当前切片含有多少客户
	//该字段后面还可以作为客户的id+1
	customerNum int
}

//编写一个方法，可以返回 *CustomerService
func NewCustomerService() *CustomerService {
	//为了能够看到有客户在切片中，初始化一个客户
	customerService := &CustomerService{}
	customerService.customerNum = 1
	customer := model.NewCustomer(1, "张三", "男", 20, "112", "zs@souhu.com")
	customerService.customers = append(customerService.customers, customer)
	return customerService
}

//返回客户切片
func (this *CustomerService) List() []model.Customer {
	return this.customers
}

//添加用户到customers切片
func (this *CustomerService) Add(customer model.Customer) bool { //没 *的话是值拷贝，

	//我们确定一个分配id的规则，就是添加的顺序
	this.customerNum++
	customer.Id = this.customerNum
	this.customers = append(this.customers, customer)
	return true
}

//根据id删除用户 从切片中删除
func (this *CustomerService) Delete(id int) bool {
	index := this.FindById(id)
	//如果index == -1 说明没这个用户
	if index == -1 {
		return false
	}
	//如果从切片中删除一个元素
	this.customers = append(this.customers[:index], this.customers[index+1:]...) //通过切片删除元素
	return true
}

//根据id查找客户在切片中对应的下标
func (this *CustomerService) FindById(id int) int {

	index := -1
	//遍历this.customers 切片
	for i := 0; i < len(this.customers); i++ {
		if this.customers[i].Id == id {
			//找到
			index = i
		}
	}
	return index
}

//update用户信息        --------------------------------------未完成 //TODO
func (this *CustomerService) Update(id int) bool {
	index := this.FindById(id)
	//如果index == -1 说明没这个用户
	if index == -1 {
		return false
	}
	//对用户进行修改
	fmt.Println("姓名：")
	name := ""
	fmt.Scanln(&name)

	fmt.Println("性别：")
	gender := ""
	fmt.Scanln(&gender)

	fmt.Println("年龄：")
	age := 0
	fmt.Scanln(&age)

	fmt.Println("电话：")
	phone := ""
	fmt.Scanln(&phone)

	fmt.Println("电邮：")
	email := ""
	fmt.Scanln(&email)
	//customer.Id = this.customerNum
	customer := model.NewCustomer2(name, gender, age, phone, email)
	//调用
	this.customers = append(this.customers, customer)
	return true
}
