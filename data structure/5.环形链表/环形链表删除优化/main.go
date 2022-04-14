package main

import "fmt"

//定义猫的结构体结点
type CatNode struct {
	no   int
	name string
	next *CatNode
}

func InsertCatNode(head *CatNode, newCatNode *CatNode) {

	//判断是不是第一只猫
	if head.next == nil {
		head.no = newCatNode.no
		head.name = newCatNode.name
		head.next = head //形成一个环形
		fmt.Println(newCatNode, "加入到环形链表")
		return
	}

	//定义一个临时变量 为找到环形的最后结点
	temp := head
	for {
		if temp.next == head {
			break
		}
		temp = temp.next
	}

	//加入到链表中
	temp.next = newCatNode
	newCatNode.next = head
}

//输出环形链表
func ListCircleLink(head *CatNode) {
	fmt.Println("环形链表如下：")
	temp := head
	if temp.next == nil {
		fmt.Println("空环形链表")
		return
	}
	for {
		fmt.Printf("猫的信息为:[id=%d name=%s] ->\n", temp.no, temp.name)
		if temp.next == head {
			break
		}
		temp = temp.next
	}

}

//删除一只猫
func DelCatNode(head *CatNode, id int) *CatNode {

	temp := head
	helper := head

	if temp.next == nil {
		fmt.Println("这是一个空的环形链表，不能删除")
		return head
	}

	//如果只有一个结点
	if temp.next == head {
		temp.next = nil //将头结点的next置空
		return head
	}

	//将 helper 定位到链表得最后
	for {
		if helper.next == head {
			break
		}
		helper = helper.next
	}

	//如果包含两个以上的结点
	flag := true
	for {
		if temp.no == id {
			//找到 直接删除
			helper.next = temp.next
			temp = temp.next //退出循环之前将 temp 和 head 的位置进行后移，否则输出链表 temp.next == head 无法完成判断
			head = head.next
			fmt.Printf("要删除的猫=%d\n", id)
			flag = false
			break
		}
		temp = temp.next     //移动 【作用：比较】
		helper = helper.next //移动 【作用：辅助删除】  helper 总是指向 temp 的前一个结点 为了删除 temp 直接 helper.next = temp.next

		//fmt.Println(temp)
		if temp.next == head { //循环终止条件
			break
		}
	}

	if flag {
		fmt.Printf("没有这只猫no=%d\n", id)
	}
	return head
}

func main() {
	//初始化一个环形链表的头结点
	head := &CatNode{}

	//创建一只猫
	cat1 := &CatNode{
		no:   1,
		name: "tom",
	}
	cat2 := &CatNode{
		no:   2,
		name: "cim",
	}
	cat3 := &CatNode{
		no:   3,
		name: "lilly",
	}
	cat4 := &CatNode{
		no:   4,
		name: "tim",
	}
	InsertCatNode(head, cat1)
	InsertCatNode(head, cat2)
	InsertCatNode(head, cat3)
	InsertCatNode(head, cat4)
	ListCircleLink(head)

	fmt.Println()
	head = DelCatNode(head, 10)
	ListCircleLink(head)

}
