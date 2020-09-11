/* 已知有个升序排好序的数组，要求插入一个元素，
最后打印该数组，顺序依然是升序
使用append内置函数进行元素追加
切片的创建 方式1：通过引用数组 方式2：通过make来创建切片 */
package main
import (
	"fmt"
)
func BubbleSort(slice1 []int){
	temp :=0
	for i:=0; i<len(slice1)-1;i++{
		for j:=0; j<len(slice1)-1; j++ {
			if(slice1)[j] > (slice1)[j+1]{
				temp = (slice1)[j]
				(slice1)[j] = (slice1)[j+1]
				(slice1)[j+1] = temp
			}
		}
	}
	fmt.Println("排序后slice1=",(slice1))
}
func main(){
	//way one
	var arr[5] int = [...]int{8,4,80,87,94}
	var slice1 = arr[0:5]
	//fmt.Println("slice1=",slice1)
	slice1 = append(slice1,10)
	fmt.Println("slice1=",slice1)
	BubbleSort(slice1)
	
	//way two
	/* var slice2 []int = make([]int,5,10)
	for i:=0; i<len(slice2);i++ {
		fmt.Printf("请输入5个数字：\n",i)
		fmt.Scanln(&slice2[i])
	}
	fmt.Println(slice2)
	slice2 = append(slice2,18)
	fmt.Println(slice2)
	// haven't set bubblesort a public metho */d!!!!!
}

