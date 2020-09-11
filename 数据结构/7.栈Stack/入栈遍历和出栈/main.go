package main

import (
	"errors"
	"fmt"
)

//使用数组模拟栈

//栈顶 栈底 数组本身 ->结构体
type Stack struct {
	MaxTop int    //表示栈最大存储数
	Top    int    //栈顶，栈底不变 直接使用top
	arr    [5]int //数组模拟栈
}

//入栈
func (this *Stack) Push(val int) (err error) {

	//判满
	if this.Top == this.MaxTop-1 {
		fmt.Println("stack full")
		return errors.New("stack full")
	}

	this.Top++
	//放入数据
	this.arr[this.Top] = val
	return
}

//遍历栈 需要从栈顶开始
func (this *Stack) List() {
	//先判空
	if this.Top == -1 {
		fmt.Println("stack empty")
		return
	}

	fmt.Println("栈的情况如下：")
	for i := this.Top; i >= 0; i-- {
		fmt.Printf("arr[%d]=%d\n", i, this.arr[i])
	}
}

//出栈
func (this *Stack) Pop() (val int, err error) {

	//判空
	if this.Top == -1 {
		fmt.Println("stack empty")
		return 0, errors.New("stack empty")
	}

	//先取值 再this.Top--
	val = this.arr[this.Top]
	this.Top--
	return val, nil
}

func main() {

	stack := &Stack{
		MaxTop: 5,  //最多存放5个数到栈中
		Top:    -1, //当栈顶为-1，表示栈为空
	}

	//入栈
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	//显示
	stack.List()
	//出栈
	val, _ := stack.Pop()
	fmt.Println("出栈val=", val)
	//显示
	stack.List()

	val, _ = stack.Pop()
	fmt.Println("出栈val=", val)
	//显示
	stack.List()
}
