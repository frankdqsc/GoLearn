package main

//稀疏矩阵存取
import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type ValNode struct {
	row int
	col int
	val int
}

func main() {
	//1.先创建一个数组
	var chessMap [11][11]int
	chessMap[1][2] = 7 //黑子
	chessMap[3][5] = 8 //白子

	//2.输出查看
	for _, v := range chessMap { // v 为一维数组
		for _, v2 := range v { // 对 v 再遍历 取出具体的值 v2
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}

	//3.转成稀疏矩阵
	//思路 -> 遍历 chessMap 如果某一元素不为0 创建node结构体
	//将其放入到对应的切片即可v

	var sparseArr []ValNode

	//标准的稀疏数组应该还有一个记录元素二维数组得规模 (行和列 以及默认值)
	valNode := ValNode{
		row: 11,
		col: 11,
		val: 0,
	}
	sparseArr = append(sparseArr, valNode)

	for i, v := range chessMap {
		for j, v2 := range v {
			if v2 != 0 {
				//创建一个ValNode值结点
				valNode := ValNode{
					row: i,
					col: j,
					val: v2,
				}
				sparseArr = append(sparseArr, valNode)
			}
		}
	}

	//输出稀疏数组
	// 0: 11 11 0
	// 1: 1 2 7
	// 2: 3 5 8

	fmt.Println("当前的稀疏数组是：：：：：")
	str, str1, str2 := "", "", ""
	for i, valNode := range sparseArr {
		if i > 0 && i < 2 {
			str1 = strconv.Itoa(valNode.row) + " " + strconv.Itoa(valNode.col) + " " + strconv.Itoa(valNode.val)
		}
		if i == 2 {
			str2 = strconv.Itoa(valNode.row) + " " + strconv.Itoa(valNode.col) + " " + strconv.Itoa(valNode.val)
		}
		fmt.Printf("%d: %d %d %d\n", i, valNode.row, valNode.col, valNode.val)
	}
	str = str1 + " " + "\n" + str2
	fmt.Println(str)
	fmt.Println()

	//将这个稀疏数组 存盘到 e:/chessmap.data

	filePath := "e:/chessmap.data"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("Open file err = %v\n", err)
		return
	}
	//及时关闭file句柄
	defer file.Close()
	//写入时使用带缓存的write
	write := bufio.NewWriter(file)
	write.WriteString(str)
	write.Flush()

	// 如何恢复原始数组？
	// 1.打开 e:/chessmap.data  =>恢复原始数组
	// 1 2 7
	// 3 5 8
	//思路
	//首相将文本中的数值读到另一个二维数组
	//然后比较两个数组的值进行赋值
	var chessMap1 [2][3]int

	f, err := ioutil.ReadFile("e:/map.data")
	// f:
	// 1 2 7
	// 3 5 8
	if err != nil {
		fmt.Println("read fail", err)
	}
	fmt.Println(string(f))
	for m := 0; m < 2; m++ {
		for n := 0; n < 3; n++ {
			fmt.Scanln(string(f), &chessMap1[m][n])
		}
	}
	for m := 0; m < 2; m++ {
		for n := 0; n < 3; n++ {
			fmt.Println(chessMap1[m][n])
		}
	}

	// var chessMap2 [11][11]int
	// for i := 0; i < 11; i++ {
	// 	for j := 0; j < 11; j++ {
	// 		if i == chessMap1[0][0] && j == chessMap1[0][1] {
	// 			chessMap2[i][j] = chessMap1[0][2]
	// 		}
	// 		if i == chessMap1[1][1] && j == chessMap1[1][2] {
	// 			chessMap2[i][j] = chessMap1[1][2]
	// 		}
	// 		fmt.Printf("%v ", chessMap2[i][j])
	// 	}
	// 	//fmt.Println()
	// }

	//fmt.Println("chessMap1", chessMap)
	// fmt.Println("恢复后的原始数据：：：")
	// for _, v := range chessMap2 {
	// 	for _, v2 := range v {
	// 		fmt.Printf("%d\t", v2)
	// 	}
	// 	fmt.Println()
	// }
}
