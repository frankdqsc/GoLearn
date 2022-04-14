package main

import (
	"fmt"
)

func main() {
	var arr[3][3]int
	for i:=0;i<3;i++{
		for j:=0;j<3;j++{
			fmt.Scanln(&arr[i][j])
		}
	}
	fmt.Println(arr)
	for i:=0;i<3;i++{
		for j:=0;j<2;j++{
			arr[i][j],arr[j][i] = arr[j][i],arr[i][j]
		}
	}
	fmt.Println(arr)
}
