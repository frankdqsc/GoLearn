package main

import(
	"fmt"
	"time"
	"sync"
)

//要求：现在要计算1-200各个数的阶乘，并把各个数的阶乘放到map中
//最后显示出来，要求使用goroutine完成
//思路
//1.编写一个函数，来计算各个数的阶乘，并放入map中
//2.启动多个协程，统计的结果放入map中
//3.map应该做出一个全局

//问题：由于同时向map中写，出现资源争夺问题
//解决方法：
//1.加入全局锁
//2.channel

var (
	myMap = make(map[int]int,10)
	//声明一个全局互斥锁  lock是全局互斥锁  sync是包 synchornized同步  Mutex互斥
	lock sync.Mutex
)
func test(n int){
	res := 1
	for i:=1; i<=n; i++{
		res *= i
	}
	//将res放入到myMap
	//加锁
	lock.Lock()
	myMap[n] = res
	//解锁
	lock.Unlock()
}

func main(){
	//在这里开启多个协程完成这个任务 200个
	for i:=1;i<=20;i++{
		go test(i)
	}

	//休眠10秒
	time.Sleep(time.Second*3)  //这里不知道子程序多久才能运行完 →channel

	//输出结果
	lock.Lock()
	for i,v:=range myMap{
		fmt.Printf("map[%d]=%d\n",i,v)
	}
	lock.Unlock()
}