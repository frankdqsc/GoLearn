// 编写一个Cal结构体，两个字段Num1和Num2
// 方法GetSub(name string)
// 使用反射遍历Cal结构体的所有字段信息
// 使用反射机制完成对GetSub的调用 输出形式为“Tom 完成了减法运算，8-3=5”

//之前→创建结构体实例 调用方法
//现在→创建结构体实例 将结构体实例传递给方法 用反射调用方法

package main

import(
	"reflect"
	"fmt"
)

//定义结构体 Cal
type Cal struct{
	Num1 int
	Num2 int
}

//给结构体绑定方法GetSub
func (s Cal) GetSub(Num1,Num2 int) int {
	return Num1 - Num2

}

//给两个数字赋值
func (s Cal) Set(num1 int,num2 int){
	s.Num1 = num1
	s.Num2 = num2
}

func TestStruct(a interface{}){
	typ := reflect.TypeOf(a)
	fmt.Printf("type is %d fields\n",typ)
	val := reflect.ValueOf(a)
	//类别判断 如果不是结构体则退出
	kd := val.Kind()
	if kd != reflect.Struct{
		fmt.Println("except struct")
		return
	}

	num := val.NumField()
	fmt.Printf("struct has %d fields\n",num)
	
	for i:=0;i<num;i++{
		fmt.Printf("Field %d:值为=%v\n",i,val.Field(i))
	}

	//获取该结构体有多少个方法
	numOfMethod := val.NumMethod()
	fmt.Printf("struct has %d method\n",numOfMethod)
	
	//传入切片 取出也是切片
	// var params []reflect.Value
	// params = append(params,reflect.ValueOf(900))
	// params = append(params,reflect.ValueOf(100))
	// res := val.Method(0).Call(params)
	// fmt.Println("res=",res[0].Int())

	//fmt.Println("val.Field(0)=",val.Field(0))
	//fmt.Printf("val.Field(0)类型是:%T\n",val.Field(0))  //val.Field(0)类型是:reflect.Value
	//传入切片 取出也是切片
	var params []reflect.Value
	params = append(params,val.Field(0))
	params = append(params,val.Field(1))
	res := val.Method(0).Call(params)
	fmt.Printf("%d -%v",val.Field(0),val.Field(1),res[0].Int())
}
func main(){
	//创建一个Cal实例
	var a Cal = Cal{
		Num1 : 300,
		Num2 : 200,
	}
	//将Cal实例传递给TestStruct函数
	TestStruct(a)

}