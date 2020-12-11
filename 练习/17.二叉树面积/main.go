package main

import (
	"container/list"
	"fmt"
)

type TreeNode struct {
	//Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

	root := &TreeNode{}
	left1 := &TreeNode{}
	right1 := &TreeNode{}
	right2 := &TreeNode{}
	right3 := &TreeNode{}
	left2 := &TreeNode{}

	root.Left = left1
	left1.Left = left2
	root.Right = right1
	right1.Right = right2
	right1.Left = right3

	width := widthOfBinaryTree(root)
	depth := maxDepth(root)
	area := width * depth
	fmt.Printf("width %d * depth %d = area %d\n", width, depth, area)
}
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := list.New()
	queue.PushFront(root)
	//记录深度
	depth := 0

	for queue.Len() > 0 {
		//当前层的数量
		count := queue.Len()
		for count > 0 {

			element := queue.Back()
			queue.Remove(element)
			node := element.Value.(*TreeNode)
			if node.Left != nil {
				queue.PushFront(node.Left)
			}

			if node.Right != nil {
				queue.PushFront(node.Right)
			}

			count--
		}
		depth++
	}
	return depth
}

func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	stack := make([]*TreeNode, 0)
	stack = append(stack, root)

	ids := make([]uint64, 1) //记录每个节点的标记
	ids = append(ids, 0)

	var maxWidth uint64 = 1

	for len(stack) > 0 {
		if ids[len(ids)-1]-ids[0]+1 > maxWidth {
			maxWidth = ids[len(ids)-1] - ids[0] + 1
		}
		newStack := make([]*TreeNode, 0)
		newids := make([]uint64, 0)
		for i := range stack {

			if stack[i].Left != nil {
				newStack = append(newStack, stack[i].Left)
				newids = append(newids, ids[i]<<1)
			}

			if stack[i].Right != nil {
				newStack = append(newStack, stack[i].Right)
				newids = append(newids, (ids[i]<<1)+1)
			}

		}
		stack = newStack
		ids = newids

	}
	return int(maxWidth)
}
