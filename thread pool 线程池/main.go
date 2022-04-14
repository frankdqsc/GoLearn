package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)
var wg sync.WaitGroup
var IDS_ALL = []int {40,42,43,44,164,166,171,173,174,175,177,518,192,193}

func main () {
	fmt.Println("BEGIN")

	poolSize := runtime.NumCPU()
	fmt.Println(poolSize)
	runtime.GOMAXPROCS(poolSize)

	ch := make(chan int,poolSize)
	for _,catid := range IDS_ALL {
		ch <- 1
		wg.Add(1)
		go work(catid,ch)
	}
	wg.Wait()
	close(ch)

	fmt.Println("DONE")
}

// 工作
func work(catid int,ch chan int) {

	fmt.Println("WORKING ",catid)
	time.Sleep(2 * 1e9)

	wg.Done()
	<-ch
}