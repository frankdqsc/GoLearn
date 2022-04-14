package main

//随机生成10 个整数 （1-100）保存到数组，
//并倒序打印以及求平均值、求最大值和最大值的下标、
//并查找里面是否有55
import (
	"fmt"
	"math/rand"
	"time"
)

func BubbleSort(intArr *[10]int) {
	temp := 0
	for i := 0; i < len(*intArr)-1; i++ {
		for j := 0; j < 9; j++ {
			if (*intArr)[j] > (*intArr)[j+1] {
				temp = (*intArr)[j]
				(*intArr)[j] = (*intArr)[j+1]
				(*intArr)[j+1] = temp
			}
		}
	}
	fmt.Println("排序后=", (*intArr))
}

func average(intArr *[10]int) {
	var count int
	max := intArr[0] //最大值
	maxindex := 0    //最大值下标（排序后）

	for i := 0; i < len(intArr); i++ {
		count += intArr[i]
		if max < intArr[i] {
			max = intArr[i]
			maxindex = i
		}
	}
	fmt.Println("平均数为：", count/len(intArr))
	fmt.Printf("最大值为max = %v 最大值下标为maxindex = %v\n", max, maxindex)
}

func check(intArr *[10]int) {
	for i := 0; i < len(intArr); i++ {
		if intArr[i] != 55 {
			fmt.Printf("没有55")
			break
		} else {
			fmt.Printf("有55")
		}
	}
}
func main() {
	var intArr [10]int
	//求随机数
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(intArr); i++ {
		intArr[i] = rand.Intn(100)
	}
	fmt.Println(intArr)
	BubbleSort(&intArr) //冒泡排序
	average(&intArr)    //求平均值和最大值、最大值下标
	check(&intArr)      //检查是否有55
}
