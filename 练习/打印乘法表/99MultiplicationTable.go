//打印9 * 9 乘法表
package main

import (
	"fmt"
)

type MethodUtils struct {
}

func (excel MethodUtils) Print(num int) {
	for i := 1; i < num+1; i++ {
		for j := 1; j < i+1; j++ {
			//fmt.Printf("i*j=v%",i*j)
			fmt.Printf("%d*%d=%d ", j, i, i*j)
		}
		fmt.Println()
	}
}
func main() {
	var excel MethodUtils
	excel.Print(9)
}
