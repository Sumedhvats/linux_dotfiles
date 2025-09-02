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

// result is a package-level variable to store the outcome of the benchmarked function.
// It prevents compiler optimizations that might remove the function call.
var result string

func BenchmarkRepeat(b *testing.B) {
	var r string
	for i := 0; i < b.N; i++ {
		r = Repeat("a")
	}
	// Assign to the global variable so staticcheck and the compiler
	// know this value is used.
	result = r
}
