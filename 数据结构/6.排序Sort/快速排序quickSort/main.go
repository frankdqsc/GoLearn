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

func QuickSort(left, right int, array *[160000]int) {

	if left < right {
		i, j := left, right
		//pivot 为中轴
		pivot := array[(left+right)/2]
		//fmt.Println("pivot", pivot)

		for i <= j {
			//从 pivot 的左边找到大于等于 pivot 的值，如果是 pivot本身也交换
			for array[i] < pivot {
				i++
			}
			//从 pivot 的右边找到小于等于pivot 的值
			for array[j] > pivot {
				j--
			}
			if i <= j {
				array[i], array[j] = array[j], array[i]
				i++
				j--
			}
		}

		//fmt.Println("一趟后", array)
		//向左递归  比较,只要j没走完left右边的就继续
		if left < j {
			QuickSort(left, j, array)
		}
		//向右递归  只要i走完right左边的就继续
		if right > i {
			QuickSort(i, right, array)
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
	var arr [160000]int
	for i := 0; i < 160000; i++ {
		arr[i] = rand.Intn(200000)
	}
	//fmt.Println(arr)
	QuickSort(0, len(arr)-1, &arr)
	//fmt.Println("main函数")
	//fmt.Println(arr)

	elapsed := time.Since(t1)
	fmt.Println("time: ", elapsed)

}

// 18ms
