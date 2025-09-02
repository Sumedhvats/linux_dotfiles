package main

import (
	"fmt"
	"io"
	"os"
	"time"
)
type Sleeper interface {
	Sleep()
}
type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}
func 

func Countdown(out io.Writer,s Sleeper) {
	for i:=3;i>0;i--{
		fmt.Fprintln(out,i)
		time.Sleep(1*time.Second)
	}
	fmt.Fprint(out,"Go!")
}

func main() {
	Countdown(os.Stdout,)
}