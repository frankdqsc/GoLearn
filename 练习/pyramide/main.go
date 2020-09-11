package main

import "fmt"

func main() {
	var totalLevel int = 9
	//i 表示层数
	for i:=1;i <=totalLevel;i++{
		//在打印*前先打印空格
		for k := 1;k <= totalLevel - i; k++{
			fmt.Print(" ")
		}
		//j 表示每层打印多少*
		for j :=1; j <=2 * i - 1; j ++{
			fmt.Print("*")
		}
		fmt.Println()
	}
}