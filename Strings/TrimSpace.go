package main

import (
	"fmt"
	"strings"
)
//golang字符串操作
func main(){
	s := " Hello world hello world    "  //删除首尾的空格

	//删除s首尾连续的的空白字符
	ret := strings.TrimSpace(s)
	fmt.Println(ret) // Hello world hello world

	//以连续的空白字符为分隔符，将s切分成多个子串，结果中不包含空白字符本身。
	//空白字符有：\t, \n, \v, \f, \r, ’ ‘, U+0085 (NEL), U+00A0 (NBSP) 。
	//如果 s 中只包含空白字符，则返回一个空列表
	ret1 := strings.Fields(s)
	fmt.Println(ret1)
}