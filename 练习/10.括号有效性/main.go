//括号检验
//多行输入
package main

import (
	"fmt"
	"reflect"
)

func Cal(s string) bool {
	keys := map[rune]rune{')': '(', ']': '[', '}': '{'}
	var stack []rune
	for _, char := range s {
		fmt.Println(reflect.TypeOf(char))
		if char == '(' || char == '{' || char == '[' {
			stack = append(stack, char)
		} else if len(stack) > 0 && keys[char] == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	return true
}

func main() {

	var s string
	//s = "{ [ ( ) ] ( [ ] ) } ( ) [ ]"
	var T int
	fmt.Scanf("%d\n", &T)

	for i := 0; i < T; i++ {
		fmt.Scanf("%s\n", &s)
		fmt.Println(Cal(s))
	}
}
