package performanceTestDemo

import "testing"

func GetArea(weight, height int) int {
	return weight * height
}
func BenchmarkGetArea(t *testing.B) {
	for i := 0; i < t.N; i++ {
		GetArea(40, 50)
	}
}
