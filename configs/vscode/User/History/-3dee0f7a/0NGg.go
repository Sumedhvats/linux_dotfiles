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
type DefaultSleeper struct{
}
func (s *DefaultSleeper) Sleep(){
	time.Sleep(1*time.Second);
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}
type SpyCountdownOperations struct {
	Calls []string
}
func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}
func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}
const sleep = "sleep"
const write = "write"

func Countdown(out io.Writer,s Sleeper) {
	for i:=3;i>0;i--{
		fmt.Fprintln(out,i)
		s.Sleep()
	}
	fmt.Fprint(out,"Go!")
}

func main() {
	sl:=&DefaultSleeper{}
	Countdown(os.Stdout,sl)
}