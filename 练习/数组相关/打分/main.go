/* 跳水比赛，8个评委打分，运动员成绩8个人取消一个最高分和最低分
剩下6个分数的平均分为最后得分，使用一维数组实现：
1.把最高分和最低分的评委找出来
2.找出最佳和最差评委，是打分和最后得分最接近的评委 */

package main
import (
	"fmt"
	"math"
)

func main (){
	var arr = [8]int{75,90,62,82,84,67,83,96}
	max:= arr[0]
	min:= arr[0]
	maxindex := 0
	minindex := 0
	for i:=0; i<len(arr); i++{
		if arr[i]>max{
			max = arr[i]
			maxindex = i
		}
		if arr[i]<min{
			min = arr[i]
			minindex = i
		}
	}
	fmt.Printf("最高分评委:%v号 最高分:%v\n最低分评委:%v号 最低分:%v\n",maxindex,max,minindex,min)
	var ave int
	for i:=0; i<len(arr);i++{
		if arr[i] == max || arr[i] == min {
			
		}else{
			ave += arr[i]
		}
	}
	ave = ave/6
	fmt.Println("ave=",ave)
	var difmaxi float64
	var difmini float64
	var difmax float64
	var difmin float64
	for i :=0; i<len(arr); i++ {
		//求距离平均分数的绝对值
		dif := math.Abs(float64(arr[i]-ave))
		//fmt.Println("dif=",dif)
		//求最大绝对值
		if dif > difmax {           
			difmax = dif
			difmaxi = float64(i)
		}
		//fmt.Println("difmax=",difmax)
		//求最小绝对值，最小绝对值选一个参照数：如最小绝对值肯定小于最大绝对值
		if dif < difmax {   
			difmin = dif
			difmini = float64(i)
		}
	}
	fmt.Printf("最差评委:%v最差评委分距离:%v\n最好评委:%v最好评委分距离:%v",difmaxi,difmax,difmini,difmin)
}