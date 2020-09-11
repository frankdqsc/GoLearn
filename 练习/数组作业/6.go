/* 试写出实现查找的核心代码，比如已知数组arr[10]string;
里面保存了十个元素，现要查找“AA”在其中是否存在，打印提示
如果有多个，也要找到对应的下标 */
package main
import (
	"fmt"
)
func main (){
	count :=0
	countIndex := 0
	arr1 := [6]string{"f","Ag","AA","An","AAA","AA"}
	for i:=0; i<len(arr1); i++{
		if arr1[i] == "AA"{
			count++
			countIndex = i
			fmt.Println("index=",countIndex)
		}
	}
	fmt.Println("count=",count)
}