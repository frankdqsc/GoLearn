package monster

import (
	_ "fmt"
	"testing"
)

//测试用例，测试store方法
func TestStore(t *testing.T ){

	//先创建一个Monster实例
	monster := Monster{
		Name:"红孩儿",
		Age:18,
		Skill:"吐火",
	}
	res := monster.Store()
	if !res {
		t.Fatalf("monster.Store()错误，希望为=%v 实际为=%v",true,res)
	}
	t.Logf("monster.Store()测试成功！")
}

func TestReStore(t *testing.T ){
	//先创建monster实例，不需要指定字段的值
	var monster = &Monster{}
	res := monster.Restore()
	if !res {
		t.Fatalf("monster.ReStore()错误，希望为=%v 实际为=%v",true,res)
	}

	//进一步判断
	if monster.Name != "红孩儿"{
		t.Fatalf("monster.ReStore()错误，希望为=%v 实际为=%v","红孩儿",monster.Name)
	}

	t.Logf("monster.ReStore()测试成功！")
}