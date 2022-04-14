package main

import (
	"runtime"
	"time"
)
/*其中创建一个goroutine并挂起， main goroutine 优先调用了 休眠，
此时唯一的 P 会转去执行 for 循环所创建的 goroutine，
进而 main goroutine 永远不会再被调度。
换一句话说在Go1.14之前，上边的代码永远不会输出OK，
因为这种协作式的抢占式调度是不会使一个没有主动放弃执行权、
且不参与任何函数调用的goroutine被抢占。

Go1.14 实现了基于信号的真抢占式调度解决了上述问题。Go1.14程序启动时，
在 runtime.sighandler 函数中注册了 SIGURG 信号的处理函数 runtime.doSigPreempt，
在触发垃圾回收的栈扫描时，调用函数挂起goroutine，并向M发送信号，M收到信号后，
会让当前goroutine陷入休眠继续执行其他的goroutine。*/
func main() {
	runtime.GOMAXPROCS(1)
	var count int
	go func() {
		for {
			println("O111111K")
			count++
		}
	}()

	time.Sleep(time.Millisecond)
	println("OK")
	println(count)
}

//go 1.13报错