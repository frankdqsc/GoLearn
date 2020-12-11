package main

//给定一个包含 0, 1, 2, ..., n 中 n 个数的序列，找出 0 .. n 中没有出现在序列中的那个数
//序列不一定从0开始
import (
	"fmt"
	"sort"
)
func missingNumber(nums []int) int {
	sort.Ints(nums)
	l:=len(nums)

	l1:= nums[0]
	l2:= nums[l-1]

	sum:= (l+1)*(l1+l2)/2
	for _,v:=range nums{
		sum-=v
	}
	return sum
}

func main(){
	var res int
	var slice = []int{1,3,4,6,7,8}
	res = missingNumber(slice)
	fmt.Println(res)
}
