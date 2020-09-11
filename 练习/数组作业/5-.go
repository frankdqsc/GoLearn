/* 试保存 13579五个奇数到数组 并倒序打印 */
package main 
import(
	"fmt"
)
func main(){
	var arr[5]int = [...]int{1,3,5,7,9}
	for i:= 4; i>-1; i-- {
		fmt.Println(arr[i])
	} 
}
/* //define stack
type Stack struct{
	size int
	top int
	data []interface{}
}
//initialize
func MakeStack(size int)*Stack{
	q:=Stack{}
	q.size = size
	q.data = make([]interface{},size)
	return q
}
//in
func (s *Stack)Push(i interface{})bool{
	if s.IsFull(){
		return false
	}
	s.data[s.top] = int
	s.top++
	return true
}
//out
func (s *Stack)Pop()(interface{},error){
	if s.IsEmpty(){
		//return nil,fmt.Errorf( format:"empty stack")'
		return nil, fmt.Errorf("empty stack")
	}
	r:=s.data[s.top]
	s.top--
	return r,nil
}

func main(){
	var arr[5]int = [...]int{1,3,5,7,9}
	for i:=0; i<len(arr); i++ {
		fmt.Println(arr[i])
		Push();
		Pop();
	} 
} */