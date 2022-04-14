package main
import "fmt"
func main() {

	//var a [][]int
	var tmp []int
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			tmp = append(tmp, j)
		}
		fmt.Println(tmp)
	}


}
