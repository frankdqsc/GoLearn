/*
123
456
789
→
147
258
369 */
package main

import (
	"fmt"
)

type MethodUtils struct {
}

func (am MethodUtils) Trans(arr [3][3]int) {
	//var tem int
	//通过中间变量交换下标
	for i := 1; i < 3; i++ {
		for j := 0; j < 2; j++ { //只上三角交换  i到 0 1 2，j到 0 1
			// tem = arr[i][j]
			// fmt.Println("tem=", tem)
			// arr[i][j] = arr[j][i]
			// arr[j][i] = tem
			arr[i][j], arr[j][i] = arr[j][i], arr[i][j] //go交换特性
		}
	}
	fmt.Println("arr=", arr)
}
func main() {
	var am MethodUtils
	var arr [3][3]int
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Printf("请从键盘输入第%d行的第%d个数字：\n", i, j)
			fmt.Scanln(&arr[i][j])
		}
	}
	fmt.Println("arr=", arr)
	am.Trans(arr)
}
