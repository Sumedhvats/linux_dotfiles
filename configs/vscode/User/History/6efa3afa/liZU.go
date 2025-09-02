package m

import "testing"

const repeatCount = 5

func Repeat(character string) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated
}

var result string // store the output to avoid elimination

func BenchmarkRepeat(b *testing.B) {
	var r string
	for i := 0; i < b.N; i++ {
		r = Repeat("a")
	}
	result = r
}
