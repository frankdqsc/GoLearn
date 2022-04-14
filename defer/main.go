package main

import "fmt"

/*10 1 2 3
20 0 2 2
2 0 2 2
1 1 3 4
defer中的函数 calc("10", a, b)和 calc("20", a, b) 先执行，然后按照 FILO 执行 defer*/

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

//return 和 panic 会终止当前函数流程，引发延迟调用
func test()(z int){  // return先更新返回值z=100; call defer z=200; return 200
	defer func(){
		println("defer:", z)  //defer: 100
		z +=100 //test: 200
	}()

	return 100
}
func main() {
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b))
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1

	println("test:", test())
}

