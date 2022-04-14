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

	////将这个稀疏数组 存盘到 e:/chessmap.data
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
	var chessMap1 [2][3]int
	var chessMap2 [11][11]int
	f, err := ioutil.ReadFile("e:/chessmap.data")
	if err != nil {
		fmt.Println("read fail", err)
	}

	for m := 0; m < 2; m++ {
		for n := 0; n < 3; n++ {
			for _, v := range f {
				chessMap1[m][n], err = strconv.Atoi(string(v))
				if err != nil {
					fmt.Println("strconv.Atoi(string(f)) err =", err)
				}
			}
			for i := 0; i < 11; i++ {
				for j := 0; j < 11; j++ {
					if i == chessMap1[0][0] && j == chessMap1[0][1] {
						chessMap2[i][j] = chessMap1[0][2]
					}
					if i == chessMap1[1][1] && j == chessMap1[1][2] {
						chessMap2[i][j] = chessMap1[1][2]
					}
					fmt.Printf("%v ", chessMap2[i][j])
				}
				//fmt.Println()
			}
		}
	}

	//fmt.Println("string(f)", string(f))
	//fmt.Println()

	// file, err = os.Open("e:/chessmap.data")
	// if err != nil {
	// 	fmt.Println("open file err=", err)
	// }
	//输出文件
	//fmt.Printf("file=%v", file) //  file=&{0xc00014aa00} 文件是指针

	//创建一个 reader 带缓冲 缓冲区4096
	//reader := bufio.NewReader(file)
	//循环读取文件中的内容
	// 1 2 7
	// 3 5 8
	//var chessMap2 [11][11]int
	// for {
	// 	str3, err = reader.ReadString('\n') //读到换行结束
	// 	if err == io.EOF {                  //io.EOF代表文件末尾
	// 		break
	// 	}
	// 	fmt.Println("str3:", str3)
	// }
	// for m := 0; m < 11; m++ {
	// 	for n := 0; n < 11; n++ {

	// 	}
	// }

	// var chessMap2 [11][11]int
	// for i, _ := range str {
	// 	chessMap2[valNode.row][valNode.col] = str
	// }

	//及时关闭文件 否则会内存泄漏
	// file.Close()
	// if err != nil {
	// 	fmt.Println("close file err=", err)
	// }
	//2.这里使用稀疏数组恢复

	//先创建一个原始数组
	// var chessMap2 [11][11]int

	// //遍历 sparseArr  （遍历文件每一行）
	// 0: 11 11 0
	// 1: 1 2 7
	// 2: 3 5 8
	// for i, valNode := range sparseArr {
	// 	if i != 0 { //跳过第一行
	// 		chessMap2[valNode.row][valNode.col] = valNode.val
	// 	}
	// }
	//看看chessMap2 是不是恢复

}
