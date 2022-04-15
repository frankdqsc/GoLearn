package main

import (
	"testing"
)

func add(x, y int)int{
	return x + y
}
func TestAdd(t *testing.T){ //专用类型 T 用来控制测试结果和行为
	if add(1,2) != 3{
		t.FailNow()
	}
}
