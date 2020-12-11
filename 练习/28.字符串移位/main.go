package main

import "fmt"

func main() {

	str:= "abcde"
	step:= 5

	rs := []rune(str)
	len:= len(rs)
	var t []rune

	if step == len{
		fmt.Println(string(rs))
		return
	}
		t = make([]rune, 0)

	for i:=len-step; i<len; i++{
		t = append(t,rs[i])
	}
	for i:=0; i<len-step; i++{
		t = append(t,rs[i])
	}
	fmt.Println(string(t))
}