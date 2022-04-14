package main

import (
	"errors"
	"fmt"
	"strconv"
)

//使用数组模拟栈

//栈顶 栈底 数组本身 ->结构体
type Stack struct {
	MaxTop int     //表示栈最大存储数
	Top    int     //栈顶，栈底不变 直接使用top
	arr    [20]int //数组模拟栈
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

//判断一个字符是不是运算符 加减乘除
func (this *Stack) IsOper(val int) bool {
	if val == 42 || val == 43 || val == 45 || val == 47 {
		return true
	} else {
		return false
	}
}

//运算的方法
func (this *Stack) Cal(num1, num2, oper int) int {
	res := 0
	switch oper {
	case 42:
		res = num2 * num1
	case 43:
		res = num2 + num1
	case 45:
		res = num2 - num1
	case 47:
		res = num2 / num1
	default:
		fmt.Println("运算符错误")
	}
	return res
}

//方法：运算符优先级定义
func (this *Stack) Priority(oper int) int {

	res := 0
	if oper == 42 || oper == 47 {
		res = 1
	} else if oper == 43 || oper == 45 { //加43 减45 乘42 除47
		res = 0
	}
	return res
}

func main() {

	//数栈
	numStack := &Stack{
		MaxTop: 20,
		Top:    -1,
	}
	//符号栈
	operStack := &Stack{
		MaxTop: 20,
		Top:    -1,
	}

	exp := "3+30*6-400"
	//定义一个index 帮助扫描exp

	index := 0
	//为了配合运算 定义需要的变量
	num1, num2, oper, result := 0, 0, 0, 0
	keepNum := ""

	for {
		ch := exp[index : index+1] //字符串
		// ch==> "+"  -->  43
		temp := int([]byte(ch)[0])  // 放入切片 取切片第一个 转int 得到字符对应的ASCII码
		if operStack.IsOper(temp) { //说明是符号
			//如果是符号 分两个逻辑,空栈和不是空栈
			if operStack.Top == -1 { //如果 operStack 是一个空栈 直接入栈
				operStack.Push(temp)
			} else {
				//如果发现operStack栈顶的运算符的优先级大于等于当前准备入栈的运算符的优先级
				//就从符号栈 pop 出，并从数栈也pop 两个数，进行运算，运算的结果在重新入栈到
				//数栈，符号再入符号栈
				if operStack.Priority(operStack.arr[operStack.Top]) >= operStack.Priority(temp) {
					num1, _ = numStack.Pop()
					num2, _ = numStack.Pop()
					oper, _ = operStack.Pop()
					result = operStack.Cal(num1, num2, oper)
					//将计算结果重新入数栈
					numStack.Push(result)
					//当前的符号压入符号栈
					operStack.Push(temp)
				} else {
					operStack.Push(temp)
				}
			}
		} else { //说明是数

			//处理多位数
			//1.定义一个变量 keepNum string 做拼接
			keepNum += ch
			//2.每次向 index 的前面字符测试一下 看看是不是运算符

			//如果已经到表达式最后 直接将 keepNum
			if index == len(exp)-1 {
				val, _ := strconv.ParseInt(keepNum, 10, 64)
				numStack.Push(int(val))
			} else {
				//没到最后一位，向 index 后测试 是不是运算符
				if operStack.IsOper(int([]byte(exp[index+1 : index+2])[0])) {
					val, _ := strconv.ParseInt(keepNum, 10, 64)
					numStack.Push(int(val))
					keepNum = ""
				}
			}
		}

		//继续扫描
		//先判断 index 是否已经扫描到计算表达式的最后
		if index+1 == len(exp) {
			break
		} else {
			index++
		}
	}

	//如果扫面表达式完毕 依次从符号栈取出符号 然后从数栈取出两个数
	//运算后的结果，入数栈，直到符号栈为空
	for {
		if operStack.Top == -1 {
			break
		}
		num1, _ = numStack.Pop()
		num2, _ = numStack.Pop()
		oper, _ = operStack.Pop()
		result = operStack.Cal(num1, num2, oper)
		//将计算结果重新入数栈
		numStack.Push(result)
	}
	//如果算法没问题 表达式也正确 则结果就是 numStack 最后数
	res, _ := numStack.Pop()
	fmt.Printf("表达式%s = %v", exp, res)
}
