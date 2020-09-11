package main

//传统方法计算8000内所有素数 时间  time:  120.6596ms+
import (
	"fmt"
	"time"
)

func main() {
	t1 := time.Now() //计算算法时间

	for num := 1; num <= 8000; num++ {
		flag := true
		for i := 2; i < num; i++ {
			if num%i == 0 { //不是素数
				flag = false
				break
			}
		}
		if flag {
			fmt.Println("num: ", num) //计算算法时间不算打印部分
		}
	}
	elapsed := time.Since(t1)
	fmt.Println("time: ", elapsed)
}
