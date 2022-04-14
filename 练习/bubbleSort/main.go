package main

import (
	"fmt"
)

//数组为值传递 切片为引用传递
//冒泡排序
func BubbleSort(arr [5]int) {

	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < 4; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println("排序后arr=", (arr))
}
func main() {
	//定义数组
	arr := [5]int{24, 69, 80, 57, 13}
	//将数组传递给一个函数，完成排序

	BubbleSort(arr)
}
