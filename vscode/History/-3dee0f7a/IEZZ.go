package main

import (
	"fmt"
	"io"
	"os"
)

func Countdown(out io.Writer) {
	for i:=0;i<3;i++{
		fmt.Fprintln(out,i)
	}
}

func main() {
	Countdown(os.Stdout)
}