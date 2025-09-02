package main

import (
	"net/http/httptest"
	"net/http/test"
	"testing"
)
func TestRacer(t *testing.T) {

	slowURL := httptest.
	fastURL := "http://www.quii.dev"

	want := fastURL
	got := Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}