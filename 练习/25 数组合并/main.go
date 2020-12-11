package main

import (
	"sort"
	"fmt"
)

func main(){

	nums1:=[]int{1,3,5}
	nums2:=[]int{2,4,6}
	t := len(nums1)
	nums1 = append(nums1[:t], nums2...)
	sort.Ints(nums1)

	fmt.Println(nums1)
}