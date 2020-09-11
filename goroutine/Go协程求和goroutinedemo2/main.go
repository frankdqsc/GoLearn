//要求
//1.启动一个协程，将1-2000的数放入到一个channel 中，比如 numChan
//2.启动8个协程，从numChan取出数（比如n）计算1+2+...+n的值，并存放到resChan
//3.最后8个协程同时完成工作后，在遍历resChan 显示结果（如res[1] = 1_res[10] = 55..）
//注意，考虑resChan chan int是否合适

// 4内核

package main

import (
	"fmt"
	"time"
)

func numIn(numChan chan int){  //协程
	for i:=1;i<=20000;i++{
		numChan <- i
	}
	//关闭numChan
	close(numChan)
}

func numCal(numChan chan int,resChan chan int,exitChan chan bool){

	for{
		num,ok := <-numChan
		//fmt.Println("num=",num)
		if !ok{ //intChan取不到
			break
		}
		sum:= 0 
		for i:=1;i<=num;i++{
			sum += i
		}
		//fmt.Println("sum=",sum)
		resChan <- sum
		sum = 0
	}
	fmt.Println("有一个resChan协程因为取不到数据而退出")
	//这里不关闭resChan （自己取不到，别人可能还在操作）向exitChan写入true
	exitChan <- true
}

func main(){
	t1 := time.Now() //计算算法时间
	
	numChan := make(chan int,2000)
	resChan := make(chan int,2000)
	exitChan := make(chan bool,4)
	go numIn(numChan)

	//2.启动8个协程，从numChan取出数（比如n）计算1+2+...+n的值，并存放到resChan
	// resChan := 0
	// for v := range numChan{
	// 	resChan += v
	// 	fmt.Println("resChan=",resChan)
	
	for i:=0;i<4;i++{
		go numCal(numChan,resChan,exitChan)
	}                                                                                                                                                                                                                              

	//这里进行主线程处理，防止主线程过快结束
	go func(){ //匿名函数
		for i:=0;i<4;i++{
			<-exitChan
		}
		elapsed := time.Since(t1)
   		fmt.Println("time: ", elapsed)
		//当我们从exitChan中取出4个结果 就可以关闭resChan
		close(resChan)
	}()
	
	//遍历resChan把结果取出来
	for {
		_,ok := <- resChan
		if !ok{
			break
		}
		//结果取出
		//fmt.Printf("res=%d\n",res)
	}

	fmt.Println("main线程退出")
}