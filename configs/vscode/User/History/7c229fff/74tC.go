package golearning

import (
	"testing"
)
func TestSearch(t *testing.T) {
	dict := Dictionary{"test":"this is a test"}
t.Run("known word",func(t *testing.T) {
	got,_=dict.Search()
})
}