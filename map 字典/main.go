package main

import "fmt"

// map 为无序键值对集合。字典不允许直接地址指向。在结构体和数组作为字典成员时，修改时只能返回整个值再修改内部（没举例）
func main(){
	m := make(map[string]int)
	m["a"] = 1
	m["b"] = 2
	fmt.Println(m)

	if v, ok := m["b"]; ok{
		println(v)
	}

}
