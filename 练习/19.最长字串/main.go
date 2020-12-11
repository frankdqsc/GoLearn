package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "abcabe"
	res := lengthOfLongestSubstring(s)
	fmt.Println(res)
}

func lengthOfLongestSubstring(s string) int {
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		index := strings.Index(s[start:i], string(s[i])) //判断 string(s[i]) 在 s[start:i] 中首次出现的位置，没有返回-1

		if index == -1 {
			if i+1 > end {
				end = i + 1
			}
		} else {
			start += index + 1
			end += index + 1
		}
	}

	fmt.Println(end - start)
	return end - start
}
