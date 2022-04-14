package main

import "fmt"

func main(){
	slice := make([]int, 3,4)
	fmt.Println(slice, cap(slice))

	slice = append(slice, 4,6,7,7,7,9,7) //append 后原 cap 不够，进行 cap扩容，扩容不一定严格2倍，也会1/4
	fmt.Println(slice, cap(slice))
}
