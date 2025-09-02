func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a")
	}
}