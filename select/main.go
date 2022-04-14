package main

import "fmt"

func main() {
	size := 10
	ch := make(chan int, size)
	for i := 0; i < size; i++ {
		ch <- 1
	}

	ch2 := make(chan int, size)
	for i := 0; i < size; i++ {
		ch2 <- 2
	}

	ch3 := make(chan int, 1)

	select {
	case v := <-ch:
		fmt.Print(v)
	case b := <-ch2:
		fmt.Print(b)
	case ch3 <- 10:
		fmt.Print("write")
	default:
		fmt.Println("none")
	}
}
