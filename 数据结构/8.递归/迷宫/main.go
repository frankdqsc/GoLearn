package main

import "fmt"

//myMap *[8][7]int 保证是同一个地图
//i j 表示对地图的哪个点进行测试
func SetWay(myMap *[8][7]int, i, j int) bool {

	//分析出什么情况下 就找到出路
	//myMap[6][5] == 2
	if myMap[6][5] == 2 {
		return true
	} else {
		//说明继续找
		if myMap[i][j] == 0 { //可以探测且没探测

			//假设这个点可以通过 但是需要探测上下左右
			//换成 下右上左
			myMap[i][j] = 2
			if SetWay(myMap, i+1, j) { //下
				return true
			} else if SetWay(myMap, i, j+1) { //右
				return true
			} else if SetWay(myMap, i-1, j) { //上
				return true
			} else if SetWay(myMap, i, j-1) { //左
				return true
			} else { //死路
				myMap[i][j] = 3
				return false
			}
		} else { //不通 为1
			return false
		}
	}
}

func main() {

	//先创建一个二维数组 模拟迷宫
	//规则
	//1.如果元素的值为1,就是墙
	//2.如果元素的值为0，是没有走过的点；为2，是一个通路，为3，是走过的不通路
	var myMap [8][7]int

	myMap[3][1] = 1
	myMap[3][2] = 1
	// myMap[1][2] = 1
	// myMap[2][2] = 1 //封住

	//先把地图的最上和最下设置为1
	for i := 0; i < 7; i++ {
		myMap[0][i] = 1
		myMap[7][i] = 1
	}
	//最左 最右
	for i := 0; i < 7; i++ {
		myMap[i][0] = 1
		myMap[i][6] = 1
	}
	//输出地图
	for i := 0; i < 8; i++ {
		for j := 0; j < 7; j++ {
			fmt.Print(myMap[i][j], " ")
		}
		fmt.Println()
	}

	//测试
	SetWay(&myMap, 1, 1)
	fmt.Println("探测完成的地图")
	//输出地图
	for i := 0; i < 8; i++ {
		for j := 0; j < 7; j++ {
			fmt.Print(myMap[i][j], " ")
		}
		fmt.Println()
	}
}
