package main

import (
	"fmt"
	"strconv"
)

// 思路：通过在map中 东西、南北方向指定加减1，对走过的坐标进行记录
// 如果map中不存在就加入map，如果map中存在就说明交叉了
func main() {
	var path string
	path = "NSE"
	res := isPathCrossing(path)
	fmt.Println(res)
}

func isPathCrossing(path string) bool {
	lenth := len(path)
	x, y := 0, 0
	numMap := make(map[string]int, 10)
	numMap["00"] = 1 //机器人从00原点开始走的
	var ok bool
	tmp := ""
	for i := 0; i < lenth; i++ {
		if path[i] == 'N' {
			y += 1
		} else if path[i] == 'S' {
			y -= 1
		} else if path[i] == 'W' {
			x -= 1
		} else if path[i] == 'E' {
			x += 1
		}
		tmp = strconv.Itoa(x) + strconv.Itoa(y) // strconv.Itoa将整形转化为 字符串类型数字
		_, ok = numMap[tmp]                     //ok 判断是否在map中存在
		if !ok {
			numMap[tmp] = 1
		} else {
			return true
		}
	}
	return false
}
