package main

import (
	"fmt"
)

func main() {
	//map使用for-range遍历 不能使用for
	cities := make(map[string]string)
	cities["1"] = "Tom"
	cities["2"] = "Alice"
	cities["3"] = "Georgia"
	cities["4"] = "Georgia"
	cities["4"] = "Georgia"

	for k, v := range cities {
		fmt.Printf("k=%v v=%v\n", k, v)
	}
}
