package main

import "fmt"

//对于N为2的幂的数，都有其二进制数 N&(N-1)=0
func isPowerOfTwo(n int) bool {
	return n > 0 && n&(n-1) == 0
}

func main(){
	var m bool
	m = isPowerOfTwo(16)
	fmt.Println(m)
}