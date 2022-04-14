package main

import "fmt"

//数组为值传递 切片为引用传递
//冒泡排序
// func BubbleSort(slice [5]int) {
// 	for i := 0; i < len(slice)-1; i++ {
// 		for j := 0; j < 4; j++ {
// 			if (slice)[j] > (slice)[j+1] {
// 				(slice)[j], (slice)[j+1] = (slice)[j+1], (slice)[j]
// 			}
// 		}
// 	}
// 	fmt.Println("排序后arr=", (slice))
// }
// func main() {
// 	//定义数组
// 	slice := [5]int{24, 69, 80, 57, 13}
// 	//将数组传递给一个函数，完成排序
// 	BubbleSort(slice)
// }

func BubbleSort(slice []int, k int) []int{
	for i := 0; i < len(slice)-1; i++ {
		for j := 0; j < len(slice)-1; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
	slice = slice[:k]
	return slice
}
func main() {

	slice := make([]int, 0, 10)
	var n, k int   //第k大/小 数
	for i := 0; i < 4; i++ {
		fmt.Scan(&n)
		slice = append(slice, n)
	}
	fmt.Scanln(&k)

	slice = BubbleSort(slice, k)
	fmt.Println(slice)
}
