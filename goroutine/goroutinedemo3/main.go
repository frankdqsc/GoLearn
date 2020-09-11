// 要求统计1-8000的数字中哪些是素数？
// 思路：开三个协程 1放8000个数  2取出判断  3判断退出
//一个协程putNum 1-8000  四个协程primeNum  三个管道intChan[1000] primeChan[2000] exitChanp[4]

package main

//go开协程（goroutine）计算8000内所有素数 时间  time:  5~6 ms
import (
	"fmt"
	"time"
)

//向intChan中放入1-8000个数
func putNum(intChan chan int) { //协程
	for i := 1; i <= 8000; i++ {
		intChan <- i
	}
	//关闭intChan
	close(intChan)
}

//从intChan取出数据 并判断是否为素数 如果是放入到primeChan
func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
	//使用for循环
	//var num int
	//time.Sleep(time.Millisecond * 10)
	var flag bool
	for {
		num, ok := <-intChan
		if !ok { //intChan取不到
			break
		}
		flag = true
		for i := 2; i < num; i++ {
			if num%i == 0 { //不是素数
				flag = false
				break
			}
		}
		if flag {
			//将这个数放到primeChan
			primeChan <- num
		}
		//fmt.Println("路过一次")
	}
	fmt.Println("有一个primeChan协程因为取不到数据而退出")
	//这里不关闭primeChan （自己取不到，别人可能还在操作）向exitChan写入true
	exitChan <- true
}

func main() {
	t1 := time.Now() //计算算法时间

	intChan := make(chan int, 1000)
	primeChan := make(chan int, 2000) //放入结果
	//标识退出的管道
	exitChan := make(chan bool, 4) //4个

	//开启一个协程，向intChan放入1-8000个数
	go putNum(intChan)

	//开启4个协程 从intChan取出数据，并判断是u否为素数，如果是就放入primeChan
	//放入到primeChan
	for i := 0; i < 8; i++ {
		go primeNum(intChan, primeChan, exitChan)
	}

	//这里进行主线程处理，防止主线程过快结束
	go func() { //匿名函数
		for i := 0; i < 4; i++ {
			<-exitChan
		}
		//当我们从exitChan中取出4个结果 就可以关闭primeChan

		elapsed := time.Since(t1)
		fmt.Println("time: ", elapsed)

		close(primeChan)
	}()

	//遍历primeChan把结果取出来
	for {
		_, ok := <-primeChan
		if !ok {
			break
		}
		//结果取出
		//fmt.Printf("素数=%d\n", res)
	}

	fmt.Println("main线程退出")
}
