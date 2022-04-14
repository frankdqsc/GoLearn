package main

import (
	"fmt"
	"math/rand"
	"time"
)

//  原始数组= [23 0 12 56 34 -1 50]
// 第1次插入后[23 0 12 56 34 -1 50]
// 第2次插入  [23 0  0 56 34 -1 50]  arr[insertIndex+1] = arr[insertIndex] 数据后移

// 第2次插入后[23 12 0 56 34 -1 50]
// 第3次插入后[56 23 12 0 34 -1 50]
// 第4次插入后[56 34 23 12 0 -1 50]
// 第5次插入后[56 34 23 12 0 -1 50]
// 第6次插入后[56 50 34 23 12 0 -1]
func InsertSort(arr *[160000]int) {

	for i := 1; i < len(arr); i++ {
		//完成第一次 给第二次元素找到合适的位置 并插入
		insertVal := arr[i]  //要插入的值先保存在 insertVal
		insertIndex := i - 1 //插入下标，总是指向要插元素前一个

		//从大到小
		//【i最小为0,如果没有insertIndex >= 0 ,i会越界  导致 如果没有insertIndex，index out of range [-1]
		for insertIndex >= 0 && insertVal > arr[insertIndex] { //insertVal > arr[insertIndex 向前比较
			arr[insertIndex+1] = arr[insertIndex] //数据后移
			insertIndex--
			//fmt.Println("insertIndex", insertIndex)
		}
		//插入
		arr[insertIndex+1] = insertVal
		//fmt.Printf("第%d次插入后%v\n", i, *arr)
	}

	// //完成第2次 给第3次元素找到合适的位置 并插入
	// insertVal = arr[2]
	// insertIndex = 2 - 1 //下标

	// //从大到小
	// for insertIndex >= 1 && insertVal > arr[insertIndex] {
	// 	arr[insertIndex+1] = arr[insertIndex] //数据后移
	// 	insertIndex--
	// }
	// //插入
	// if insertIndex+1 != 2 {
	// 	arr[insertIndex+1] = insertVal
	// }
	// fmt.Println("第2次插入后", *arr)
}

func main() {

	// arr := [7]int{23, 0, 12, 56, 34, -1, 50}
	// fmt.Println("原始数组=", arr)
	// InsertSort(&arr)

	// fmt.Println("main函数")
	// fmt.Println(arr)

	t1 := time.Now() //计算算法时间
	var arr [160000]int
	for i := 0; i < 160000; i++ {
		arr[i] = rand.Intn(200000)
	}
	//fmt.Println(arr)
	InsertSort(&arr)
	//fmt.Println("main函数")
	//fmt.Println(arr)

	elapsed := time.Since(t1)
	fmt.Println("time: ", elapsed)
}

// 7.7s
