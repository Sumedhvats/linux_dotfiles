package golearning

import (
	"testing"
)
func TestSearch(t *testing.T) {
	dict := Dictionary{"test":"this is a test"}
	got:=dict.Search("test")
	want:="this is a test"
	if got!=want {
		t.Errorf("got %q instead of %q",got,want)
	}

}