package main

import (
	"fmt"
)
func main() {
	var n int
	fmt.Scanln(&n)
	arr := make([][]int,n)
	/*	输入维度后先把里面一维，初始化为 0
		4
		arr: [[0 0 0 0] [0 0 0 0] [0 0 0 0] [0 0 0 0]]
		3
		arr: [[0 0 0] [0 0 0] [0 0 0]]
	*/
	for i := 0; i < n; i++ {
		arr[i] = make([]int, n)
	}
	fmt.Println("arr:",arr)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("请从键盘输入第%d行的第%d个数字：\n", i, j)
			fmt.Scanln(&arr[i][j])
		}
	}
	fmt.Println("arr:",arr)
}
