package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(1) // 1个线程
	wg := sync.WaitGroup{}
	wg.Add(20)

	for j := 0; j < 10; j++ { //j是外部 for的一个变量，i=1 i=2...i=10 地址不变化。遍历完成后，最终 i=10。故 go func执行时，i 的值始终是10。
		go func() {
			fmt.Println("A: ", j)
			wg.Done()
		}()
	}
	for i := 0; i < 10; i++ { //i是函数参数，与外部 for中的 i完全是两个变量。尾部(i)将发生值拷贝，go func内部指向值拷贝地址。
		go func(i int) {
			fmt.Println("B: ", i)
			wg.Done()
		}(i) // i的地址变化
	}
	wg.Wait()
}

