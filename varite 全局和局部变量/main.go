package main

import "fmt"

var nickname = "大虾"

func show() {
	print(nickname)
}

func change() {
	nickname := "111"
	print(nickname)
}

func main() {
	show()
	change()
	show()
	fmt.Println(nickname)
}

// 大虾小虾大虾大虾
