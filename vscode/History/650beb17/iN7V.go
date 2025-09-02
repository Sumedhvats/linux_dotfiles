package main

import (
	"net/http"
	"net/http/httptest"
	"net/http/test"
	"testing"
	"time"
)
func TestRacer(t *testing.T) {

	slowURL := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				time.Sleep(20 * time.Millisecond)
				w.Write(http.StatusOK)

	}))
	fastURL := "http://www.quii.dev"

	want := fastURL
	got := Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}