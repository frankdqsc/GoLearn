package main 
import "fmt"

type Visitor struct {
	Name string
	Age int
}

func (visitor *Visitor) showPrice (){
	if visitor.Age > 18 {
		fmt.Printf("游客的名字为 %v 年龄为 %v 收费20元\n", visitor.Name, visitor.Age)
	}else{
		fmt.Printf("游客的名字为 %v 年龄为 %v 免费\n", visitor.Name, visitor.Age)
	}
}

func main (){
	var v Visitor
	
	for {
		fmt.Println("name：")
		fmt.Scanln(&v.Name)
		if v.Name == "n" {
			fmt.Println("quit")
			break
		}
		fmt.Println("age:")
		fmt.Scanln(&v.Age)
		
		v.showPrice()
	}
}