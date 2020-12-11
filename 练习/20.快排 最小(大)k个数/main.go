package main

import "fmt"

func QuickSort(left,right int, arr*[5]int) {
	if left < right{
		i , j ,p := left,right,arr[(left+right)/2]
		for i<= j{
			for arr[i] < p{
				i++
			}
			for arr[j] > p{
				j--
			}
			if i<=j{
				arr[i],arr[j]=arr[j],arr[i]
				i++
				j--
			}
		}
		if left < j{
			QuickSort(left,j,arr)
		}
		if right>i{
			QuickSort(i,right,arr)
		}
	}
}
func main(){
	var arr [5]int
	var n,k int
	for i:=0;i<5;i++{
		fmt.Scan(&n)
		arr[i] = n
	}
	fmt.Scan(&k)

	QuickSort(0,len(arr)-1,&arr)
	fmt.Println(arr[:k])
}