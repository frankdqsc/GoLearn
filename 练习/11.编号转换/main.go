package main

import (
	"fmt"
	"strconv"
)

func change(str string) (res string) {
	r, c := 0, 0
	_, err := fmt.Sscanf(str, "R%dC%d", &r, &c)
	if err == nil {
		res = strconv.Itoa(r)
		for c > 26 {
			if c%26 == 0 {
				res = string('Z') + res
				c = c/26 - 1
			} else {
				res = string('A'+c%26-1) + res
				c = c / 26
			}
		}
		if c%26 == 0 {
			res = string('Z') + res
		} else {
			res = string('A'+c%26-1) + res
		}
		return
	}
	r, c = 0, 0
	for t, v := range str {
		if v >= 'A' && v <= 'Z' {
			c = 26*c + int(v-'A') + 1
		} else {
			r = t
			break
		}
	}
	res = "R" + string(str[r:]) + "C" + strconv.Itoa(c)
	return
}

func main() {
	n := 0
	fmt.Scanf("%d\n", &n)
	str := ""
	for i := 0; i < n; i++ {
		fmt.Scanf("%s\n", &str)
		fmt.Println(change(str))
	}
}
