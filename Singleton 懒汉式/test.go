package main

import (
	"fmt"
	"sync"
)

type Singleton struct {
}

var once sync.Once
var singleton *Singleton

func GetSingletonObj() *Singleton{  //懒汉式：调用时才执行创建
	once.Do(func(){
			fmt.Println(&singleton)
			singleton = new(Singleton) // 或者singleton = &Singleton{}
	})
	return singleton
}
func main(){
	var wg sync.WaitGroup  //加waitGroup保证线程安全
	for i:=0;i<5;i++{
		wg.Add(1)
		go func(){
			GetSingletonObj()
			wg.Done()
		}()
	}
	wg.Wait()
}
