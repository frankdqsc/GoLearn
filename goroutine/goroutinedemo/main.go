package main

import(
	"fmt"
	"time"
	"strconv"
)

//1.在主线程（可以理解为进程）中，开启一个goroutine,该协程每隔1秒输出“Hello,world”
//2.在主线程中也每隔1秒输出“Hello，Golang”,输出10次后，退出程序
//3.要求主线程和goroutine同时执行

//编写一个函数，每隔1秒输出“Hello,world”
func test(){
	for i:=1; i<10; i++{
		fmt.Println("test ()hello,world"+ strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func main(){

	go test() //加上go表示开启了协程

	for i:=1; i<10; i++{
		fmt.Println("main ()hello,golang"+ strconv.Itoa(i))
		time.Sleep(time.Second)
	}
	
}
