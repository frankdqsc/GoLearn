package main

import (
	"fmt"
)

type Singleton struct {
}

var singleton *Singleton = new(Singleton)  //饿汉式：提前new出来，后面调用只返回这个

func GetSingletonObj() *Singleton{
	fmt.Println(&singleton)
	return singleton
}
func main(){
	for i:=0;i<5;i++{
		GetSingletonObj()
	}
}
