package main

import (
	"fmt"
	"strconv"
	"time"
)

func test03() {
	str := ""
	for i := 0; i < 10; i++ {
		str += "hello" + strconv.Itoa(i)
		fmt.Printf(str)
	}
}

func main() {
	t1 := time.Now()
	test03()
	timecost := time.Since(t1)
	fmt.Println("time: ", timecost)
}
