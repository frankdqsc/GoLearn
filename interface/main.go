package main

func Print(v interface{}) {
	println(v)
}
func main() {
	type Test struct{}
	v := Test{}
	Print(v)
}

