package main

import (
	"fmt"
	"math/rand"
	"time"
)

func SelectSort(arr *[160000]int) {
	//选择排序较于冒泡排序的优点在于先比较出大值 最后才交换 而冒泡每次比较都交换
	//(*arr)[1] = 600 等价于 arr[1] = 600
	//arr[1] = 900
	//先假设arr[0]为最大值
	for j := 0; j < len(arr)-1; j++ {
		max := arr[j]
		maxIndex := j
		//遍历后面 1- [len(arr) - 1]比较
		for i := j + 1; i < len(arr); i++ {
			if max < arr[i] {
				max = arr[i]
				maxIndex = i
			}
		}
		//交换
		if maxIndex != j { //比较到初始位置 第一次即假设的 maxIndex = 0
			arr[j], arr[maxIndex] = arr[maxIndex], arr[j] //golang特有交换
		}
		//fmt.Printf("第%d次 %v\n", j+1, *arr)
	}
}

func main() {
	t1 := time.Now() //计算算法时间
	//定义一个数组
	//arr := [6]int{10, 24, 19, 100, 80, 700}
	var arr [160000]int
	for i := 0; i < 160000; i++ {
		arr[i] = rand.Intn(200000)
	}
	//fmt.Println(arr)
	SelectSort(&arr)
	//fmt.Println("main函数")
	//fmt.Println(arr)

	elapsed := time.Since(t1)
	fmt.Println("time: ", elapsed)
}

// 14s
