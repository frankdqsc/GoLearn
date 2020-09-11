package main
import (
	"fmt"
	"go_code/customerManager/service"
	"go_code/customerManager/model"
)

type customerView struct{
	//定义必要的字段
	key string //接收用户的输入
	loop bool  //表示是否循环的显示菜单
	//增加一个字段customerService
	customerService *service.CustomerService
}

//显示所有客户信息
func (this *customerView) list (){

	//首先，获取到当前所有的客户信息（在切片中）
	customers := this.customerService.List()
	//显示
	fmt.Println("------------------客户列表-------------------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for i:= 0; i< len(customers); i++ {
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Println("------------------客户列表完成----------------------")
}


//得到用户输入。创建新客户，完成添加
func (this *customerView) add(){
	fmt.Println("-------------------添加客户-------------------------")
	fmt.Println("姓名：")
	name :=""
	fmt.Scanln(&name)

	fmt.Println("性别：")
	gender :=""
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

	//构建一个新的customer实例
	// id由系统分派
	customer := model.NewCustomer2(name,gender,age,phone,email)
	//调用
	if this.customerService.Add(customer){
		fmt.Println("------------------添加完成-------------------------")
	}else{
		fmt.Println("------------------添加失败-------------------------")
	}
} 

//得到用户的id，删除该id对应的用户
func (this *customerView) delete(){

	fmt.Println("-----------------------删除用户-------------------------")
	fmt.Println("请选择待删除用户编号（-1退出）：")
	id := -1
	fmt.Scanln(&id)
	if id == -1{
		return  //放弃删除操作
	}

	fmt.Println("确认是否删除（Y/N）:")

	//这里进入循环判断 知道输入Y或者N才退出
	choice := ""
	fmt.Scanln(&choice)
	if choice == "y" || choice =="Y"{
		
		//调用customerService的delete方法
		if this.customerService.Delete(id){
			fmt.Println("-----------------------删除完成-------------------------")
		}else{
			fmt.Println("-----------------------删除失败 id不存在-------------------------")
		}
	}
}
//退出软件
func (this *customerView) exit(){
	fmt.Println("请确认是否退出（Y/N）:")
	for {
		fmt.Scanln(&this.key)
		if this.key == "Y" || this.key =="y" ||this.key =="N" ||this.key == "n"{
			break
		}
		fmt.Println("你的输入有误，确认是否退出（Y/N）:")

	}
	if this.key == "Y" || this.key =="y"{
		this.loop = false
	}
}

//得到用户的id，修改该id对应的用户
func (this *customerView) update(){

	fmt.Println("-----------------------修改用户-------------------------")
	fmt.Println("请选择待修改用户编号（-1退出）：")
	id := -1
	fmt.Scanln(&id)
	if id == -1{
		return  //放弃删除操作
	}

	fmt.Println("确认是否修改（Y/N）:")

	//这里进入循环判断 知道输入Y或者N才退出
	choice := ""
	fmt.Scanln(&choice)
	if choice == "y" || choice =="Y"{
		
		//调用customerService的update方法
		if this.customerService.Update(id){
			fmt.Println("-----------------------修改完成-------------------------")
		}else{
			fmt.Println("-----------------------修改失败 id不存在-------------------------")
		}
	}
}

//显示主菜单
func (this *customerView) mainMenu(){
	for{
		fmt.Println("------------------客户信息管理软件---------------------")
		fmt.Println("		 1.添加客户")
		fmt.Println("		 2.修改客户")
		fmt.Println("		 3.删除客户")
		fmt.Println("		 4.客户列表")
		fmt.Println("		 5.退    出")
		fmt.Println("请选择（1-5）：")

		fmt.Scanln(&this.key)
		switch this.key {
			case "1":
				this.add()
			case "2":
				this.update()
			case "3":
				this.delete()
			case "4":
				this.list()
			case "5":
				this.exit()
				default :
					fmt.Println("输入有误，请重新输入")
		}

		if !this.loop{
			break
		}
	}

	fmt.Println("退出了系统")
}

func main(){
	//fmt.Println("ok")
	//在主函数中创建customerView，并运行显示主菜单
	customerView := customerView{
		key :"",
		loop : true,
	}
	
	//这里对customerView 结构体的customerService 字段初始化
	customerView.customerService = service.NewCustomerService()
	//显示主菜单
	customerView.mainMenu();
}