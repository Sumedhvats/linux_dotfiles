package main
import "testing"

func TestAdder(t *testing.T){
	got:=Add(1,2)
	want:=3
	if got!=want{
		t.Errorf("expected %d but got ")
	}
}