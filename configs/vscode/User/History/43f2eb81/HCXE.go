package m
import "testing"
func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a")
	}
}