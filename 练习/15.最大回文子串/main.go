package main

import (
	"fmt"
	"io"
)

func main() {

	//s := "abaca"
	var n int
	var s string
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		for {
			_, err := fmt.Scan(&s)
			if err == io.EOF {
				return
			} else {
				countSubstrings(s)
			}
		}
	}
}

func countSubstrings(s string) int {
	n := len(s)
	ans := 0
	for i := 0; i < n; i++ {
		for j := 0; j <= 1; j++ { // 0: 中心为1个字符；1：中心为2个字符
			left := i
			right := i + j
			// 中心扩散检测
			for left >= 0 && right < n && s[left] == s[right] {
				ans++
				left--
				right++
			}
		}
	}
	fmt.Println(ans)
	return ans
}
