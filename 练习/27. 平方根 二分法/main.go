package main

import "fmt"

func main() {

	x:= 10

	l, r := 0, x
	ans := -1
	for l <= r {
		mid := (l + r) / 2
		//fmt.Println(mid)
		if mid*mid > x {
			r = mid - 1
		} else {
			ans = mid
			l = mid + 1
		}
	}
	fmt.Println(ans)
}
