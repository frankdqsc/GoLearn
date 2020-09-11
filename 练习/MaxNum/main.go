package main

import (
    "fmt"
)

func main(){
	var intArr [5]int = [...]int {1, -1, 9, 90, 11}
	maxVal :=intArr[0]

	for i := 1; i < len(intArr); i++ {
		if maxVal < intArr[i] {
			maxVal = intArr[i]
		}
	}
	fmt.Printf("maxVal=%v",maxVal)
}