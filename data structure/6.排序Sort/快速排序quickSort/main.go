package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 原始数组 [9 -1 45 23 2 70 100 10]
// pivot 23
// 一趟后 &[9 -1 10 2 23 70 100 45]
// pivot -1
// 一趟后 &[-1 9 10 2 23 70 100 45]
// pivot 10
// 一趟后 &[-1 9 2 10 23 70 100 45]
// pivot 9
// 一趟后 &[-1 2 9 10 23 70 100 45]

// pivot 70
// 一趟后 &[-1 2 9 10 23 45 100 70]
// pivot 23
// 一趟后 &[-1 2 9 10 23 45 100 70]
// pivot 100
// 一趟后 &[-1 2 9 10 23 45 70 100]
// 结果
// [-1 2 9 10 23 45 70 100]

func QuickSort(left, right int, array *[1600]int) {
	if left < right {
		i, j, pivot := left, right, array[(left+right)/2] //pivot 为中轴
		for i <= j {
			for array[i] < pivot {  //从 pivot 的左边找到大于等于 pivot 的值，如果是 pivot本身也交换
				i++
			}
			for array[j] > pivot {  //从 pivot 的右边找到小于等于pivot 的值
				j--
			}
			if i <= j {
				array[i], array[j] = array[j], array[i]
				i++
				j--
			}
		}
		if left < j {  //向左递归  比较,只要j没走完left右边的就继续
			QuickSort(left, j, array)
		}
		if right > i {
			QuickSort(i, right, array)  //向右递归  只要i没走完right左边的就继续
		}
	}
}
func main() {

	// arr := [8]int{9, -1, 45, 23, 2, 70, 100, 10}
	// fmt.Println("原始数组", arr)
	// QuickSort(0, len(arr)-1, &arr)
	// fmt.Println("结果")
	// fmt.Println(arr)

	t1 := time.Now() //计算算法时间
	var arr [1600]int
	for i := 0; i < 1600; i++ {
		arr[i] = rand.Intn(2000)
	}
	//fmt.Println(arr)
	QuickSort(0, len(arr)-1, &arr)
	//fmt.Println("main函数")
	fmt.Println(arr)

	elapsed := time.Since(t1)
	fmt.Println("time: ", elapsed)

}

// 18ms
