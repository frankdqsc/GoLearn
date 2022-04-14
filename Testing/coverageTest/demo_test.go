package coverageTest

import "testing"

func GetArea(weight, height int) int {
	return weight * height
}

func TestGetArea(t *testing.T) {
	area := GetArea(40, 50)
	if area != 2000 {
		t.Error("测试失败")
	}
}
func BenchmarkGetArea(t *testing.B) {
	for i := 0; i < t.N; i++ {
		GetArea(40, 50)
	}
}
