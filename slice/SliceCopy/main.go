package main

import "fmt"

func f(slice[]int, i int) []int{
	copy(slice[i:],slice[i+1:])  //789 89 -> 899  -> 46899
	fmt.Println(slice)
	return slice[:len(slice)-1]
}

func main() {

	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3}

	//copy(slice2, slice1) // 只会复制slice1的前3个元素到slice2中
	fmt.Println(copy(slice1, slice2)) // 只会复制slice2的3个元素到slice1的前3个位置
	fmt.Println(slice1)
	//fmt.Println(slice2)

	s:= []int{4,6,7,8,9}
	fmt.Println(f(s,2))
}
