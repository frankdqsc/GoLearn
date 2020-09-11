package main
import (
	"fmt"
)

func fbn(n int) ([]uint64) {
	//声明一个切片，大小为 n
	fbnSlice := make ([]uint64,n)
	//第一个和第二个数的斐波那契为1
	fbnSlice[0] = 1
	fbnSlice[1] = 1
	//进行for循环来存放斐波那契数列
	for i :=2; i<n; i++ {
		fbnSlice[i] = fbnSlice[i - 1]+fbnSlice[i - 2]
	}

	return fbnSlice
}

func main(){
	fbnSlice := fbn(20)
	fmt.Println("fnbSlice=",fbnSlice)
}