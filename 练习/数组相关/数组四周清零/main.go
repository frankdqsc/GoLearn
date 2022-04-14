package main

import (
	"fmt"
)

/*
2222
2222
2222
定义一个3行4列的二维数组，
逐个从键盘输入值，
编写程序将四周的数据清0 */
func main() {
	var arr1 [3][4]int
	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr1[i]); j++ {
			fmt.Printf("请从键盘输入第%d行的第%d个数字：\n", i, j)
			fmt.Scanln(&arr1[i][j])
		}
		fmt.Println(arr1)
	}
	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr1[i]); j++ {
			if i == 1 && j == 1 || i == 1 && j == 2 { } else {
				arr1[i][j] = 0
			}
		}
	}
	fmt.Println(arr1)
}
