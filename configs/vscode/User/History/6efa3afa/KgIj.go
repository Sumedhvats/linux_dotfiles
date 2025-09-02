package m
const repeatCount = 5

func Repeat(character string) string {
	var repeated string.builder
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated
}