package main

type ListNode struct{
	Val int
	Next *ListNode
}
func reverseList(head *ListNode) *ListNode{
	var newHead, next *ListNode
	for head != nil{
		next = head.Next
		head.Next = newHead
		newHead = head
		head = next
	}
	return newHead
}

func main(){

}
