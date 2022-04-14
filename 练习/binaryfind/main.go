package main

import (
	"fmt"
)

func BinaryFind(arr *[6]int, leftIndex int, rightIndex int, target int) {
	//判断leftindex 是否大于 rightindex
	if leftIndex > rightIndex {
		fmt.Println("Can't find")
		return
	}
	//find middle index
	middle := (leftIndex + rightIndex) / 2

	if arr[middle] > target {
		BinaryFind(arr, leftIndex, middle-1, target) //递归调用
	} else if arr[middle] < target {
		BinaryFind(arr, middle+1, rightIndex, target)
	} else {
		//find it
		fmt.Println("找到了，下标为:", middle)
	}

}
func main() {
	arr := [6]int{1, 8, 10, 89, 1000, 1234}
	BinaryFind(&arr, 0, len(arr)-1, 89) //二分查找
}
