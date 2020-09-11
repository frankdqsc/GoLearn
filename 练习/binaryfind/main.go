package main
import (
	"fmt"
)

func BinaryFind(arr *[6]int, leftIndex int, rightIndex int, findVal int){
	//判断leftindex 是否大于 rightindex	
	if leftIndex > rightIndex  {
		fmt.Println("Can't find")
		return
	}
	//find middle index
	middle := (leftIndex + rightIndex)/2
	
	if (*arr)[middle] > findVal {
		BinaryFind(arr, leftIndex, middle - 1, findVal) //递归调用  
	} else if (*arr)[middle] < findVal {
		BinaryFind(arr, middle + 1, rightIndex , findVal)
	} else {
		//find it
		fmt.Printf("找到了，下标为：%v \n",middle) //格式化输出
		fmt.Println("找到了，下标为:",middle)
	}
}
func main(){ 
	arr := [6]int{1, 8, 10, 89, 1000, 1234}
	//test
	BinaryFind(&arr, 0, len(arr)-1, 1000) //二分查找
}