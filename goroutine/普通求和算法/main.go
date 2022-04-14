package main

import (
	"fmt"
	"time"
)

func main() {
	t1 := time.Now() //计算算法时间
	sum := 0
	for num := 1; num <= 20000; num++ {
		for i := 1; i <= num; i++ {
			sum += i
		}
		//fmt.Println("sum=",sum)
		sum = 0
	}
	elapsed := time.Since(t1)
	fmt.Println("time: ", elapsed)
}
