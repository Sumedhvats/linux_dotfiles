package golearning

import (
	"testing"
)
func TestWallet(t *testing.T) {
	dict := map[string]string{"test":"this is a test"}
	got:=Search(dict,"test")
	want:="this is a test"
	if got!=want {
		t.Errorf("got %q instead of %q",got,want)
	}

}