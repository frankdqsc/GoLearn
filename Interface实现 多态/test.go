package main

import (
	"fmt"
)

/*通过golang里面有一个接口类型interface，
任何类型只要实现了接口类型，都可以赋值，如果接口类型是空，那么所有的类型都实现了它。其作⽤用类似⾯面向对象语⾔言中的根对象 object。
golang里面的多态就是用接口类型实现的，即定义一个接口类型，里面声明一些要实现的功能，注意，只要声明，不要实现*/

type Biology interface {
	sayhi()
}


type Man struct {
	name string
	age  int
}

type Monster struct {
	name string
	age  int
}

func (this *Man) sayhi()  { // 实现抽象方法1
	fmt.Printf("Man[%s, %d] sayhi\n", this.name, this.age)
}

func (this *Monster) sayhi()  { // 实现抽象方法1
	fmt.Printf("Monster[%s, %d] sayhi\n", this.name, this.age)
}

func WhoSayHi(i Biology) {
	i.sayhi()
}


func main() {
	man := &Man{"我是人", 100}
	monster := &Monster{"妖怪", 1000}
	defer println("c")
	WhoSayHi(man)
	WhoSayHi(monster)
}