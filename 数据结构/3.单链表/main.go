package main

import "fmt"

//定义一个HeroNode
type HeroNode struct {
	no       int
	name     string
	nickname string
	next     *HeroNode //指向下一个结点
}

//给链表插入一个结点
//编写第1种插入方法 在单链表的最后加入
func InsertHeroNode(head *HeroNode, newHeroNode *HeroNode) {

	//1.先找到该链表最后一个结点
	//2.创建一个辅助结点
	temp := head
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next //让temp不断指向下一个结点
	}
	//3.将 newHeroNode 加入到链表的最后
	temp.next = newHeroNode
}

//给链表插入一个结点
//编写第2种插入方法
//根据 NO 的编号从小到大插入
func InsertHeroNode2(head *HeroNode, newHeroNode *HeroNode) {

	//1.先找到该链表最后一个结点
	//2.创建一个辅助结点
	temp := head
	flag := true
	//让插入的结点的no 和 temp的下一个结点的no比较
	for {
		if temp.next == nil {
			break
		} else if temp.next.no > newHeroNode.no {
			//说明newHeroNode 应该插到 temp后面
			break
		} else if temp.next.no == newHeroNode.no {
			//说明链表中已经有no 就不让插入
			flag = false
			break
		}
		temp = temp.next
	}

	if !flag {
		fmt.Println("已经存在 no=", newHeroNode.no)
	} else {
		newHeroNode.next = temp.next
		temp.next = newHeroNode
	}

}

//删除一个结点
func DelHerNode(head *HeroNode, id int) {
	temp := head

	flag := false
	//找到要删除的结点的no 和temp的下一个结点的no比较
	for {
		if temp.next == nil { //说明到链表的最后
			break
		} else if temp.next.no == id {
			flag = true
			break
		}
		temp = temp.next
	}
	if flag { //找到 删除
		temp.next = temp.next.next
	} else {
		fmt.Println("要删除的结点id不存在")
	}
}

//显示链表的所有结点信息
func ListHeroNode(head *HeroNode) {
	//1.创建一个辅助结点
	temp := head
	//先判断该链表是不是一个空的链表
	if temp.next == nil {
		fmt.Println("空链表")
		return
	}

	//2.遍历这个链表
	for {
		fmt.Printf("[%d ,%s, %s]==>", temp.next.no, temp.next.name, temp.next.nickname)
		//判断是否有结点
		temp = temp.next
		if temp.next == nil {
			break
		}
	}
}

func main() {

	//1.先创建一个头结点
	head := &HeroNode{}

	//2.创建一个新的HeroNode
	hero1 := &HeroNode{
		no:       1,
		name:     "宋江",
		nickname: "及时雨",
	}
	hero2 := &HeroNode{
		no:       2,
		name:     "卢俊义",
		nickname: "玉麒麟",
	}
	hero3 := &HeroNode{
		no:       3,
		name:     "林冲",
		nickname: "豹子头",
	}

	//3.加入
	InsertHeroNode2(head, hero3)
	InsertHeroNode2(head, hero1)
	InsertHeroNode2(head, hero2)

	//4.显示及
	ListHeroNode(head)

	//5.删除
	fmt.Println()
	DelHerNode(head, 2)
	ListHeroNode(head)
}
