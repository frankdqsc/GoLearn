/* 定义一个4行4列的二维数组，逐个从键盘输入值
将第1行和第4行的数据进行交换
将第2行和第3行的数据进行交换 
*/
package main
import (
	"fmt"
)

func main (){
	var arr[4][4]int
	for i:=0; i<len(arr); i++ {
		for j:=0; j<len(arr[i]); j++ {
			fmt.Printf("请从键盘输入第%d行的第%d个数字：\n",i,j)
			fmt.Scanln(&arr[i][j])
		}
	}
	fmt.Println("arr=",arr)
	
/* 	for i :=0; i<1; i++{
		for j:=0; j<4; j++{

		}
	} */
	var slice = arr
	fmt.Println("slice=",slice)
	slice1:= slice[0:1]
	slice2:= slice[1:2]
	slice3:= slice[2:3]
	slice4:= slice[3:4]
	//must add "..."if append slice to another slice
	fmt.Println("slice4=",slice4)
	slice4 = append(slice4,slice3...) 
	slice4 = append(slice4,slice2...)
	slice4 = append(slice4,slice1...)
	
	fmt.Println("slice4=",slice4)
/* 
result:
arr= [[43 5 3 7] [2 8 2 9] [3 5 5 8] [3 8 3 978]]
slice= [[43 5 3 7] [2 8 2 9] [3 5 5 8] [3 8 3 978]]
slice4= [[3 8 3 978]]
slice4= [[3 8 3 978] [3 5 5 8] [2 8 2 9] [43 5 3 7]] */
}