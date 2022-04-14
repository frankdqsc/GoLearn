/* 编写一个函数，可以接收一个数组，该数组有5个数，
请找出最大的数和最小的数和对应的数组下标 */
package main
import "fmt"

func fb(arr [5]int){
	var maxindex,minindex int
	max:= arr[0]
	min:= arr[0]
	for i:=0; i <len(arr); i++{
		if arr[i] > max{
		max = arr[i]
		maxindex = i
		}
		if arr[i] < min{
		min = arr[i]
		minindex = i
		}
	}
	fmt.Printf("max值为：%v 下标为: %v\n",max,maxindex)
	fmt.Printf("min值为：%v 下标为: %v",min,minindex)
}

func main(){
	var arr = [5]int{9,32,12,54,67}
	fb(arr)
}