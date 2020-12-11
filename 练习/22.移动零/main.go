package main

import (
	"fmt"
	"time"
)
func movezero(slice []int)[]int{
	t := 0
	for i := 0; i < len(slice); i++ {
		if (slice[i] != 0){
			if(i != t){
				slice[t] = slice[i]
				slice[i] = 0
			}
			t++
		}
	}
	return slice
}
func main() {

	/*slice := make([]int,0)
	var n,m int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		//fmt.Scan(&m)
		//slice = append(slice,rand.Intn(2000))
		fmt.Scan(&m)
		slice = append(slice,m)
	}*/
	slice := []int{1,3,4,0,3,4,0,5,6}
	t1 := time.Now() //计算算法时间
	slice = movezero(slice)

	elapsed := time.Since(t1)
	fmt.Println("time: ", elapsed)

	fmt.Println(slice)
}

