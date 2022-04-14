package main

import "fmt"

func main(){
	// 不管 for循环还是 range 迭代，其定义的局部变量都会重复使用
	slice := [3]string{"a","b","c"}
	for i, s:= range slice{
		fmt.Println(&i, &s)
	}
	for i,_:= range slice{
		fmt.Println(i, slice[i])
	}
	data := [3]int{10, 20, 30}
	//for i, x:= range data{ //从 data 复制中取值，data 复制中包含3个值，第一次时就完成三次相加
	//	if i == 0{
	//		data[0] += 100
	//		data[1] += 200
	//		data[2] += 300
	//	}
	//	fmt.Printf("x: %d, data: %d\n", x, data[i])
	//}

	for i, x:= range data[:]{  //仅复制 slice，不包括底层 array --> i==0 时，if里完成 110 220 330， i==1 时，x= data[1] = 220
		if i == 0{
			data[0] += 100
			data[1] += 200
			data[2] += 300
		}
		fmt.Printf("x: %d, data: %d\n", x, data[i])
	}
}
