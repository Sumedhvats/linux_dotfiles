package golearning

import (
	"bytes"
	"fmt"
	"golearning"
)

func main() {
	buffer := &bytes.Buffer{}
	golearning.Countdown(buffer)
	fmt.Println("Output:", buffer.String())
}
