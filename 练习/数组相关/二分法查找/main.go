/* 随机生成10个整数（1-100）使用冒泡排序进行排序，
然后使用二分查找法，查找是否有90这个元素，并显示下标
如果没有提示没有 */
package main
import (
	"fmt"
	"math/rand"
	"time"
)
func BubbleSort(intArr *[10]int){ 
	temp :=0
	for i:=0; i<len(*intArr)-1;i++{
		for j:=0; j<9; j++ {
			if(*intArr)[j] > (*intArr)[j+1]{
				temp = (*intArr)[j]
				(*intArr)[j] = (*intArr)[j+1]
				(*intArr)[j+1] = temp
			}
		}
	}
	fmt.Println("排序后=",(*intArr))
}
func BinaryFind(intArr *[10]int, leftIndex int, rightIndex int, findVal int){
	//判断leftindex 是否大于 rightindex	
	if leftIndex > rightIndex  {
		fmt.Println("Can't find")
		return
	}
	//find middle index
	middle := (leftIndex + rightIndex)/2
	
	if (*intArr)[middle] > findVal {
		BinaryFind(intArr, leftIndex, middle - 1, findVal) //递归调用  
	} else if (*intArr)[middle] < findVal {
		BinaryFind(intArr, middle + 1, rightIndex , findVal)
	} else {
		//find it
		fmt.Printf("找到了，下标为：%v \n",middle) //格式化输出
		fmt.Println("找到了，下标为:",middle)
	}
}
func main(){
	var intArr[10]int
	//求随机数
	rand.Seed(time.Now().UnixNano())
	for i:=0; i<len(intArr); i++ {
		intArr[i] = rand.Intn(100)
	}
	fmt.Println(intArr)
	BubbleSort(&intArr) //冒泡排序
	fmt.Println(intArr)
	BinaryFind(&intArr, 0, len(intArr)-1, 90)
}