package unitTestDemo

import (
	"testing"
)

func GetArea(weight, height int) int {
	return weight * height
}

//专用类型 T 用来控制测试结果和行为, 单元测试用t *testing.T， 性能测试用t *testing.B
func TestGetArea(t *testing.T) {
	area := GetArea(40, 50)
	if area != 2000 {
		t.Error("测试失败")
	}
}
