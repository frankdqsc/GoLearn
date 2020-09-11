package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var count int
	rand.Seed(time.Now().UnixNano())
	for {
		n := rand.Intn(100) //范围
		fmt.Println(n)
		count++
		if n == 99 {
			break
		}
	}
	fmt.Printf("随机到99共用了%d次", count)
}
