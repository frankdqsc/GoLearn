package main

/*随机抽取扑克牌中的n张牌，判断是不是顺子，即这5张牌是不是连续的。
其中A看成1，J看成11，Q看成12，K看成13，大小王可以看成任何需要的数字。*/
import (
	"fmt"
	"sort"
)
func judge(slice[] int) bool{

	sort.Ints(slice) //[0,0,1,2,4]
	zeros := 0  //用0表示大小王,记录大小王个数
	for i,n := range slice{  // n == slice[i]

		if slice[i] == 0{
			zeros ++  //记录大小王个数
			continue  //continue循环下一个key
		}
		// i!=0 是为了slice[i-1]不越界
		if i!=0 && n == slice[i-1] {  //前后两张一样，slice[i]此时为1 和 0 比较，除了0的重复直接false
			return false
		}
		if i!=0 && slice[i-1] == 0 {  //说明slice[i-1]是大小王 计算过了，继续循环
			continue
		}
		if i!=0 && n-1 != slice[i-1] {  //两个相邻的数不是相差1，则进入判断
			pre := slice[i-1]
			target := n-1
			if zeros > 0 {
				for zeros >0 && pre != target {  //在 0 耗尽前，持续填补
					pre++
					zeros--
				}
				if pre != target {  // 0耗尽也追不上
					return false
				}
			}else {  //没大小王，不能填补
				return false
			}
		}
	}
	return true
}

func main(){
	//var n int
	var t bool
	var slice[] int
	slice = []int{0,0,8,5,4}
/*	for i:=0;i<5;i++{
		fmt.Scan(&n)
		slice = append(slice,n)
	}*/
	t = judge(slice)
	fmt.Println(t)
}
