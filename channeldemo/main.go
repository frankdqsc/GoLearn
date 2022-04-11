package main

import(
	"fmt"
)

func main()  {
	//管道的使用
	//创建一个可以存放3个int类型的管道
	var intChan chan int
	intChan = make(chan int,3)

	//看看intChan 是什么
	fmt.Printf("intChan的值=%v intChan本身的地址=%p\n",intChan,&intChan)

	//向管道中写入数据
	intChan<-10
	num := 211
	intChan <-num

	//写入数据时不能超过其容量
	// intChan <-50
	// intChan <-98

	//管道的长度
	fmt.Printf("channel len=%v cap=%v \n",len(intChan),cap(intChan))

	//从管道中读取数据
	var num2 int
	num2 = <-intChan
	fmt.Println("num2=",num2)
	fmt.Printf("channel len=%v cap=%v \n",len(intChan),cap(intChan))

	//在没有使用协程的情况下，如果数据全部取完，再取就会deadlock
	num3 := <-intChan
	//num4 := <-intChan
	//num5 := <-intChan
	fmt.Println("num3=",num3)
	
}
