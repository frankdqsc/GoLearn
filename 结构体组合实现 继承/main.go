package main

import (
	"fmt"
)

//Go 实现继承的语义不是通过 extends 关键字而是通过结构体组合的方式
type Pet struct {
}

func (p *Pet) Skill() {
	fmt.Println("能文能武的宠物")
}

type Cat struct {  //猫是能够抓老鼠的宠物,Go 采用组合的方式表达继承的语义.
	p *Pet
}

func (c *Cat) Catch() {
	fmt.Println("老鼠天敌喵喵喵")
}

type Dog struct {
	p *Pet
}

func (d *Dog) Navigate() {
	fmt.Println("自带导航汪汪汪")
}

func main() {
	p := new(Pet)

	d := new(Dog)
	d.p = p

	d.Navigate()
	d.p.Skill()

	fmt.Println()

	c := new(Cat)
	c.p = p
	c.Catch()
	c.p.Skill()
}
