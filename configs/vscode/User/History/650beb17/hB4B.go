package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {

	slowURL := delayedServer(20 * time.Millisecond)
	fastURL := delayedServer(0 * time.Millisecond)

	want := fastURL.URL
	got := Racer(slowURL.URL, fastURL.URL)

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
	slowURL.Close()
	fastURL.Close()
}
func delayedServer(t time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(t)
		w.WriteHeader(http.StatusOK)
	}))

}
