package main

import (
	"bytes"
	"fmt"
	
)

func main() {
	buffer := &bytes.Buffer{}
	Countdown(buffer)
	fmt.Println("Output:", buffer.String())
}
