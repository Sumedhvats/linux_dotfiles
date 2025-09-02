var result string

func BenchmarkRepeat(b *testing.B) {
    var r string
    for i := 0; i < b.N; i++ {
        r = Repeat("a")
    }
    result = r
}
