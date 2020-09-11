package main

import "fmt"

//定义emp
type Emp struct {
	Id   int
	Name string
	Next *Emp
}

//
func (this *Emp) ShowMe() {
	fmt.Printf("链表 %d 找到该雇员 %d \n", this.Id%7, this.Id)

}

//方法待定 添加员工的方法 编号从小到大
func (this *EmpLink) Insert(emp *Emp) {

	cur := this.Head   //辅助指针
	var pre *Emp = nil //辅助指针 pre 始终在cur 前面
	//如果当前的 EmpLink 是空链表
	if cur == nil {
		this.Head = emp
		return
	}
	//如果不是空链表 给 emp 找到对应的位置并插入
	//思路：让 cur 和emp 比较，让 pre 保持在 cur 前面
	for {
		if cur != nil {
			//比较
			if cur.Id > emp.Id {
				//找到位置
				break
			}
			pre = cur //保证同步
			cur = cur.Next
		} else {
			//没有哪个大于它 emp.Id 为最大
			break
		}
	}
	//退出时 看 是否将 emp 添加到链表最后
	if cur == nil { //最后加入
		pre.Next = emp
		emp.Next = cur
	} else {
		pre.Next = emp
		emp.Next = cur
	}
}

//我们这里的EmpLink 不带表头，即第一个结点就是存放雇员
type EmpLink struct {
	Head *Emp
}

//显示链表信息
func (this *EmpLink) ShowLink(no int) {

	if this.Head == nil {
		fmt.Printf("链表%d 为空\n", no)
		return
	}

	//遍历当前链表 显示数据
	cur := this.Head
	for {
		if cur != nil {
			fmt.Printf("链表%d 雇员id =%d 名字=%s ->", no, cur.Id, cur.Name)
			cur = cur.Next
		} else {
			break
		}
	}
	fmt.Println()
}

//根据id查询 雇员，没有返回 nil
func (this *EmpLink) FindById(id int) *Emp {
	cur := this.Head
	for {
		if cur != nil && cur.Id == id { //找到
			return cur
		} else if cur == nil {
			break
		}
		cur = cur.Next
	}
	return nil
}

//定义hashtable 含有一个链表数组
type HashTable struct {
	LinkArr [7]EmpLink
}

//给 HashTable 写 Insert 雇员的方法
func (this *HashTable) Insert(emp *Emp) {

	//使用散列函数 确定将该雇员添加到哪个链表
	linkNo := this.HashFun(emp.Id)
	//使用对应的链表添加
	this.LinkArr[linkNo].Insert(emp)

}

//编写一个散列方法,将雇员id 传入返回应该存入的链表标号
func (this *HashTable) HashFun(id int) int {
	return id % 7 //得到一个值，就是链表的下标
}

//编写方法 显示哈希表 hashtable 的所有雇员
func (this *HashTable) ShowAll() {

	for i := 0; i < len(this.LinkArr); i++ {
		this.LinkArr[i].ShowLink(i)
	}
}

//查找雇员
func (this *HashTable) FindById(id int) *Emp { //返回 Emp指针
	//使用散列函数 确定该雇员在哪一个链表
	linkNo := this.HashFun(id)
	return this.LinkArr[linkNo].FindById(id) //调用链表的 FindById
}
func main() {

	var hashtable HashTable
	key, name, id := " ", " ", 0
	for {
		fmt.Println("====================雇员系统菜单==================")
		fmt.Println("input 添加雇员")
		fmt.Println("show 显示雇员")
		fmt.Println("find 查找雇员")
		fmt.Println("exit 退出系统")
		fmt.Println("请输入选择")
		fmt.Scanln(&key)
		switch key {
		case "input":
			fmt.Println("请输入雇员ID")
			fmt.Scanln(&id)
			fmt.Println("请输入雇员Name")
			fmt.Scanln(&name)

			emp := &Emp{
				Id:   id,
				Name: name,
			}
			hashtable.Insert(emp)
		case "show":
			hashtable.ShowAll()
		case "find":
			fmt.Println("请输入要查找的雇员ID")
			fmt.Scanln(&id)
			emp := hashtable.FindById(id)
			if emp == nil {
				fmt.Printf("id= %d 的雇员不存在 \n", id)
			} else {
				//编写一个方法 显示雇员信息
				emp.ShowMe()
			}
		case "exit":
		default:
			fmt.Println("输入错误")

		}
	}
}
