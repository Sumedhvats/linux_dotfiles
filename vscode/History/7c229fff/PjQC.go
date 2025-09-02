package golearning

import (
	"testing"
)
func assertStrings(t testing.TB,got,want string){
	t.Helper()
	if got!=want{
		t.Errorf("got %q want %q", got, want)
	}

}
func TestSearch(t *testing.T) {
	dict := Dictionary{"test": "this is a test"}
	t.Run("known word", func(t *testing.T) {
		got, _ = dict.Search("test")
		want := "this is a test"
		assertStrings(t, got,want)
	})
	t.Run("unknown word", func(t *testing.T) {
		_, err = dict.Search("adgaeg")
		want := "could not find the word you were looking for"
		if err == nil {
			t.Fatal("expected to get an error.")
		}
		assertStrings(t, err.Error(), want)

	})
}
