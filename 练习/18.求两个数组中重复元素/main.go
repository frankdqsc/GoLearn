package main

import (
	"fmt"
	"io"
)

func intersect(nums1 []int, nums2 []int) []int {
	m0 := map[int]int{}
	for _, v := range nums1 {
		m0[v] += 1
	}
	k := 0
	//比如nums2中k0 v=4 则在m0中判断 m0[4]是否大于1 大于就是nums1中存在
	//然后将 重复的此值 v 加到 nums2[0] 位置
	for _, v := range nums2 {
		if m0[v] > 0 {
			m0[v] -= 1
			nums2[k] = v
			k++
		}
	}
	return nums2[0:k] //将重复的值分割出
}

func main() {
	var n, p, q int
	var nums1 []int
	var nums2 []int
	var t1, t2, norepe, repe int
	var res []int
	fmt.Scan(&n, &p, &q)

	for i := 0; i < 3; i++ {
		_, err := fmt.Scan(&t1)
		if t1 > n {
			return
		} else {
			nums1 = append(nums1, t1)
			if err == io.EOF {
				return
			}
		}
	}
	for i := 0; i < 3; i++ {
		_, err := fmt.Scan(&t2)
		if t2 > n {
			return
		} else {
			nums2 = append(nums2, t2)
			if err == io.EOF {
				return
			}
		}
	}
	//fmt.Println("nums1:", nums1)
	//fmt.Println("nums2:", nums2)

	nums2 = intersect(nums1, nums2)
	repe = len(nums2)
	norepe = 3 - repe
	res = append(res, norepe, norepe, repe)
	fmt.Println(res)

}
