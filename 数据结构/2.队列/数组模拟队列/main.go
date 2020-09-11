package main

import (
	"errors"
	"fmt"
	"os"
)

//使用一个结构体管理队列
type Queue struct {
	maxSize int
	array   [4]int //数组 =>模拟队列
	front   int    //指向队首
	rear    int    //指向队列尾部
}

//添加数据到队列
func (this *Queue) AddQueue(val int) (err error) {

	//先判断队列是否已满
	if this.rear == this.maxSize-1 { //rear 是队列尾部 包含最后元素
		return errors.New("queue full")
	}

	this.rear++ //rear 后移
	this.array[this.rear] = val
	return
}

//显示队列,找到队首 遍历到队尾
func (this *Queue) ShowQuene() {
	//this.front 不包含队首的元素
	for i := this.front + 1; i <= this.rear; i++ {
		fmt.Printf("array[%d] = %d\t", i, this.array[i])
	}
	fmt.Println()
}

//从队列中取出数据
func (this *Queue) GetQueue() (val int, err error) {
	//先判断队列是否为空
	if this.rear == this.front {
		return -1, errors.New("queue empty")
	}
	this.front++
	val = this.array[this.front]
	return val, err
}

//编写主函数测试
func main() {

	//创建一个队列
	queue := &Queue{
		maxSize: 5,
		front:   -1,
		rear:    -1,
	}
	var key string
	var val int
	for {
		fmt.Println("1.输入add 表示添加数据到队列")
		fmt.Println("2.输入get 表示从队列获取数据")
		fmt.Println("3.输入show 表示显示队列")
		fmt.Println("4.输入exit 退出")

		fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Println("请输入你要入队列数")
			fmt.Scanln(&val)
			err := queue.AddQueue(val)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("加入队列成功")
			}
		case "get":
			val, err := queue.GetQueue()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("从队列取出了一个数", val)
			}

		case "show":
			queue.ShowQuene()
		case "exit":
			os.Exit(0)
		}
	}
}
