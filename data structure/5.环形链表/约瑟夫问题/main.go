package main

import "fmt"

//定义孩子的结构体结点
type Kid struct {
	No   int
	Next *Kid //指向下一个孩子的指针
}

//构成环形链表
//*Boy 返回该环形链表的第一个小孩的指针
func AddKid(num int) *Kid {

	first := &Kid{}  //空结点
	curKid := &Kid{} //空结点

	//判断
	if num < 1 {
		fmt.Println("只有一个结点")
		return first
	}

	//循环的构建这个环形链表
	for i := 1; i <= num; i++ {
		kid := &Kid{
			No: i,
		}
		//分析构成环形链表 需要一个辅助指针
		//1.一个小孩时比较特殊
		if i == 1 {
			first = kid
			curKid = kid
			curKid.Next = first //构建环形链表
		} else {
			curKid.Next = kid
			curKid = kid
			curKid.Next = first //构建环形链表
		}
	}
	return first
}

//显示单向的环形链表（=遍历）
func ShowKid(first *Kid) {

	//如果为空链表
	if first.Next == nil {
		fmt.Println("链表为空，没有小孩")
		return
	}

	//至少有一个小孩时，创建一个指针 帮助遍历
	curKid := first
	for {
		fmt.Printf("小孩标号=%d ->", curKid.No)
		//退出条件
		if curKid.Next == first {
			break
		}
		//移动
		curKid = curKid.Next
	}
}

//函数 PlayGame
func PlayGame(first *Kid, startNo int, countNum int) {

	//1.空的链表单独处置
	if first.Next == nil {
		fmt.Println("空的链表 没有kid")
		return
	}

	tail := first
	//3.让tail指向环形链表得最后一个
	for {
		if tail.Next == first { //tail转了到尾
			break
		}
		tail = tail.Next
	}
	//4. 让first 移动到 startNo
	for i := 0; i < startNo-1; i++ {
		first = first.Next
		tail = tail.Next
	}
	fmt.Println()
	//5.数 countNum 然后删除 first 指向的小孩
	for {
		//开始数 countNum -1 次
		for i := 0; i < countNum-1; i++ {
			first = first.Next
			tail = tail.Next
		}
		fmt.Printf("小孩编号为 %d 出圈 ->\n", first.No)
		//删除 first 代表的小孩
		first = first.Next
		tail.Next = first
		//出循环条件
		if tail == first {
			break
		}
	}
	fmt.Printf("小孩编号为 %d 出圈 ->", first.No)
}

func main() {
	addNum := 5
	startNo := 5
	first := AddKid(addNum)
	//显示
	ShowKid(first)
	fmt.Println()
	//startNo 要小于等于小孩的总数
	if startNo > addNum {
		fmt.Println("startNo 要小于等于小孩的总数")
		return
	}
	PlayGame(first, startNo, 3)
}
